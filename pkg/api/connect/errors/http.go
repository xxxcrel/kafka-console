// Copyright 2024 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package errors

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/textproto"
	"strings"

	commonv1alpha1 "buf.build/gen/go/redpandadata/common/protocolbuffers/go/redpanda/api/common/v1alpha1"
	"connectrpc.com/connect"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/genproto/googleapis/rpc/code"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/anypb"
)

var protoJSONMarshaler = &runtime.JSONPb{
	MarshalOptions: protojson.MarshalOptions{
		UseProtoNames: true, // use snake_case
		// Do not use EmitUnpopulated, so we don't emit nulls (they are ugly, and provide no benefit. they transport no information, even in "normal" json).
		EmitUnpopulated: false,
		// Instead, use EmitDefaultValues, which is new and like EmitUnpopulated, but
		// skips nulls (which we consider ugly, and provides no benefit over skipping the field)
		EmitDefaultValues: true,
	},
	UnmarshalOptions: protojson.UnmarshalOptions{
		DiscardUnknown: true,
	},
}

// HandleHTTPError serializes the given error and writes it using the protoJSONMarshaler.
// This function can handle errors of type connect.Error as well, so that err details are
// printed properly.
func HandleHTTPError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	var st *spb.Status

	var connectErr *connect.Error
	if errors.As(err, &connectErr) {
		st = &spb.Status{
			Code:    int32(connectErr.Code()),
			Message: connectErr.Message(),
		}
		for _, detail := range connectErr.Details() {
			anyDetail := &anypb.Any{
				TypeUrl: detail.Type(),
				Value:   detail.Bytes(),
			}
			st.Details = append(st.Details, anyDetail)
		}
	} else {
		st = &spb.Status{
			Code:    int32(connect.CodeOf(err)),
			Message: err.Error(),
		}
	}

	NiceHTTPErrorHandler(ctx, nil, protoJSONMarshaler, w, r, status.ErrorProto(st))
}

// NiceHTTPErrorHandler is a clone of grpc-gateway's
// runtime.DefaultHTTPErrorHandler, with one difference: it uses a modified
// variant of google.rpc.Status, where code is ENUM instead of int32.
func NiceHTTPErrorHandler(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	const fallback = `{"code":"INTERNAL", "message":"failed to marshal error message"}`

	var customStatus *runtime.HTTPStatusError
	if errors.As(err, &customStatus) {
		err = customStatus.Err
	}
	s := status.Convert(err)
	pb := StatusToNice(s.Proto())

	w.Header().Del("Trailer")
	w.Header().Del("Transfer-Encoding")

	contentType := marshaler.ContentType(pb)
	w.Header().Set("Content-Type", contentType)

	if s.Code() == codes.Unauthenticated {
		w.Header().Set("WWW-Authenticate", s.Message())
	}

	buf, merr := marshaler.Marshal(pb)
	if merr != nil {
		grpclog.Infof("Failed to marshal error message %q: %v", s, merr)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := io.WriteString(w, fallback); err != nil {
			grpclog.Infof("Failed to write response: %v", err)
		}
		return
	}

	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		grpclog.Infof("Failed to extract ServerMetadata from context")
	}

	handleForwardResponseServerMetadata(w, md)

	// RFC 7230 https://tools.ietf.org/html/rfc7230#section-4.1.2
	// Unless the request includes a TE header field indicating "trailers"
	// is acceptable, as described in Section 4.3, a server SHOULD NOT
	// generate trailer fields that it believes are necessary for the user
	// agent to receive.
	doForwardTrailers := requestAcceptsTrailers(r)

	if doForwardTrailers {
		handleForwardResponseTrailerHeader(w, md)
		w.Header().Set("Transfer-Encoding", "chunked")
	}

	st := runtime.HTTPStatusFromCode(s.Code())
	if customStatus != nil {
		st = customStatus.HTTPStatus
	}

	w.WriteHeader(st)
	if _, err := w.Write(buf); err != nil {
		grpclog.Infof("Failed to write response: %v", err)
	}

	if doForwardTrailers {
		handleForwardResponseTrailer(w, md)
	}
}

var defaultOutgoingHeaderMatcher = func(key string) (string, bool) {
	return fmt.Sprintf("%s%s", runtime.MetadataHeaderPrefix, key), true
}

func handleForwardResponseServerMetadata(w http.ResponseWriter, md runtime.ServerMetadata) {
	for k, vs := range md.HeaderMD {
		if h, ok := defaultOutgoingHeaderMatcher(k); ok {
			for _, v := range vs {
				w.Header().Add(h, v)
			}
		}
	}
}

func requestAcceptsTrailers(req *http.Request) bool {
	te := req.Header.Get("TE")
	return strings.Contains(strings.ToLower(te), "trailers")
}

func handleForwardResponseTrailerHeader(w http.ResponseWriter, md runtime.ServerMetadata) {
	for k := range md.TrailerMD {
		tKey := textproto.CanonicalMIMEHeaderKey(fmt.Sprintf("%s%s", runtime.MetadataTrailerPrefix, k))
		w.Header().Add("Trailer", tKey)
	}
}

func handleForwardResponseTrailer(w http.ResponseWriter, md runtime.ServerMetadata) {
	for k, vs := range md.TrailerMD {
		tKey := fmt.Sprintf("%s%s", runtime.MetadataTrailerPrefix, k)
		for _, v := range vs {
			w.Header().Add(tKey, v)
		}
	}
}

// StatusToNice converts a google.rpc.Status to cloudv1alpha1.ErrorStatus,
// which is "nicer" variant with Code as Enum.
func StatusToNice(s *spb.Status) *commonv1alpha1.ErrorStatus {
	pb := commonv1alpha1.ErrorStatus{
		Code:    code.Code(s.Code),
		Message: s.Message,
		Details: s.Details,
	}

	return &pb
}

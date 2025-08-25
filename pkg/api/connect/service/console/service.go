// Copyright 2023 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

// Package kconsole contains the implementation of all Console service RPC endpoints.
package console

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/xxxcrel/kafka-console/pkg/api/connect/service/clusterstatus"
	"go.uber.org/zap"

	apierrors "github.com/xxxcrel/kafka-console/pkg/api/connect/errors"
	"github.com/xxxcrel/kafka-console/pkg/kconsole"
)

// Service that implements the ConsoleServiceHandler interface.
type Service struct {
	logger     *zap.Logger
	consoleSvc kconsole.Servicer
}

// NewService creates a new Console service handler.
func NewService(
	logger *zap.Logger,
	consoleSvc kconsole.Servicer,
) *Service {
	return &Service{
		logger:     logger,
		consoleSvc: consoleSvc,
	}
}

// ListMessages consumes a Kafka topic and streams the Kafka records back.
func (api *Service) ListMessages(
	ctx context.Context,
	req *ListMessagesRequest,
) error {
	//lmq := httptypes.ListMessagesRequest{
	//	TopicName:             req.Msg.GetTopic(),
	//	StartOffset:           req.Msg.GetStartOffset(),
	//	StartTimestamp:        req.Msg.GetStartTimestamp(),
	//	PartitionID:           req.Msg.GetPartitionId(),
	//	MaxResults:            int(req.Msg.GetMaxResults()),
	//	FilterInterpreterCode: req.Msg.GetFilterInterpreterCode(),
	//	Enterprise:            req.Msg.GetEnterprise(),
	//}
	//
	//interpreterCode, err := lmq.DecodeInterpreterCode()
	//if err != nil {
	//	return apierrors.NewConnectError(
	//		connect.CodeInvalidArgument,
	//		fmt.Errorf("failed decoding provided interpreter code: %w", err),
	//		apierrors.NewErrorInfo(commonv1alpha1.Reason_REASON_INVALID_INPUT.String()),
	//	)
	//}
	//
	//// test compile
	//code := fmt.Sprintf(`var isMessageOk = function() {%s}`, interpreterCode)
	//_, err = goja.Compile("", code, true)
	//if err != nil {
	//	return apierrors.NewConnectError(
	//		connect.CodeInvalidArgument,
	//		fmt.Errorf("failed to compile provided interpreter code: %w", err),
	//		apierrors.NewErrorInfo(commonv1alpha1.Reason_REASON_INVALID_INPUT.String()),
	//	)
	//}
	//
	//// Request messages from kafka and return them once we got all the messages or the context is done
	//listReq := kconsole.ListMessageRequest{
	//	TopicName:             lmq.TopicName,
	//	PartitionID:           lmq.PartitionID,
	//	StartOffset:           lmq.StartOffset,
	//	StartTimestamp:        lmq.StartTimestamp,
	//	MessageCount:          lmq.MaxResults,
	//	FilterInterpreterCode: interpreterCode,
	//	Troubleshoot:          req.Msg.GetTroubleshoot(),
	//	IncludeRawPayload:     req.Msg.GetIncludeOriginalRawPayload(),
	//	IgnoreMaxSizeLimit:    req.Msg.GetIgnoreMaxSizeLimit(),
	//	KeyDeserializer:       fromProtoEncoding(req.GetKeyDeserializer()),
	//	ValueDeserializer:     fromProtoEncoding(req.GetValueDeserializer()),
	//}
	//
	//timeout := 35 * time.Second
	//if req.Msg.GetFilterInterpreterCode() != "" || req.Msg.GetStartOffset() == kconsole.StartOffsetNewest {
	//	// Push-down filters and StartOffset = Newest may be long-running streams.
	//	// There's already a client-side provided timeout which we usually trust.
	//	// But additionally we want to ensure it never takes much longer than that.
	//	timeout = 31 * time.Minute
	//}
	//
	//ctx, cancel := context.WithTimeoutCause(ctx, timeout, errors.New("list fetch timeout"))
	//defer cancel()
	//
	//progress := &streamProgressReporter{
	//	logger:           api.logger,
	//	request:          &listReq,
	//	stream:           stream,
	//	messagesConsumed: atomic.Int64{},
	//	bytesConsumed:    atomic.Int64{},
	//}
	//progress.Start(ctx)
	//
	//return api.consoleSvc.ListMessages(ctx, listReq, progress)
	fmt.Printf(req.Topic)
	return nil
}

// PublishMessage serialized and produces the records.
//
//nolint:gocognit // Complexity is rather high, but not unreasonable
func (api *Service) PublishMessage(
	ctx context.Context,
	req *PublishMessageRequest,
) (*PublishMessageResponse, error) {

	recordHeaders := make([]kgo.RecordHeader, 0, len(req.Headers))
	for _, h := range req.Headers {
		recordHeaders = append(
			recordHeaders, kgo.RecordHeader{
				Key:   h.Key,
				Value: h.Value,
			},
		)
	}

	keyInput := rpcPublishMessagePayloadOptionsToSerializeInput(req.Key)
	valueInput := rpcPublishMessagePayloadOptionsToSerializeInput(req.Value)
	compression := rpcCompressionTypeToKgoCodec(req.Compression)

	prRes, prErr := api.consoleSvc.ProduceRecord(
		ctx, req.Topic, req.PartitionId, recordHeaders,
		keyInput, valueInput, req.UseTransactions, compression,
	)

	if prErr == nil && prRes != nil && prRes.Error != "" {
		prErr = errors.New(prRes.Error)
	}

	if prErr != nil {
		code := connect.CodeInternal

		details := []*connect.ErrorDetail{}

		if prRes != nil {
			if len(prRes.KeyTroubleshooting) > 0 {
				code = connect.CodeInvalidArgument

				for _, ktr := range prRes.KeyTroubleshooting {
					errInfo := apierrors.NewErrorInfo(
						clusterstatus.Reason_REASON_CONSOLE_ERROR.String(), apierrors.KeyVal{
							Key: ktr.SerdeName, Value: ktr.Message,
						},
					)

					if detail, detailErr := connect.NewErrorDetail(errInfo); detailErr == nil {
						details = append(details, detail)
					}
				}
			}

			if len(prRes.ValueTroubleshooting) > 0 {
				code = connect.CodeInvalidArgument

				for _, vtr := range prRes.ValueTroubleshooting {
					errInfo := apierrors.NewErrorInfo(
						clusterstatus.Reason_REASON_CONSOLE_ERROR.String(), apierrors.KeyVal{
							Key: vtr.SerdeName, Value: vtr.Message,
						},
					)

					if detail, detailErr := connect.NewErrorDetail(errInfo); detailErr == nil {
						details = append(details, detail)
					}
				}
			}
		}

		err := connect.NewError(
			code,
			prErr,
		)

		for _, ed := range details {
			ed := ed
			err.AddDetail(ed)
		}

		return nil, err
	}

	return &PublishMessageResponse{
		Topic:       prRes.TopicName,
		PartitionId: prRes.PartitionID,
		Offset:      prRes.Offset,
	}, nil
}

// Copyright 2022 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file https://github.com/xxxcrel/redpanda/blob/dev/licenses/bsl.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package kconsole

import (
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kmsg"
)

// DeleteACLsResponse is the response to deleting ACL resources.
type DeleteACLsResponse struct {
	ErrorMessages []string `json:"errorMessage"`
	MatchedACLs   int      `json:"matchedACLs"`
	DeletedACLs   int      `json:"deletedACLs"`
}

// DeleteACLs deletes Kafka ACLs based on a given filter.
func (s *Service) DeleteACLs(ctx context.Context, filter kmsg.DeleteACLsRequestFilter) (DeleteACLsResponse, error) {
	cl, _, err := s.kafkaClientFactory.GetKafkaClient(ctx)
	if err != nil {
		return DeleteACLsResponse{}, err
	}

	req := kmsg.NewDeleteACLsRequest()
	req.Filters = []kmsg.DeleteACLsRequestFilter{filter}

	res, err := req.RequestWith(ctx, cl)
	if err != nil {
		return DeleteACLsResponse{}, fmt.Errorf("Failed to execute delete topic command: %v", err.Error())
	}

	deleteAclsRes := DeleteACLsResponse{
		ErrorMessages: make([]string, 0),
		MatchedACLs:   0,
		DeletedACLs:   0,
	}
	for _, aclRes := range res.Results {
		err := newKafkaErrorWithDynamicMessage(aclRes.ErrorCode, aclRes.ErrorMessage)
		if err != nil {
			return DeleteACLsResponse{}, fmt.Errorf("Failed to delete Kafka ACL: %v", err.Error())
		}

		for _, item := range aclRes.MatchingACLs {
			deleteAclsRes.MatchedACLs++
			err := newKafkaErrorWithDynamicMessage(item.ErrorCode, item.ErrorMessage)
			if err != nil {
				deleteAclsRes.ErrorMessages = append(deleteAclsRes.ErrorMessages, err.Error())
				continue
			}
			deleteAclsRes.DeletedACLs++
		}
	}

	return deleteAclsRes, nil
}

// DeleteACLsKafka proxies the request/response via the Kafka API.
func (s *Service) DeleteACLsKafka(ctx context.Context, req *kmsg.DeleteACLsRequest) (*kmsg.DeleteACLsResponse, error) {
	cl, _, err := s.kafkaClientFactory.GetKafkaClient(ctx)
	if err != nil {
		return nil, err
	}
	return req.RequestWith(ctx, cl)
}

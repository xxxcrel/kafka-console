// Copyright 2022 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package kconsole

import (
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kerr"
	"github.com/twmb/franz-go/pkg/kmsg"
)

// CreateACL creates an ACL resource in your target Kafka cluster.
func (s *Service) CreateACL(ctx context.Context, createReq kmsg.CreateACLsRequestCreation) error {
	cl, _, err := s.kafkaClientFactory.GetKafkaClient(ctx)
	if err != nil {
		return err
	}

	req := kmsg.NewCreateACLsRequest()
	req.Creations = []kmsg.CreateACLsRequestCreation{createReq}
	res, err := req.RequestWith(ctx, cl)
	if err != nil {
		return fmt.Errorf("Failed to execute create ACL command: %v", err.Error())
	}

	if len(res.Results) != 1 {
		return fmt.Errorf("unexpected number of results in create ACL response")
	}

	aclRes := res.Results[0]
	err = kerr.ErrorForCode(aclRes.ErrorCode)
	if err != nil {
		return fmt.Errorf("Failed to execute create ACL command: %v", err.Error())
	}

	return nil
}

// CreateACLs proxies the request/response to CreateACLs via the Kafka API.
func (s *Service) CreateACLs(ctx context.Context, req *kmsg.CreateACLsRequest) (*kmsg.CreateACLsResponse, error) {
	cl, _, err := s.kafkaClientFactory.GetKafkaClient(ctx)
	if err != nil {
		return nil, err
	}
	return req.RequestWith(ctx, cl)
}

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

// DeleteTopic deletes a Kafka Topic (if possible and not disabled).
func (s *Service) DeleteTopic(ctx context.Context, topicName string) error {
	cl, _, err := s.kafkaClientFactory.GetKafkaClient(ctx)
	if err != nil {
		return err
	}

	req := kmsg.NewDeleteTopicsRequest()
	req.TopicNames = []string{topicName}
	req.TimeoutMillis = 30 * 1000 // 30s

	res, err := req.RequestWith(ctx, cl)
	if err != nil {
		return fmt.Errorf("Failed to execute delete topic command: %v", err.Error())
	}

	if len(res.Topics) != 1 {
		return fmt.Errorf("topics array in response is empty")
	}

	topicRes := res.Topics[0]
	if err := newKafkaErrorWithDynamicMessage(topicRes.ErrorCode, topicRes.ErrorMessage); err != nil {
		return fmt.Errorf("Failed to delete Kafka topic: %v", err.Error())
	}

	return nil
}

// DeleteTopics proxies the Kafka request/response between the Console service and Kafka.
func (s *Service) DeleteTopics(ctx context.Context, req *kmsg.DeleteTopicsRequest) (*kmsg.DeleteTopicsResponse, error) {
	cl, _, err := s.kafkaClientFactory.GetKafkaClient(ctx)
	if err != nil {
		return nil, err
	}

	return req.RequestWith(ctx, cl)
}

// DeleteTopicRecordsResponse is the response to deleting a Kafka topic.
type DeleteTopicRecordsResponse struct {
	TopicName  string                                `json:"topicName"`
	Partitions []DeleteTopicRecordsResponsePartition `json:"partitions"`
}

// DeleteTopicRecordsResponsePartition os the partition-scoped response to deleting
// a Kafka topic.
type DeleteTopicRecordsResponsePartition struct {
	PartitionID  int32  `json:"partitionId"`
	LowWaterMark int64  `json:"lowWaterMark"`
	ErrorMsg     string `json:"error,omitempty"`
}

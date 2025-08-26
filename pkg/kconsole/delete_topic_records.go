// Copyright 2024 Redpanda Data, Inc.
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

// DeleteTopicRecords deletes records within a Kafka topic until a certain offset.
func (s *Service) DeleteTopicRecords(ctx context.Context, deleteReq kmsg.DeleteRecordsRequestTopic) (DeleteTopicRecordsResponse, error) {
	cl, _, err := s.kafkaClientFactory.GetKafkaClient(ctx)
	if err != nil {
		return DeleteTopicRecordsResponse{}, err
	}

	req := kmsg.NewDeleteRecordsRequest()
	req.Topics = []kmsg.DeleteRecordsRequestTopic{deleteReq}
	res, err := req.RequestWith(ctx, cl)
	if err != nil {
		return DeleteTopicRecordsResponse{}, fmt.Errorf("Failed to execute delete topic command: %v", err.Error())
	}

	if len(res.Topics) != 1 {
		return DeleteTopicRecordsResponse{}, fmt.Errorf("topics array in response is empty")
	}

	topicRes := res.Topics[0]
	partitions := make([]DeleteTopicRecordsResponsePartition, len(topicRes.Partitions))
	for i, partitionRes := range topicRes.Partitions {
		err = kerr.ErrorForCode(partitionRes.ErrorCode)
		errMsg := ""
		if err != nil {
			errMsg = err.Error()
		}
		partitions[i] = DeleteTopicRecordsResponsePartition{
			PartitionID:  partitionRes.Partition,
			LowWaterMark: partitionRes.LowWatermark,
			ErrorMsg:     errMsg,
		}
	}

	return DeleteTopicRecordsResponse{
		TopicName:  topicRes.Topic,
		Partitions: partitions,
	}, nil
}

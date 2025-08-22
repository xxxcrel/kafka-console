// Copyright 2022 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package console

import (
	"context"

	"github.com/twmb/franz-go/pkg/kerr"
	"github.com/twmb/franz-go/pkg/kmsg"
)

// DeleteConsumerGroupOffsetsResponseTopic is the topic-scoped response to deleting
// a consumer group's offsets.
type DeleteConsumerGroupOffsetsResponseTopic struct {
	TopicName  string                                             `json:"topicName"`
	Partitions []DeleteConsumerGroupOffsetsResponseTopicPartition `json:"partitions"`
}

// DeleteConsumerGroupOffsetsResponseTopicPartition is the partition-scoped response to deleting
// a consumer group's offsets.
type DeleteConsumerGroupOffsetsResponseTopicPartition struct {
	ID    int32  `json:"partitionID"`
	Error string `json:"error,omitempty"`
}

// DeleteConsumerGroupOffsets requests to delete some or all consumer group's topic/partition offsets.
func (s *Service) DeleteConsumerGroupOffsets(ctx context.Context, groupID string, topics []kmsg.OffsetDeleteRequestTopic) ([]DeleteConsumerGroupOffsetsResponseTopic, error) {
	cl, _, err := s.kafkaClientFactory.GetKafkaClient(ctx)
	if err != nil {
		return nil, err
	}
	req := kmsg.NewOffsetDeleteRequest()
	req.Group = groupID
	req.Topics = topics

	offsetDeleteRes, err := req.RequestWith(ctx, cl)
	if err != nil {
		return nil, err
	}

	res := make([]DeleteConsumerGroupOffsetsResponseTopic, len(offsetDeleteRes.Topics))
	for i, topic := range offsetDeleteRes.Topics {
		partitions := make([]DeleteConsumerGroupOffsetsResponseTopicPartition, len(topic.Partitions))
		for j, partition := range topic.Partitions {
			err := kerr.ErrorForCode(partition.ErrorCode)
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			partitions[j] = DeleteConsumerGroupOffsetsResponseTopicPartition{
				ID:    partition.Partition,
				Error: errMsg,
			}
		}
		res[i] = DeleteConsumerGroupOffsetsResponseTopic{
			TopicName:  topic.Topic,
			Partitions: partitions,
		}
	}

	return res, nil
}

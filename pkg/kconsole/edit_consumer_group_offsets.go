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
	"strings"

	"github.com/twmb/franz-go/pkg/kerr"
	"github.com/twmb/franz-go/pkg/kmsg"
)

// EditConsumerGroupOffsetsResponse is the sum of all brokers' response shards for
// requesting a consumer group offset edit.
type EditConsumerGroupOffsetsResponse struct {
	Error  string                                  `json:"error,omitempty"`
	Topics []EditConsumerGroupOffsetsResponseTopic `json:"topics"`
}

// EditConsumerGroupOffsetsResponseTopic is the topic-scoped response to editing
// a consumer group's offset.
type EditConsumerGroupOffsetsResponseTopic struct {
	TopicName  string                                           `json:"topicName"`
	Partitions []EditConsumerGroupOffsetsResponseTopicPartition `json:"partitions"`
}

// EditConsumerGroupOffsetsResponseTopicPartition is the partition-scoped response to editing
// a consumer group's offset.
type EditConsumerGroupOffsetsResponseTopicPartition struct {
	ID    int32  `json:"partitionID"`
	Error string `json:"error,omitempty"`
}

// EditConsumerGroupOffsets edits the group offsets of one or more partitions.
//
//nolint:cyclop // Eventually this should be refactored to use the franz-go admin client
func (s *Service) EditConsumerGroupOffsets(ctx context.Context, groupID string, topics []kmsg.OffsetCommitRequestTopic) (*EditConsumerGroupOffsetsResponse, error) {
	cl, adminCl, err := s.kafkaClientFactory.GetKafkaClient(ctx)
	if err != nil {
		return nil, err
	}

	// 0. Check if consumer group is empty, otherwise we can't edit the group offsets and want to provide a proper
	// error message for the frontend.
	describedGroups, err := adminCl.DescribeGroups(ctx, groupID)
	if err != nil {
		return nil, fmt.Errorf("failed to check group state: %w", err)
	}
	describedGroup, exists := describedGroups[groupID]
	if !exists {
		return nil, fmt.Errorf("failed to check group state, group does not exist in response: %w", err)
	}

	if !strings.EqualFold(describedGroup.State, "empty") {
		return &EditConsumerGroupOffsetsResponse{
			Error:  fmt.Sprintf("Consumer group is still active and therefore can't be edited. Current Group State is: %v", describedGroup.State),
			Topics: nil,
		}, nil
	}

	// 1. The provided topic partitions might use special offsets (-2, -1) that need to be resolved to the earliest
	// or oldest offset before sending the edit group request to Kafka.
	topicPartitions := make(map[string][]int32)
	topicNames := make([]string, len(topics))
	for i, topic := range topics {
		topicNames[i] = topic.Topic
		for _, partition := range topic.Partitions {
			topicPartitions[topic.Topic] = append(topicPartitions[topic.Topic], partition.Partition)
		}
	}

	startOffsets, err := adminCl.ListStartOffsets(ctx, topicNames...)
	if err != nil {
		return nil, fmt.Errorf("Failed to list partition start offsets: %v", err.Error())
	}
	endOffsets, err := adminCl.ListEndOffsets(ctx, topicNames...)
	if err != nil {
		return nil, fmt.Errorf("Failed to list partition end offsets: %v", err.Error())
	}

	// Because topics is immutable and we want to replace special offsets
	// (earliest/oldest) with the actual watermarks we are effectively rebuilding
	// the offset commit request slice.
	substitutedTopics := make([]kmsg.OffsetCommitRequestTopic, len(topics))
	for i, topic := range topics {
		substitutedPartitions := make([]kmsg.OffsetCommitRequestTopicPartition, len(topic.Partitions))
		for j, partition := range topic.Partitions {
			switch partition.Partition {
			case TimestampLatest:
				offset, exists := endOffsets.Lookup(topic.Topic, partition.Partition)
				if !exists {
					return nil, fmt.Errorf("end offset for topic '%v' and partition %d is missing", topic.Topic, partition.Partition)
				}
				partition.Offset = offset.Offset
			case TimestampEarliest:
				offset, exists := startOffsets.Lookup(topic.Topic, partition.Partition)
				if !exists {
					return nil, fmt.Errorf("start offset for topic '%v' and partition %d is missing", topic.Topic, partition.Partition)
				}
				partition.Offset = offset.Offset
			default:
				// Nothing to edit, the provided offset is an exact match, and we don't need
				// to edit the oldest or earliest offset into it
			}
			substitutedPartitions[j] = partition
		}
		substitutedTopics[i] = kmsg.OffsetCommitRequestTopic{
			Topic:      topic.Topic,
			Partitions: substitutedPartitions,
		}
	}

	req := kmsg.NewOffsetCommitRequest()
	req.Group = groupID
	req.Topics = substitutedTopics
	commitResponse, err := req.RequestWith(ctx, cl)
	if err != nil {
		return nil, fmt.Errorf("Edit consumer group offsets failed: %v", err.Error())
	}

	editedTopics := make([]EditConsumerGroupOffsetsResponseTopic, len(commitResponse.Topics))
	for i, topic := range commitResponse.Topics {
		partitions := make([]EditConsumerGroupOffsetsResponseTopicPartition, len(topic.Partitions))
		for j, partition := range topic.Partitions {
			err := kerr.ErrorForCode(partition.ErrorCode)
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			partitions[j] = EditConsumerGroupOffsetsResponseTopicPartition{
				ID:    partition.Partition,
				Error: errMsg,
			}
		}
		editedTopics[i] = EditConsumerGroupOffsetsResponseTopic{
			TopicName:  topic.Topic,
			Partitions: partitions,
		}
	}

	return &EditConsumerGroupOffsetsResponse{
		Error:  "",
		Topics: editedTopics,
	}, nil
}

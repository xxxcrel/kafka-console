// Copyright 2022 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file https://github.com/xxxcrel/redpanda/blob/dev/licenses/bsl.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package console

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudhut/common/rest"
	"github.com/twmb/franz-go/pkg/kmsg"
)

// IncrementalAlterConfigsResourceResponse is the response to changing Kafka configurations
// via the Kafka API.
type IncrementalAlterConfigsResourceResponse struct {
	Error        string `json:"error,omitempty"`
	ResourceName string `json:"resourceName"`
	ResourceType int8   `json:"resourceType"`
}

// IncrementalAlterConfigs alters the configuration of a Kafka resource (broker/topic/...)
// via the Kafka API.
func (s *Service) IncrementalAlterConfigs(ctx context.Context,
	alterConfigs []kmsg.IncrementalAlterConfigsRequestResource,
) ([]IncrementalAlterConfigsResourceResponse, *rest.Error) {
	cl, _, err := s.kafkaClientFactory.GetKafkaClient(ctx)
	if err != nil {
		return nil, errorToRestError(err)
	}

	req := kmsg.NewIncrementalAlterConfigsRequest()
	req.Resources = alterConfigs

	configRes, err := req.RequestWith(ctx, cl)
	if err != nil {
		return nil, &rest.Error{
			Err:      err,
			Status:   http.StatusServiceUnavailable,
			Message:  fmt.Sprintf("Incremental Alter Config request has failed: %v", err.Error()),
			IsSilent: false,
		}
	}

	patchedConfigs := make([]IncrementalAlterConfigsResourceResponse, len(configRes.Resources))
	for i, res := range configRes.Resources {
		errMessage := ""
		kafkaErr := newKafkaErrorWithDynamicMessage(res.ErrorCode, res.ErrorMessage)
		if kafkaErr != nil {
			errMessage = err.Error()
		}
		patchedConfigs[i] = IncrementalAlterConfigsResourceResponse{
			Error:        errMessage,
			ResourceName: res.ResourceName,
			ResourceType: int8(res.ResourceType),
		}
	}

	return patchedConfigs, nil
}

// IncrementalAlterConfigsKafka alters the configuration of a Kafka resource (broker/topic/...)
// via the Kafka API. In contrast to IncrementalAlterConfigs the request and response re-uses the
// original Kafka client kmsg types and thus is only a proxy function which is used for abstracting
// and virtualizing the Console service.
func (s *Service) IncrementalAlterConfigsKafka(ctx context.Context, req *kmsg.IncrementalAlterConfigsRequest) (*kmsg.IncrementalAlterConfigsResponse, error) {
	cl, _, err := s.kafkaClientFactory.GetKafkaClient(ctx)
	if err != nil {
		return nil, err
	}

	return req.RequestWith(ctx, cl)
}

// AlterConfigs proxies the request/response to set configs (not incrementally) via the Kafka API. The difference
// between AlterConfigs and IncrementalAlterConfigs is that AlterConfigs sets the entire configuration so that
// all properties that are not set as part of this request will be reset to their default values.
func (s *Service) AlterConfigs(ctx context.Context, req *kmsg.AlterConfigsRequest) (*kmsg.AlterConfigsResponse, error) {
	cl, _, err := s.kafkaClientFactory.GetKafkaClient(ctx)
	if err != nil {
		return nil, err
	}

	return req.RequestWith(ctx, cl)
}

func errorToRestError(err error) *rest.Error {
	return &rest.Error{
		Err:      err,
		Status:   http.StatusServiceUnavailable,
		Message:  err.Error(),
		IsSilent: false,
	}
}

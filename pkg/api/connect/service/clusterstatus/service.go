// Copyright 2025 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

// Package clusterstatus implements RPCs that retrieve high level information
// providing insights about the health and deployed resources on all connected
// clusters and APIs.
package clusterstatus

import (
	"context"
	"fmt"
	"time"

	"connectrpc.com/connect"
	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kmsg"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	apierrors "github.com/xxxcrel/kafka-console/pkg/api/connect/errors"
	"github.com/xxxcrel/kafka-console/pkg/config"
	kafkaconnect "github.com/xxxcrel/kafka-console/pkg/connect"
	kafkafactory "github.com/xxxcrel/kafka-console/pkg/factory/kafka"
	"github.com/xxxcrel/kafka-console/pkg/factory/schema"
	"github.com/xxxcrel/kafka-console/pkg/version"
)

// Service that implements the ClusterStatusServiceHandler interface. This includes all
// RPCs to retrieve cluster information about all connected APIs.
type Service struct {
	cfg    *config.Config
	logger *zap.Logger

	kafkaClientProvider  kafkafactory.ClientFactory
	schemaClientProvider schema.ClientFactory
	connectSvc           *kafkaconnect.Service

	kafkaStatusChecker *kafkaStatusChecker
}

// NewService creates a new Service that serves the RPCs for retrieving cluster statuses.
func NewService(
	cfg *config.Config,
	logger *zap.Logger,
	kafkaClientProvider kafkafactory.ClientFactory,
	schemaClientProvider schema.ClientFactory,
	connectSvc *kafkaconnect.Service,
) *Service {
	return &Service{
		cfg:    cfg,
		logger: logger,

		kafkaClientProvider:  kafkaClientProvider,
		schemaClientProvider: schemaClientProvider,
		connectSvc:           connectSvc,

		kafkaStatusChecker: &kafkaStatusChecker{logger: logger},
	}
}

// GetKafkaInfo retrieves Kafka cluster metadata and API version concurrently,
// aggregates details (such as broker counts, topics, partitions, and replicas),
// and returns a comprehensive Kafka status response.
func (s *Service) GetKafkaInfo(ctx context.Context) (*KafkaInfo, error) {
	_, adminCl, err := s.kafkaClientProvider.GetKafkaClient(ctx)
	if err != nil {
		return nil, err
	}

	// We use a child context with a shorter timeout because otherwise we'll potentially have very long response
	// times in case of a single broker being down.
	childCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	grp, grpCtx := errgroup.WithContext(childCtx)

	// Fetch cluster metadata
	var metadata kadm.Metadata
	grp.Go(func() error {
		var err error
		metadata, err = adminCl.Metadata(grpCtx)
		return err
	})

	// Fetch Kafka API version
	clusterVersion := "unknown"
	grp.Go(func() error {
		var err error

		apiVersions, err := adminCl.ApiVersions(ctx)
		if err != nil {
			s.logger.Warn("failed to request kafka version", zap.Error(err))
		}
		apiVersions.Each(func(versions kadm.BrokerApiVersions) {
			if versions.Err != nil {
				s.logger.Warn("failed to request kafka version", zap.Int32("broker_id", versions.NodeID), zap.Error(versions.Err))
				return
			}
			clusterVersion = versions.VersionGuess()
		})

		return nil
	})

	if err := grp.Wait(); err != nil {
		return nil, apierrors.NewConnectError(
			connect.CodeInternal,
			err,
			apierrors.NewErrorInfo(Reason_REASON_KAFKA_API_ERROR.String(), apierrors.KeyValsFromKafkaError(err)...),
		)
	}

	partitionCount := int32(0)
	replicaCount := int32(0)
	metadata.Topics.EachPartition(func(p kadm.PartitionDetail) {
		if p.Err != nil {
			return
		}
		partitionCount++
		replicaCount += int32(len(p.Replicas))
	})

	kafkaInfoResponse := &KafkaInfo{
		Status:          s.kafkaStatusChecker.statusFromMetadata(metadata),
		Version:         clusterVersion,
		Distribution:    s.kafkaStatusChecker.distributionFromMetadata(metadata),
		BrokersOnline:   int32(len(metadata.Brokers)),
		BrokersExpected: int32(len(s.kafkaStatusChecker.getExpectedBrokers(metadata))),
		TopicsCount:     int32(len(metadata.Topics)),
		PartitionsCount: partitionCount,
		ReplicasCount:   replicaCount,
		ControllerId:    metadata.Controller,
		Brokers:         s.kafkaStatusChecker.brokersFromMetadata(metadata),
		ClusterId:       metadata.Cluster,
	}

	return kafkaInfoResponse, nil
}

// GetKafkaAuthorizerInfo fetches Kafka ACLs using a describe request and
// returns the total count of ACL resources, converting any Kafka API errors
// into ConnectRPC errors.
func (s *Service) GetKafkaAuthorizerInfo(ctx context.Context) (int32, error) {
	kafkaCl, _, err := s.kafkaClientProvider.GetKafkaClient(ctx)
	if err != nil {
		return 0, err
	}

	listAllReq := kmsg.NewDescribeACLsRequest()
	listAllReq.ResourcePatternType = kmsg.ACLResourcePatternTypeAny
	listAllReq.Operation = kmsg.ACLOperationAny
	listAllReq.PermissionType = kmsg.ACLPermissionTypeAny
	listAllReq.ResourceType = kmsg.ACLResourceTypeAny

	aclResponses, err := listAllReq.RequestWith(ctx, kafkaCl)
	if err != nil {
		return 0, apierrors.NewConnectError(
			connect.CodeInternal,
			err,
			apierrors.NewErrorInfo(Reason_REASON_KAFKA_API_ERROR.String(), apierrors.KeyValsFromKafkaError(err)...),
		)
	}

	connectErr := apierrors.NewConnectErrorFromKafkaErrorCode(aclResponses.ErrorCode, aclResponses.ErrorMessage)
	if connectErr != nil {
		return 0, connectErr
	}

	return int32(len(aclResponses.Resources)), nil
}

// GetConsoleInfo returns version and build timestamp information for Console.
func (*Service) GetConsoleInfo(context.Context) (string, string, error) {
	return version.Version, version.BuiltAt, nil
}

// GetKafkaConnectInfo retrieves information from all configured Kafka KafkaConnect
// clusters, including cluster name, host, version, health status, and the count
// of installed plugins, while ensuring that Kafka KafkaConnect is enabled.
func (s *Service) GetKafkaConnectInfo(ctx context.Context) ([]*KafkaConnectCluster, error) {
	// Currently the connectSvc is always configured, even if it's not enabled.
	// The connectSvc itself will then return errors if you request resources in
	// case it hasn't been enabled in the configuration. Hence, we have to check
	// whether the Kafka connect config is enabled.
	if s.connectSvc == nil || !s.connectSvc.Cfg.Enabled {
		return nil, apierrors.NewKafkaConnectNotConfiguredError()
	}

	// Get cluster info from all clusters
	clustersInfo := s.connectSvc.GetAllClusterInfo(ctx)
	clustersOverview := make([]*KafkaConnectCluster, len(clustersInfo))

	for i, clusterInfo := range clustersInfo {
		status := &ComponentStatus{Status: StatusType_STATUS_TYPE_HEALTHY}
		if clusterInfo.RequestError != nil {
			status = &ComponentStatus{Status: StatusType_STATUS_TYPE_UNHEALTHY, StatusReason: clusterInfo.RequestError.Error()}
		}
		clustersOverview[i] = &KafkaConnectCluster{
			Name:                  clusterInfo.Name,
			Status:                status,
			Host:                  clusterInfo.Host,
			Version:               clusterInfo.Version,
			InstalledPluginsCount: int32(len(clusterInfo.Plugins)),
		}
	}
	return clustersOverview, nil
}

// GetSchemaRegistryInfo obtains the status of the Schema Registry and the number
// of registered subjects. It reports an unhealthy status if subjects cannot be
// fetched, ensuring that errors are properly reflected in the response.
func (s *Service) GetSchemaRegistryInfo(ctx context.Context) (*SchemaRegistryInfo, error) {
	if !s.cfg.SchemaRegistry.Enabled {
		return nil, apierrors.NewSchemaRegistryNotConfiguredError()
	}

	status := &ComponentStatus{
		Status:       StatusType_STATUS_TYPE_HEALTHY,
		StatusReason: "",
	}

	srClient, err := s.schemaClientProvider.GetSchemaRegistryClient(ctx)
	if err != nil {
		return nil, err
	}

	registeredSubjects := int32(0)
	subjects, err := srClient.Subjects(ctx)
	if err != nil {
		setStatus(status, StatusType_STATUS_TYPE_UNHEALTHY, fmt.Sprintf("Could not fetch subjects from schema registry %q", err.Error()))
	} else {
		registeredSubjects = int32(len(subjects))
	}

	info := SchemaRegistryInfo{
		Status:                  status,
		RegisteredSubjectsCount: registeredSubjects,
	}

	return &info, nil
}

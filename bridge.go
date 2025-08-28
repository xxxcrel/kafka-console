package main

import (
	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/kmsg"
	"github.com/twmb/franz-go/pkg/sr"
	"github.com/xxxcrel/kafka-console/pkg/api/connect/service/clusterstatus"
	"github.com/xxxcrel/kafka-console/pkg/kconsole"
	"github.com/xxxcrel/kafka-console/pkg/serde"
)

func (app *App) GetClusterInfo() (*kconsole.ClusterInfo, error) {
	return app.api.ConsoleSvc.GetClusterInfo(app.ctx)
}

func (app *App) GetBrokersWithLogDirs() ([]kconsole.BrokerWithLogDirs, error) {
	return app.api.ConsoleSvc.GetBrokersWithLogDirs(app.ctx)
}

func (app *App) GetBrokerConfig(brokerId int32) ([]kconsole.BrokerConfigEntry, error) {
	return app.api.ConsoleSvc.GetBrokerConfig(app.ctx, brokerId)
}

func (app *App) GetEndpointCompatibility() (kconsole.EndpointCompatibility, error) {
	return app.api.ConsoleSvc.GetEndpointCompatibility(app.ctx)
}

func (app *App) GetConsumerGroupsOverview(groupIDS []string) ([]kconsole.ConsumerGroupOverview, error) {
	return app.api.ConsoleSvc.GetConsumerGroupsOverview(app.ctx, groupIDS)
}

func (app *App) GetKafkaInfo() (*clusterstatus.KafkaInfo, error) {
	return app.api.ClusterSvc.GetKafkaInfo(app.ctx)
}
func (app *App) GetKafkaAuthorizerInfo() (int32, error) {
	return app.api.ClusterSvc.GetKafkaAuthorizerInfo(app.ctx)
}

func (app *App) GetConsoleInfo() (string, string, error) {
	return app.api.ClusterSvc.GetConsoleInfo(app.ctx)
}

func (app *App) GetKafkaConnectInfo() ([]*clusterstatus.KafkaConnectCluster, error) {
	return app.api.ClusterSvc.GetKafkaConnectInfo(app.ctx)
}

func (app *App) GetSchemaRegistryInfo() (*clusterstatus.SchemaRegistryInfo, error) {
	return app.api.ClusterSvc.GetSchemaRegistryInfo(app.ctx)
}

func (app *App) GetAPIVersions() ([]kconsole.APIVersion, error) {
	return app.api.ConsoleSvc.GetAPIVersions(app.ctx)
}
func (app *App) GetAllBrokerConfigs() (map[int32]kconsole.BrokerConfig, error) {
	return app.api.ConsoleSvc.GetAllBrokerConfigs(app.ctx)
}
func (app *App) DeleteConsumerGroup(groupID string) error {
	return app.api.ConsoleSvc.DeleteConsumerGroup(app.ctx, groupID)
}
func (app *App) CreateACL(createReq kmsg.CreateACLsRequestCreation) error {
	return app.api.ConsoleSvc.CreateACL(app.ctx, createReq)
}
func (app *App) CreateTopic(createTopicReq kmsg.CreateTopicsRequestTopic) (kconsole.CreateTopicResponse, error) {
	return app.api.ConsoleSvc.CreateTopic(app.ctx, createTopicReq)
}
func (app *App) DeleteACLs(filter kmsg.DeleteACLsRequestFilter) (kconsole.DeleteACLsResponse, error) {
	return app.api.ConsoleSvc.DeleteACLs(app.ctx, filter)
}
func (app *App) DeleteConsumerGroupOffsets(groupID string, topics []kmsg.OffsetDeleteRequestTopic) ([]kconsole.DeleteConsumerGroupOffsetsResponseTopic, error) {
	return app.api.ConsoleSvc.DeleteConsumerGroupOffsets(app.ctx, groupID, topics)
}
func (app *App) DeleteTopic(topicName string) error {
	return app.api.ConsoleSvc.DeleteTopic(app.ctx, topicName)
}
func (app *App) DeleteTopicRecords(deleteReq kmsg.DeleteRecordsRequestTopic) (kconsole.DeleteTopicRecordsResponse, error) {
	return app.api.ConsoleSvc.DeleteTopicRecords(app.ctx, deleteReq)
}
func (app *App) DescribeQuotas() kconsole.QuotaResponse {
	return app.api.ConsoleSvc.DescribeQuotas(app.ctx)
}
func (app *App) EditConsumerGroupOffsets(groupID string, topics []kmsg.OffsetCommitRequestTopic) (*kconsole.EditConsumerGroupOffsetsResponse, error) {
	return app.api.ConsoleSvc.EditConsumerGroupOffsets(app.ctx, groupID, topics)
}
func (app *App) EditTopicConfig(topicName string, configs []kmsg.IncrementalAlterConfigsRequestResourceConfig) error {
	return app.api.ConsoleSvc.EditTopicConfig(app.ctx, topicName, configs)
}
func (app *App) IncrementalAlterConfigs(alterConfigs []kmsg.IncrementalAlterConfigsRequestResource) ([]kconsole.IncrementalAlterConfigsResourceResponse, error) {
	return app.api.ConsoleSvc.IncrementalAlterConfigs(app.ctx, alterConfigs)
}
func (app *App) ListAllACLs(req kmsg.DescribeACLsRequest) (*kconsole.ACLOverview, error) {
	return app.api.ConsoleSvc.ListAllACLs(app.ctx, req)
}
func (app *App) ListMessages(listReq kconsole.ListMessageRequest) ([]*kconsole.TopicMessage, error) {
	return app.api.ConsoleSvc.ListMessages(app.ctx, listReq)
}
func (app *App) ListOffsets(topicNames []string, timestamp int64) ([]kconsole.TopicOffset, error) {
	return app.api.ConsoleSvc.ListOffsets(app.ctx, topicNames, timestamp)
}
func (app *App) GetKafkaVersion() (string, error) {
	return app.api.ConsoleSvc.GetKafkaVersion(app.ctx)
}
func (app *App) ListPartitionReassignments() ([]kconsole.PartitionReassignments, error) {
	return app.api.ConsoleSvc.ListPartitionReassignments(app.ctx)
}
func (app *App) AlterPartitionAssignments(topics []kmsg.AlterPartitionAssignmentsRequestTopic) ([]kconsole.AlterPartitionReassignmentsResponse, error) {
	return app.api.ConsoleSvc.AlterPartitionAssignments(app.ctx, topics)
}
func (app *App) ProducePlainRecords(records []*kgo.Record, useTransactions bool, compressionOpts []kgo.CompressionCodec) *kconsole.ProduceRecordsResponse {
	return nil
}
func (app *App) ProduceRecord(string, int32, []kgo.RecordHeader, *serde.RecordPayloadInput, *serde.RecordPayloadInput, bool, []kgo.CompressionCodec) (*kconsole.ProduceRecordResponse, error) {
	return nil, nil
}
func (app *App) GetTopicConfigs(topicName string, configNames []string) (*kconsole.TopicConfig, error) {
	return app.api.ConsoleSvc.GetTopicConfigs(app.ctx, topicName, configNames)
}
func (app *App) GetTopicsConfigs(topicNames []string, configNames []string) (map[string]*kconsole.TopicConfig, error) {
	return app.api.ConsoleSvc.GetTopicsConfigs(app.ctx, topicNames, configNames)
}
func (app *App) ListTopicConsumers(topicName string) ([]*kconsole.TopicConsumerGroup, error) {
	return app.api.ConsoleSvc.ListTopicConsumers(app.ctx, topicName)
}
func (app *App) GetTopicDocumentation(topicName string) *kconsole.TopicDocumentation {
	return app.api.ConsoleSvc.GetTopicDocumentation(topicName)
}
func (app *App) GetTopicsOverview() ([]*kconsole.TopicSummary, error) {
	return app.api.ConsoleSvc.GetTopicsOverview(app.ctx)
}
func (app *App) GetAllTopicNames() ([]string, error) {
	return app.api.ConsoleSvc.GetAllTopicNames(app.ctx)
}
func (app *App) GetTopicDetails(topicNames []string) ([]kconsole.TopicDetails, error) {
	return app.api.ConsoleSvc.GetTopicDetails(app.ctx, topicNames)
}

// ------------------------------------------------------------------
// Plain Kafka requests, used by Connect API.
// The Console service was supposed to be a translation layer between the API (REST)
// and the Kafka package, but it's also used for virtualizing Console. Thus, even
// plain Kafka requests need to go through this package.
// ------------------------------------------------------------------

func (app *App) CreateACLs(createReq *kmsg.CreateACLsRequest) (*kmsg.CreateACLsResponse, error) {
	return app.api.ConsoleSvc.CreateACLs(app.ctx, createReq)
}
func (app *App) DeleteACLsKafka(deleteReq *kmsg.DeleteACLsRequest) (*kmsg.DeleteACLsResponse, error) {
	return app.api.ConsoleSvc.DeleteACLsKafka(app.ctx, deleteReq)
}
func (app *App) CreateTopics(createReq *kmsg.CreateTopicsRequest) (*kmsg.CreateTopicsResponse, error) {
	return app.api.ConsoleSvc.CreateTopics(app.ctx, createReq)
}
func (app *App) DescribeConfigs(req *kmsg.DescribeConfigsRequest) (*kmsg.DescribeConfigsResponse, error) {
	return app.api.ConsoleSvc.DescribeConfigs(app.ctx, req)
}
func (app *App) DeleteTopics(deleteReq *kmsg.DeleteTopicsRequest) (*kmsg.DeleteTopicsResponse, error) {
	return app.api.ConsoleSvc.DeleteTopics(app.ctx, deleteReq)
}
func (app *App) GetMetadata(metadataReq *kmsg.MetadataRequest) (*kmsg.MetadataResponse, error) {
	return app.api.ConsoleSvc.GetMetadata(app.ctx, metadataReq)
}
func (app *App) IncrementalAlterConfigsKafka(req *kmsg.IncrementalAlterConfigsRequest) (*kmsg.IncrementalAlterConfigsResponse, error) {
	return app.api.ConsoleSvc.IncrementalAlterConfigsKafka(app.ctx, req)
}
func (app *App) AlterConfigs(req *kmsg.AlterConfigsRequest) (*kmsg.AlterConfigsResponse, error) {
	return app.api.ConsoleSvc.AlterConfigs(app.ctx, req)
}
func (app *App) AddPartitionsToTopics(add int, topicNames []string, validateOnly bool) (kadm.CreatePartitionsResponses, error) {
	return app.api.ConsoleSvc.AddPartitionsToTopics(app.ctx, add, topicNames, validateOnly)
}
func (app *App) SetPartitionsToTopics(add int, topicNames []string, validateOnly bool) (kadm.CreatePartitionsResponses, error) {
	return app.api.ConsoleSvc.SetPartitionsToTopics(app.ctx, add, topicNames, validateOnly)
}

// SchemaRegistryServicer is the interface for schema registry servicer
func (app *App) GetSchemaRegistryMode() (*kconsole.SchemaRegistryMode, error) {
	return app.api.ConsoleSvc.GetSchemaRegistryMode(app.ctx)
}
func (app *App) GetSchemaRegistryConfig(subject string) (*kconsole.SchemaRegistryConfig, error) {
	return app.api.ConsoleSvc.GetSchemaRegistryConfig(app.ctx, subject)
}
func (app *App) PutSchemaRegistryConfig(subject string, compatibility sr.SetCompatibility) (*kconsole.SchemaRegistryConfig, error) {
	return app.api.ConsoleSvc.PutSchemaRegistryConfig(app.ctx, subject, compatibility)
}
func (app *App) DeleteSchemaRegistrySubjectConfig(subject string) error {
	return app.api.ConsoleSvc.DeleteSchemaRegistrySubjectConfig(app.ctx, subject)
}
func (app *App) GetSchemaRegistrySubjects() ([]kconsole.SchemaRegistrySubject, error) {
	return app.api.ConsoleSvc.GetSchemaRegistrySubjects(app.ctx)
}
func (app *App) GetSchemaRegistrySubjectDetails(subjectName string, version string) (*kconsole.SchemaRegistrySubjectDetails, error) {
	return app.api.ConsoleSvc.GetSchemaRegistrySubjectDetails(app.ctx, subjectName, version)
}
func (app *App) GetSchemaRegistrySchemaReferencedBy(subjectName string, version int) ([]kconsole.SchemaReference, error) {
	return app.api.ConsoleSvc.GetSchemaRegistrySchemaReferencedBy(app.ctx, subjectName, version)
}
func (app *App) DeleteSchemaRegistrySubject(subjectName string, deletePermanently bool) (*kconsole.SchemaRegistryDeleteSubjectResponse, error) {
	return app.api.ConsoleSvc.DeleteSchemaRegistrySubject(app.ctx, subjectName, deletePermanently)
}
func (app *App) DeleteSchemaRegistrySubjectVersion(subject string, version int, deletePermanently bool) (*kconsole.SchemaRegistryDeleteSubjectVersionResponse, error) {
	return app.api.ConsoleSvc.DeleteSchemaRegistrySubjectVersion(app.ctx, subject, version, deletePermanently)
}
func (app *App) GetSchemaRegistrySchemaTypes() (*kconsole.SchemaRegistrySchemaTypes, error) {
	return app.api.ConsoleSvc.GetSchemaRegistrySchemaTypes(app.ctx)
}
func (app *App) CreateSchemaRegistrySchema(subjectName string, schema sr.Schema) (*kconsole.CreateSchemaResponse, error) {
	return app.api.ConsoleSvc.CreateSchemaRegistrySchema(app.ctx, subjectName, schema)
}
func (app *App) ValidateSchemaRegistrySchema(subjectName string, version int, schema sr.Schema) (*kconsole.SchemaRegistrySchemaValidation, error) {
	return app.api.ConsoleSvc.ValidateSchemaRegistrySchema(app.ctx, subjectName, version, schema)
}
func (app *App) GetSchemaUsagesByID(schemaID int) ([]kconsole.SchemaVersion, error) {
	return app.api.ConsoleSvc.GetSchemaUsagesByID(app.ctx, schemaID)
}

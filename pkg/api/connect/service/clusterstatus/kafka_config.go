package clusterstatus

// KafkaBroker represents details about a single Kafka broker.
type KafkaBroker struct {
	// BrokerID is the unique identifier of the broker.
	BrokerId int32 `json:"brokerId"`
	// Host is the address or hostname of the broker.
	Host string `json:"host"`
	// RackID (availability zone) configured for this broker.
	RackId *string `json:"rackId"`
}

type StatusType int32

const (
	// STATUS_TYPE_UNSPECIFIED is the default value.
	StatusType_STATUS_TYPE_UNSPECIFIED StatusType = 0
	// STATUS_TYPE_HEALTHY indicates the component is healthy.
	StatusType_STATUS_TYPE_HEALTHY StatusType = 1
	// STATUS_TYPE_DEGRADED indicates the component is partially impaired.
	StatusType_STATUS_TYPE_DEGRADED StatusType = 2
	// STATUS_TYPE_UNHEALTHY indicates the component is unhealthy or unreachable.
	StatusType_STATUS_TYPE_UNHEALTHY StatusType = 3
)

type ComponentStatus struct {
	// Status is the overall health.
	Status StatusType `json:"status"`
	// StatusReason provides details if the component is not healthy.
	StatusReason string `json:"statusReason"`
}

// KafkaDistribution describes the Kafka software flavor.
type KafkaDistribution int32

const (
	// KAFKA_DISTRIBUTION_UNSPECIFIED is the default value.
	KafkaDistribution_KAFKA_DISTRIBUTION_UNSPECIFIED KafkaDistribution = 0
	// KAFKA_DISTRIBUTION_UNKNOWN represents a Kafka distribution that is neither Redpanda
	// nor Apache Kafka.
	KafkaDistribution_KAFKA_DISTRIBUTION_UNKNOWN KafkaDistribution = 1
	// KAFKA_DISTRIBUTION_APACHE_KAFKA indicates an Apache Kafka implementation.
	KafkaDistribution_KAFKA_DISTRIBUTION_APACHE_KAFKA KafkaDistribution = 2
)

type KafkaInfo struct {
	// The health status for Kafka.
	Status *ComponentStatus `json:"status,omitempty"`
	// Version is the Kafka API version.
	Version string `json:"version, omitempty"`
	// Distribution indicates the software flavor (e.g. Apache Kafka, Redpanda).
	Distribution KafkaDistribution `json:"distribution,omitempty"`
	// BrokersOnline is the number of brokers currently online.
	BrokersOnline int32 `json:"brokersOnline,omitempty"`
	// BrokersExpected is the number of brokers expected.
	BrokersExpected int32 `json:"brokersExpected,omitempty"`
	// TopicsCount is the total number of topics.
	TopicsCount int32 `json:"topicsCount,omitempty"`
	// PartitionsCount is the total number of partitions.
	PartitionsCount int32 `json:"partitionsCount,omitempty"`
	// ReplicasCount is the total number of replicas.
	ReplicasCount int32 `json:"replicasCount,omitempty"`
	// ControllerID is the ID of the controller broker.
	ControllerId int32 `json:"controllerId,omitempty"`
	// Brokers lists details of individual Kafka brokers.
	Brokers []*KafkaBroker `json:"brokers,omitempty"`
	// ClusterID is the unique identifier for the Kafka cluster.
	ClusterId string `json:"clusterId,omitempty"`
}

// KafkaConnectCluster represents a single Kafka Connect cluster.
type KafkaConnectCluster struct {
	// Name is the Kafka Connect cluster name.
	Name string `json:"name,omitempty"`
	// Status indicates the health of the cluster.
	Status *ComponentStatus `json:"status,omitempty"`
	// Host is the address of the Kafka Connect cluster.
	Host string `json:"host,omitempty"`
	// Version is the version of the Kafka Connect cluster.
	Version string `json:"version,omitempty"`
	// InstalledPlugins is the number of plugins installed.
	InstalledPluginsCount int32 `json:"installed_plugins_count,omitempty"`
}

type SchemaRegistryInfo struct {
	// Status indicates the health status of the Schema Registry.
	Status *ComponentStatus `json:"status,omitempty"`
	// RegisteredSubjectsCount is the number of subjects registered.
	RegisteredSubjectsCount int32 `json:"registered_subjects_count,omitempty"`
}

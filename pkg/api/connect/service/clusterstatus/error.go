package clusterstatus

type Reason int32

const (
	Reason_REASON_UNSPECIFIED Reason = 0
	// The feature is not configured.
	Reason_REASON_FEATURE_NOT_CONFIGURED Reason = 1
	// Internal Redpanda Console or data plane error.
	Reason_REASON_CONSOLE_ERROR Reason = 2
	// Redpanda Admin API returned an error.
	Reason_REASON_REDPANDA_ADMIN_API_ERROR Reason = 3
	// Redpanda or Kafka protocol error.
	Reason_REASON_KAFKA_API_ERROR Reason = 4
	// Kafka Connect API error.
	Reason_REASON_KAFKA_CONNECT_API_ERROR Reason = 5
	// Type mapping error translating internal or external types to API types.
	Reason_REASON_TYPE_MAPPING_ERROR Reason = 6
	// Cloud provider's secret store manager error.
	Reason_REASON_SECRET_STORE_ERROR Reason = 7
	// Invalid pipeline configuration.
	Reason_REASON_CONNECT_INVALID_PIPELINE_CONFIGURATION Reason = 8
	// The Redpanda enterprise license has expired and is no longer valid.
	Reason_REASON_ENTERPRISE_LICENSE_EXPIRED Reason = 9
)

// Enum value maps for Reason.
var (
	Reason_name = map[int32]string{
		0: "REASON_UNSPECIFIED",
		1: "REASON_FEATURE_NOT_CONFIGURED",
		2: "REASON_CONSOLE_ERROR",
		3: "REASON_REDPANDA_ADMIN_API_ERROR",
		4: "REASON_KAFKA_API_ERROR",
		5: "REASON_KAFKA_CONNECT_API_ERROR",
		6: "REASON_TYPE_MAPPING_ERROR",
		7: "REASON_SECRET_STORE_ERROR",
		8: "REASON_CONNECT_INVALID_PIPELINE_CONFIGURATION",
		9: "REASON_ENTERPRISE_LICENSE_EXPIRED",
	}
	Reason_value = map[string]int32{
		"REASON_UNSPECIFIED":                            0,
		"REASON_FEATURE_NOT_CONFIGURED":                 1,
		"REASON_CONSOLE_ERROR":                          2,
		"REASON_REDPANDA_ADMIN_API_ERROR":               3,
		"REASON_KAFKA_API_ERROR":                        4,
		"REASON_KAFKA_CONNECT_API_ERROR":                5,
		"REASON_TYPE_MAPPING_ERROR":                     6,
		"REASON_SECRET_STORE_ERROR":                     7,
		"REASON_CONNECT_INVALID_PIPELINE_CONFIGURATION": 8,
		"REASON_ENTERPRISE_LICENSE_EXPIRED":             9,
	}
)

func (x Reason) Enum() *Reason {
	p := new(Reason)
	*p = x
	return p
}

func (x Reason) String() string {
	return Reason_name[int32(x)]
}

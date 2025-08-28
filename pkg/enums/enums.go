package enums

import (
	"github.com/twmb/franz-go/pkg/kmsg"
	"github.com/twmb/franz-go/pkg/sr"
	"github.com/xxxcrel/kafka-console/pkg/api/connect/service/console"
	"github.com/xxxcrel/kafka-console/pkg/kconsole"
)

var AllSchemaType = []struct {
	Value  sr.SchemaType
	TSName string
}{
	{sr.TypeAvro, "AVRO"},
	{sr.TypeProtobuf, "PROTOBUF"},
	{sr.TypeJSON, "JSON"},
}

var AllCompatibilityLevel = []struct {
	Value  sr.CompatibilityLevel
	TSName string
}{
	{sr.CompatNone, "NONE"},
	{sr.CompatBackward, "BACKWROD"},
	{sr.CompatBackwardTransitive, "BACKWORD_TRANSITIVE"},
	{sr.CompatForward, "FORWARD"},
	{sr.CompatForwardTransitive, "FORWARD_TRANSITIVE"},
	{sr.CompatFull, "FULL"},
	{sr.CompatFullTransitive, "FULL_TRANSITIVE"},
}

var AllMode = []struct {
	Value  sr.Mode
	TSName string
}{
	{sr.ModeImport, "IMPORT"},
	{sr.ModeReadOnly, "READONLY"},
	{sr.ModeReadWrite, "READWRITE"},
}

// 所有 SchemaRuleKind 枚举选项
var AllSchemaRuleKind = []struct {
	Value  sr.SchemaRuleKind
	TSName string
}{
	{sr.SchemaRuleKindTransform, "TRANSFORM"},
	{sr.SchemaRuleKindCondition, "CONDITION"},
}

// 所有 SchemaRuleMode 枚举选项
var AllSchemaRuleMode = []struct {
	Value  sr.SchemaRuleMode
	TSName string
}{
	{sr.SchemaRuleModeUpgrade, "UPGRADE"},
	{sr.SchemaRuleModeDowngrade, "DOWNGRADE"},
	{sr.SchemaRuleModeUpdown, "UPDOWN"},
	{sr.SchemaRuleModeWrite, "WRITE"},
	{sr.SchemaRuleModeRead, "READ"},
	{sr.SchemaRuleModeWriteRead, "WRITEREAD"},
}

var AllACLResourceType = []struct {
	Value  kmsg.ACLResourceType
	TSName string
}{
	{kmsg.ACLResourceTypeUnknown, "UNKNOWN"},
	{kmsg.ACLResourceTypeAny, "ANY"},
	{kmsg.ACLResourceTypeTopic, "TOPIC"},
	{kmsg.ACLResourceTypeGroup, "GROUP"},
	{kmsg.ACLResourceTypeCluster, "CLUSTER"},
	{kmsg.ACLResourceTypeTransactionalId, "TRANSACTIONAL_ID"},
	{kmsg.ACLResourceTypeDelegationToken, "DELEGATION_TOKEN"},
	{kmsg.ACLResourceTypeUser, "USER"},
}
var AllACLResourcePatternType = []struct {
	Value  kmsg.ACLResourcePatternType
	TSName string
}{
	{kmsg.ACLResourcePatternTypeUnknown, "UNKNOWN"},
	{kmsg.ACLResourcePatternTypeAny, "ANY"},
	{kmsg.ACLResourcePatternTypeMatch, "MATCH"},
	{kmsg.ACLResourcePatternTypeLiteral, "LITERAL"},
	{kmsg.ACLResourcePatternTypePrefixed, "PREFIXED"},
}
var AllACLPermissionType = []struct {
	Value  kmsg.ACLPermissionType
	TSName string
}{
	{kmsg.ACLPermissionTypeUnknown, "UNKNOWN"},
	{kmsg.ACLPermissionTypeAny, "ANY"},
	{kmsg.ACLPermissionTypeDeny, "DENY"},
	{kmsg.ACLPermissionTypeAllow, "ALLOW"},
}

var AllACLOperation = []struct {
	Value  kmsg.ACLOperation
	TSName string
}{
	{kmsg.ACLOperationUnknown, "UNKNOWN"},
	{kmsg.ACLOperationAny, "ANY"},
	{kmsg.ACLOperationAll, "ALL"},
	{kmsg.ACLOperationRead, "READ"},
	{kmsg.ACLOperationWrite, "WRITE"},
	{kmsg.ACLOperationCreate, "CREATE"},
	{kmsg.ACLOperationDelete, "DELETE"},
	{kmsg.ACLOperationAlter, "ALTER"},
	{kmsg.ACLOperationDescribe, "DESCRIBE"},
	{kmsg.ACLOperationClusterAction, "CLUSTER_ACTION"},
	{kmsg.ACLOperationDescribeConfigs, "DESCRIBE_CONFIGS"},
	{kmsg.ACLOperationAlterConfigs, "ALTER_CONFIGS"},
	{kmsg.ACLOperationIdempotentWrite, "IDEMPOTENT_WRITE"},
	{kmsg.ACLOperationCreateTokens, "CREATE_TOKENS"},
	{kmsg.ACLOperationDescribeTokens, "DESCRIBE_TOKENS"},
}

var AllFrontendFormat = []struct {
	Value  kconsole.FrontendFormat
	TSName string
}{
	{kconsole.FrontendFormatUnknown, "UNKNOWN"},
	{kconsole.FrontendFormatBoolean, "BOOLEAN"},
	{kconsole.FrontendFormatPassword, "PASSWORD"},
	{kconsole.FrontendFormatString, "STRING"},
	{kconsole.FrontendFormatSelect, "SELECT"},
	{kconsole.FrontendFormatMultiSelect, "MULTI_SELECT"},
	{kconsole.FrontendFormatByteSize, "BYTE_SIZE"},
	{kconsole.FrontendFormatRatio, "RATIO"},
	{kconsole.FrontendFormatDuration, "DURATION"},
	{kconsole.FrontendFormatDecimal, "DECIMAL"},
	{kconsole.FrontendFormatInteger, "INTEGER"},
}

var AllConfigType = []struct {
	Value  kmsg.ConfigType
	TSName string
}{
	{kmsg.ConfigType(1), "BOOLEAN"},
	{kmsg.ConfigType(2), "STRING"},
	{kmsg.ConfigType(3), "INT"},
	{kmsg.ConfigType(4), "SHORT"},
	{kmsg.ConfigType(5), "LONG"},
	{kmsg.ConfigType(6), "DOUBLE"},
	{kmsg.ConfigType(7), "LIST"},
	{kmsg.ConfigType(8), "CLASS"},
	{kmsg.ConfigType(9), "PASSWORD"},
}

var AllIncrementalAlterConfigOp = []struct {
	Value  kmsg.IncrementalAlterConfigOp
	TSName string
}{
	{kmsg.IncrementalAlterConfigOpSet, "SET"},
	{kmsg.IncrementalAlterConfigOpDelete, "DELETE"},
	{kmsg.IncrementalAlterConfigOpAppend, "APPEND"},
	{kmsg.IncrementalAlterConfigOpSubtract, "SUBTRACT"},
}

var AllConfigResourceType = []struct {
	Value  kmsg.ConfigResourceType
	TSName string
}{
	{kmsg.ConfigResourceTypeUnknown, "UNKNOWN"},
	{kmsg.ConfigResourceTypeTopic, "TOPIC"},
	{kmsg.ConfigResourceTypeBroker, "BROKER"},
	{kmsg.ConfigResourceTypeBrokerLogger, "BROKER_LOGGER"},
}

var AllPayloadEncoding = []struct {
	Value  console.PayloadEncoding
	TSName string
}{
	{console.PayloadEncoding_PAYLOAD_ENCODING_UNSPECIFIED, "UNSPECIFIED"},
	{console.PayloadEncoding_PAYLOAD_ENCODING_NULL, "NULL"},
	{console.PayloadEncoding_PAYLOAD_ENCODING_AVRO, "AVRO"},
	{console.PayloadEncoding_PAYLOAD_ENCODING_PROTOBUF, "PROTOBUF"},
	{console.PayloadEncoding_PAYLOAD_ENCODING_PROTOBUF_SCHEMA, "PROTOBUF_SCHEMA"},
	{console.PayloadEncoding_PAYLOAD_ENCODING_JSON, "JSON"},
	{console.PayloadEncoding_PAYLOAD_ENCODING_JSON_SCHEMA, "JSON_SCHEMA"},
	{console.PayloadEncoding_PAYLOAD_ENCODING_XML, "XML"},
	{console.PayloadEncoding_PAYLOAD_ENCODING_TEXT, "TEXT"},
	{console.PayloadEncoding_PAYLOAD_ENCODING_UTF8, "UTF8"},
	{console.PayloadEncoding_PAYLOAD_ENCODING_MESSAGE_PACK, "MESSAGE_PACK"},
	{console.PayloadEncoding_PAYLOAD_ENCODING_SMILE, "SMILE"},
	{console.PayloadEncoding_PAYLOAD_ENCODING_BINARY, "BINARY"},
	{console.PayloadEncoding_PAYLOAD_ENCODING_UINT, "UINT"},
	{console.PayloadEncoding_PAYLOAD_ENCODING_CONSUMER_OFFSETS, "CONSUMER_OFFSETS"},
	{console.PayloadEncoding_PAYLOAD_ENCODING_CBOR, "CBOR"},
}

var AllCompressionType = []struct {
	Value  console.CompressionType
	TSName string
}{
	{console.CompressionType_COMPRESSION_TYPE_UNSPECIFIED, "UNSPECIFIED"},
	{console.CompressionType_COMPRESSION_TYPE_UNCOMPRESSED, "UNCOMPRESSED"},
	{console.CompressionType_COMPRESSION_TYPE_GZIP, "GZIP"},
	{console.CompressionType_COMPRESSION_TYPE_SNAPPY, "SNAPPY"},
	{console.CompressionType_COMPRESSION_TYPE_LZ4, "LZ4"},
	{console.CompressionType_COMPRESSION_TYPE_ZSTD, "ZSTD"},
}

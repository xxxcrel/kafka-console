package console

type PayloadEncoding int32

const (
	PayloadEncoding_PAYLOAD_ENCODING_UNSPECIFIED      PayloadEncoding = 0
	PayloadEncoding_PAYLOAD_ENCODING_NULL             PayloadEncoding = 1
	PayloadEncoding_PAYLOAD_ENCODING_AVRO             PayloadEncoding = 2
	PayloadEncoding_PAYLOAD_ENCODING_PROTOBUF         PayloadEncoding = 3
	PayloadEncoding_PAYLOAD_ENCODING_PROTOBUF_SCHEMA  PayloadEncoding = 4
	PayloadEncoding_PAYLOAD_ENCODING_JSON             PayloadEncoding = 5
	PayloadEncoding_PAYLOAD_ENCODING_JSON_SCHEMA      PayloadEncoding = 6
	PayloadEncoding_PAYLOAD_ENCODING_XML              PayloadEncoding = 7
	PayloadEncoding_PAYLOAD_ENCODING_TEXT             PayloadEncoding = 8
	PayloadEncoding_PAYLOAD_ENCODING_UTF8             PayloadEncoding = 9
	PayloadEncoding_PAYLOAD_ENCODING_MESSAGE_PACK     PayloadEncoding = 10
	PayloadEncoding_PAYLOAD_ENCODING_SMILE            PayloadEncoding = 11
	PayloadEncoding_PAYLOAD_ENCODING_BINARY           PayloadEncoding = 12
	PayloadEncoding_PAYLOAD_ENCODING_UINT             PayloadEncoding = 13
	PayloadEncoding_PAYLOAD_ENCODING_CONSUMER_OFFSETS PayloadEncoding = 14
	PayloadEncoding_PAYLOAD_ENCODING_CBOR             PayloadEncoding = 15
)

// Enum value maps for PayloadEncoding.
var (
	PayloadEncoding_name = map[int32]string{
		0:  "PAYLOAD_ENCODING_UNSPECIFIED",
		1:  "PAYLOAD_ENCODING_NULL",
		2:  "PAYLOAD_ENCODING_AVRO",
		3:  "PAYLOAD_ENCODING_PROTOBUF",
		4:  "PAYLOAD_ENCODING_PROTOBUF_SCHEMA",
		5:  "PAYLOAD_ENCODING_JSON",
		6:  "PAYLOAD_ENCODING_JSON_SCHEMA",
		7:  "PAYLOAD_ENCODING_XML",
		8:  "PAYLOAD_ENCODING_TEXT",
		9:  "PAYLOAD_ENCODING_UTF8",
		10: "PAYLOAD_ENCODING_MESSAGE_PACK",
		11: "PAYLOAD_ENCODING_SMILE",
		12: "PAYLOAD_ENCODING_BINARY",
		13: "PAYLOAD_ENCODING_UINT",
		14: "PAYLOAD_ENCODING_CONSUMER_OFFSETS",
		15: "PAYLOAD_ENCODING_CBOR",
	}
	PayloadEncoding_value = map[string]int32{
		"PAYLOAD_ENCODING_UNSPECIFIED":      0,
		"PAYLOAD_ENCODING_NULL":             1,
		"PAYLOAD_ENCODING_AVRO":             2,
		"PAYLOAD_ENCODING_PROTOBUF":         3,
		"PAYLOAD_ENCODING_PROTOBUF_SCHEMA":  4,
		"PAYLOAD_ENCODING_JSON":             5,
		"PAYLOAD_ENCODING_JSON_SCHEMA":      6,
		"PAYLOAD_ENCODING_XML":              7,
		"PAYLOAD_ENCODING_TEXT":             8,
		"PAYLOAD_ENCODING_UTF8":             9,
		"PAYLOAD_ENCODING_MESSAGE_PACK":     10,
		"PAYLOAD_ENCODING_SMILE":            11,
		"PAYLOAD_ENCODING_BINARY":           12,
		"PAYLOAD_ENCODING_UINT":             13,
		"PAYLOAD_ENCODING_CONSUMER_OFFSETS": 14,
		"PAYLOAD_ENCODING_CBOR":             15,
	}
)

func (x PayloadEncoding) Enum() *PayloadEncoding {
	p := new(PayloadEncoding)
	*p = x
	return p
}

func (x PayloadEncoding) String() string {
	return PayloadEncoding_name[int32(x)]
}

type ListMessagesRequest struct {
	Topic                     string           `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`                                                                                                              // Topic name.
	StartOffset               int64            `protobuf:"zigzag64,2,opt,name=start_offset,json=startOffset,proto3" json:"start_offset,omitempty"`                                                                            // Start offset. -1 for recent (newest - results), -2 for oldest offset, -3 for newest, -4 for timestamp.
	StartTimestamp            int64            `protobuf:"varint,3,opt,name=start_timestamp,json=startTimestamp,proto3" json:"start_timestamp,omitempty"`                                                                     // Start offset by unix timestamp in ms (only considered if start offset is set to -4).
	PartitionId               int32            `protobuf:"varint,4,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`                                                                              // -1 for all partition ids
	MaxResults                int32            `protobuf:"varint,5,opt,name=max_results,json=maxResults,proto3" json:"max_results,omitempty"`                                                                                 // Maximum number of results
	FilterInterpreterCode     string           `protobuf:"bytes,6,opt,name=filter_interpreter_code,json=filterInterpreterCode,proto3" json:"filter_interpreter_code,omitempty"`                                               // Base64 encoded code
	Enterprise                []byte           `protobuf:"bytes,7,opt,name=enterprise,proto3" json:"enterprise,omitempty"`                                                                                                    // Enterprise may only be set in the Enterprise mode. The JSON deserialization is deferred.
	Troubleshoot              bool             `protobuf:"varint,8,opt,name=troubleshoot,proto3" json:"troubleshoot,omitempty"`                                                                                               // Optionally include troubleshooting data in the response.
	IncludeOriginalRawPayload bool             `protobuf:"varint,9,opt,name=include_original_raw_payload,json=includeOriginalRawPayload,proto3" json:"include_original_raw_payload,omitempty"`                                // Optionally include original raw payload.
	KeyDeserializer           *PayloadEncoding `protobuf:"varint,10,opt,name=key_deserializer,json=keyDeserializer,proto3,enum=redpanda.api.kconsole.v1alpha1.PayloadEncoding,oneof" json:"key_deserializer,omitempty"`       // Optionally specify key payload deserialization strategy to use.
	ValueDeserializer         *PayloadEncoding `protobuf:"varint,11,opt,name=value_deserializer,json=valueDeserializer,proto3,enum=redpanda.api.kconsole.v1alpha1.PayloadEncoding,oneof" json:"value_deserializer,omitempty"` // Optionally specify value payload deserialization strategy to use.
	IgnoreMaxSizeLimit        bool             `protobuf:"varint,12,opt,name=ignore_max_size_limit,json=ignoreMaxSizeLimit,proto3" json:"ignore_max_size_limit,omitempty"`                                                    // Optionally ignore configured maximum payload size limit.
}

// KafkaRecordHeader is the record header.
type KafkaRecordHeader struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`     // Header key.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"` // Header value.
}

// PublishMessageRequest is the request for PublishMessage call.
type PublishMessageRequest struct {
	Topic           string                        `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`                                                                  // The topics to publish to.
	PartitionId     int32                         `protobuf:"varint,2,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`                                  // -1 for automatic partition assignment.
	Compression     CompressionType               `protobuf:"varint,3,opt,name=compression,proto3,enum=redpanda.api.kconsole.v1alpha1.CompressionType" json:"compression,omitempty"` // The compression to be used.
	UseTransactions bool                          `protobuf:"varint,4,opt,name=use_transactions,json=useTransactions,proto3" json:"use_transactions,omitempty"`                      // Use transactions.
	Headers         []*KafkaRecordHeader          `protobuf:"bytes,5,rep,name=headers,proto3" json:"headers,omitempty"`                                                              // Kafka record headers.
	Key             *PublishMessagePayloadOptions `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
	Value           *PublishMessagePayloadOptions `protobuf:"bytes,7,opt,name=value,proto3" json:"value,omitempty"`
}
type PublishMessagePayloadOptions struct {
	Encoding PayloadEncoding `protobuf:"varint,1,opt,name=encoding,proto3,enum=redpanda.api.kconsole.v1alpha1.PayloadEncoding" json:"encoding,omitempty"` // Payload encoding to use.
	Data     []byte          `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`                                                              // Data.
	SchemaId *int32          `json:"schema_id,omitempty"`                                                                                                 // Optional schema ID.
	Index    *int32          `protobuf:"varint,10,opt,name=index,proto3,oneof" json:"index,omitempty"`                                                    // Optional index. Useful for Protobuf messages.
}

// PublishMessageResponse is the response for PublishMessage call.
type PublishMessageResponse struct {
	Topic       string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	PartitionId int32  `protobuf:"varint,2,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`
	Offset      int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

type CompressionType int32

const (
	CompressionType_COMPRESSION_TYPE_UNSPECIFIED  CompressionType = 0
	CompressionType_COMPRESSION_TYPE_UNCOMPRESSED CompressionType = 1
	CompressionType_COMPRESSION_TYPE_GZIP         CompressionType = 2
	CompressionType_COMPRESSION_TYPE_SNAPPY       CompressionType = 3
	CompressionType_COMPRESSION_TYPE_LZ4          CompressionType = 4
	CompressionType_COMPRESSION_TYPE_ZSTD         CompressionType = 5
)

// Enum value maps for CompressionType.
var (
	CompressionType_name = map[int32]string{
		0: "COMPRESSION_TYPE_UNSPECIFIED",
		1: "COMPRESSION_TYPE_UNCOMPRESSED",
		2: "COMPRESSION_TYPE_GZIP",
		3: "COMPRESSION_TYPE_SNAPPY",
		4: "COMPRESSION_TYPE_LZ4",
		5: "COMPRESSION_TYPE_ZSTD",
	}
	CompressionType_value = map[string]int32{
		"COMPRESSION_TYPE_UNSPECIFIED":  0,
		"COMPRESSION_TYPE_UNCOMPRESSED": 1,
		"COMPRESSION_TYPE_GZIP":         2,
		"COMPRESSION_TYPE_SNAPPY":       3,
		"COMPRESSION_TYPE_LZ4":          4,
		"COMPRESSION_TYPE_ZSTD":         5,
	}
)

func (x CompressionType) Enum() *CompressionType {
	p := new(CompressionType)
	*p = x
	return p
}

func (x CompressionType) String() string {
	return CompressionType_name[int32(x)]
}

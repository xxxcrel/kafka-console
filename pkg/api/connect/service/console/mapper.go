// Copyright 2023 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package console

import (
	"github.com/twmb/franz-go/pkg/kgo"

	"github.com/xxxcrel/kafka-console/pkg/serde"
)

func rpcPublishMessagePayloadOptionsToSerializeInput(po *PublishMessagePayloadOptions) *serde.RecordPayloadInput { //nolint:cyclop // we have to map all possible values here
	encoding := serde.PayloadEncodingBinary

	switch po.Encoding {
	case PayloadEncoding_PAYLOAD_ENCODING_NULL:
		encoding = serde.PayloadEncodingNull
	case PayloadEncoding_PAYLOAD_ENCODING_AVRO:
		encoding = serde.PayloadEncodingAvro
	case PayloadEncoding_PAYLOAD_ENCODING_PROTOBUF:
		encoding = serde.PayloadEncodingProtobuf
	case PayloadEncoding_PAYLOAD_ENCODING_PROTOBUF_SCHEMA:
		encoding = serde.PayloadEncodingProtobufSchema
	case PayloadEncoding_PAYLOAD_ENCODING_JSON:
		encoding = serde.PayloadEncodingJSON
	case PayloadEncoding_PAYLOAD_ENCODING_JSON_SCHEMA:
		encoding = serde.PayloadEncodingJSONSchema
	case PayloadEncoding_PAYLOAD_ENCODING_XML:
		encoding = serde.PayloadEncodingXML
	case PayloadEncoding_PAYLOAD_ENCODING_TEXT:
		encoding = serde.PayloadEncodingText
	case PayloadEncoding_PAYLOAD_ENCODING_UTF8:
		encoding = serde.PayloadEncodingUtf8WithControlChars
	case PayloadEncoding_PAYLOAD_ENCODING_MESSAGE_PACK:
		encoding = serde.PayloadEncodingMsgPack
	case PayloadEncoding_PAYLOAD_ENCODING_SMILE:
		encoding = serde.PayloadEncodingSmile
	case PayloadEncoding_PAYLOAD_ENCODING_UINT:
		encoding = serde.PayloadEncodingUint
	case PayloadEncoding_PAYLOAD_ENCODING_UNSPECIFIED,
		PayloadEncoding_PAYLOAD_ENCODING_BINARY,
		PayloadEncoding_PAYLOAD_ENCODING_CONSUMER_OFFSETS:
		encoding = serde.PayloadEncodingBinary
	case PayloadEncoding_PAYLOAD_ENCODING_CBOR:
		encoding = serde.PayloadEncodingCbor
	}

	input := &serde.RecordPayloadInput{
		Payload:  po.Data,
		Encoding: encoding,
	}

	if *po.SchemaId > 0 {
		input.Options = []serde.SerdeOpt{serde.WithSchemaID(uint32(*po.SchemaId))}
	}

	if *po.Index > 0 {
		input.Options = append(input.Options, serde.WithIndex(int(*po.Index)))
	}

	return input
}

func rpcCompressionTypeToKgoCodec(compressionType CompressionType) []kgo.CompressionCodec {
	switch compressionType {
	case CompressionType_COMPRESSION_TYPE_UNCOMPRESSED, CompressionType_COMPRESSION_TYPE_UNSPECIFIED:
		return []kgo.CompressionCodec{kgo.NoCompression()}
	case CompressionType_COMPRESSION_TYPE_GZIP:
		return []kgo.CompressionCodec{kgo.GzipCompression(), kgo.NoCompression()}
	case CompressionType_COMPRESSION_TYPE_SNAPPY:
		return []kgo.CompressionCodec{kgo.SnappyCompression(), kgo.NoCompression()}
	case CompressionType_COMPRESSION_TYPE_LZ4:
		return []kgo.CompressionCodec{kgo.Lz4Compression(), kgo.NoCompression()}
	case CompressionType_COMPRESSION_TYPE_ZSTD:
		return []kgo.CompressionCodec{kgo.ZstdCompression(), kgo.NoCompression()}
	default:
		return []kgo.CompressionCodec{kgo.NoCompression()}
	}
}

func toProtoEncoding(serdeEncoding serde.PayloadEncoding) PayloadEncoding { //nolint:cyclop // we have to map all possible values here
	encoding := PayloadEncoding_PAYLOAD_ENCODING_BINARY

	switch serdeEncoding {
	case serde.PayloadEncodingNull:
		encoding = PayloadEncoding_PAYLOAD_ENCODING_NULL
	case serde.PayloadEncodingAvro:
		encoding = PayloadEncoding_PAYLOAD_ENCODING_AVRO
	case serde.PayloadEncodingProtobuf:
		encoding = PayloadEncoding_PAYLOAD_ENCODING_PROTOBUF
	case serde.PayloadEncodingProtobufSchema:
		encoding = PayloadEncoding_PAYLOAD_ENCODING_PROTOBUF_SCHEMA
	case serde.PayloadEncodingJSON:
		encoding = PayloadEncoding_PAYLOAD_ENCODING_JSON
	case serde.PayloadEncodingJSONSchema:
		encoding = PayloadEncoding_PAYLOAD_ENCODING_JSON_SCHEMA
	case serde.PayloadEncodingXML:
		encoding = PayloadEncoding_PAYLOAD_ENCODING_XML
	case serde.PayloadEncodingText:
		encoding = PayloadEncoding_PAYLOAD_ENCODING_TEXT
	case serde.PayloadEncodingUtf8WithControlChars:
		encoding = PayloadEncoding_PAYLOAD_ENCODING_UTF8
	case serde.PayloadEncodingMsgPack:
		encoding = PayloadEncoding_PAYLOAD_ENCODING_MESSAGE_PACK
	case serde.PayloadEncodingSmile:
		encoding = PayloadEncoding_PAYLOAD_ENCODING_SMILE
	case serde.PayloadEncodingUint:
		encoding = PayloadEncoding_PAYLOAD_ENCODING_UINT
	case serde.PayloadEncodingBinary:
		encoding = PayloadEncoding_PAYLOAD_ENCODING_BINARY
	case serde.PayloadEncodingConsumerOffsets:
		encoding = PayloadEncoding_PAYLOAD_ENCODING_CONSUMER_OFFSETS
	case serde.PayloadEncodingUnspecified:
		encoding = PayloadEncoding_PAYLOAD_ENCODING_UNSPECIFIED
	case serde.PayloadEncodingCbor:
		encoding = PayloadEncoding_PAYLOAD_ENCODING_CBOR
	}

	return encoding
}

func fromProtoEncoding(protoEncoding PayloadEncoding) serde.PayloadEncoding { //nolint:cyclop // we have to map all possible values here
	encoding := serde.PayloadEncodingUnspecified

	switch protoEncoding {
	case PayloadEncoding_PAYLOAD_ENCODING_NULL:
		encoding = serde.PayloadEncodingNull
	case PayloadEncoding_PAYLOAD_ENCODING_AVRO:
		encoding = serde.PayloadEncodingAvro
	case PayloadEncoding_PAYLOAD_ENCODING_PROTOBUF:
		encoding = serde.PayloadEncodingProtobuf
	case PayloadEncoding_PAYLOAD_ENCODING_PROTOBUF_SCHEMA:
		encoding = serde.PayloadEncodingProtobufSchema
	case PayloadEncoding_PAYLOAD_ENCODING_JSON:
		encoding = serde.PayloadEncodingJSON
	case PayloadEncoding_PAYLOAD_ENCODING_JSON_SCHEMA:
		encoding = serde.PayloadEncodingJSONSchema
	case PayloadEncoding_PAYLOAD_ENCODING_XML:
		encoding = serde.PayloadEncodingXML
	case PayloadEncoding_PAYLOAD_ENCODING_TEXT:
		encoding = serde.PayloadEncodingText
	case PayloadEncoding_PAYLOAD_ENCODING_UTF8:
		encoding = serde.PayloadEncodingUtf8WithControlChars
	case PayloadEncoding_PAYLOAD_ENCODING_MESSAGE_PACK:
		encoding = serde.PayloadEncodingMsgPack
	case PayloadEncoding_PAYLOAD_ENCODING_SMILE:
		encoding = serde.PayloadEncodingSmile
	case PayloadEncoding_PAYLOAD_ENCODING_UINT:
		encoding = serde.PayloadEncodingUint
	case PayloadEncoding_PAYLOAD_ENCODING_BINARY:
		encoding = serde.PayloadEncodingBinary
	case PayloadEncoding_PAYLOAD_ENCODING_CONSUMER_OFFSETS:
		encoding = serde.PayloadEncodingConsumerOffsets
	case PayloadEncoding_PAYLOAD_ENCODING_UNSPECIFIED:
		encoding = serde.PayloadEncodingUnspecified
	case PayloadEncoding_PAYLOAD_ENCODING_CBOR:
		encoding = serde.PayloadEncodingCbor
	}

	return encoding
}

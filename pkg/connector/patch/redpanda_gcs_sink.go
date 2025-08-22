// Copyright 2023 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package patch

import (
	"regexp"
	"strings"

	"github.com/xxxcrel/kafka-console/pkg/connector/model"
	"github.com/xxxcrel/kafka-console/pkg/random"
)

// ConfigPatchRedpandaGCS is a config patch that includes changes that shall be applied to the
// Redpanda GCS Sink connector.
type ConfigPatchRedpandaGCS struct {
	ConfigurationKeySelector IncludeExcludeSelector
	ConnectorClassSelector   IncludeExcludeSelector
}

var _ ConfigPatch = (*ConfigPatchRedpandaGCS)(nil)

const formatOutputType = "format.output.type"

// NewConfigPatchRedpandaGCS returns a new Patch for the Redpanda GCS connector.
func NewConfigPatchRedpandaGCS() *ConfigPatchRedpandaGCS {
	return &ConfigPatchRedpandaGCS{
		ConfigurationKeySelector: IncludeExcludeSelector{
			Include: regexp.MustCompile(`.*`),
			Exclude: nil,
		},
		ConnectorClassSelector: IncludeExcludeSelector{
			Include: regexp.MustCompile(`com.redpanda.kafka.connect.gcs\..*`),
			Exclude: nil,
		},
	}
}

// IsMatch implements the ConfigPatch.IsMatch interface.
func (c *ConfigPatchRedpandaGCS) IsMatch(configKey, connectorClass string) bool {
	return c.ConfigurationKeySelector.IsMatch(configKey) && c.ConnectorClassSelector.IsMatch(connectorClass)
}

// PatchDefinition implements the ConfigPatch.PatchDefinition interface.
func (*ConfigPatchRedpandaGCS) PatchDefinition(d model.ConfigDefinition, _ string) model.ConfigDefinition {
	// Misc patches
	switch d.Definition.Name {
	case "gcs.credentials.json":
		d.SetDocumentation("")
	case keyConverter:
		d.ClearRecommendedValuesWithMetadata().
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.converters.ByteArrayConverter", "BYTEARRAY").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.storage.StringConverter", "STRING").
			AddRecommendedValueWithMetadata("io.confluent.connect.avro.AvroConverter", "AVRO").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.json.JsonConverter", "JSON").
			AddRecommendedValueWithMetadata("com.redpanda.connectors.converter.protobuf.ProtobufConverter", "PROTOBUF").
			SetDefaultValue("org.apache.kafka.connect.converters.ByteArrayConverter")
	case valueConverter:
		d.ClearRecommendedValuesWithMetadata().
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.converters.ByteArrayConverter", "BYTEARRAY").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.storage.StringConverter", "STRING").
			AddRecommendedValueWithMetadata("io.confluent.connect.avro.AvroConverter", "AVRO").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.json.JsonConverter", "JSON").
			AddRecommendedValueWithMetadata("com.redpanda.connectors.converter.protobuf.ProtobufConverter", "PROTOBUF").
			SetDefaultValue("org.apache.kafka.connect.converters.ByteArrayConverter")
	case formatOutputType:
		d.SetDisplayName("GCS file format").
			SetDocumentation("Format of the key coming from the Redpanda topic").
			SetComponentType(model.ComponentRadioGroup)
	case "file.compression.type":
		d.SetDisplayName("Output file compression").
			SetComponentType(model.ComponentRadioGroup)
	case "format.output.fields":
		d.SetDisplayName("Output fields")
	case "format.output.fields.value.encoding":
		d.SetDisplayName("Value field encoding").
			SetComponentType(model.ComponentRadioGroup)
	case "format.output.envelope":
		d.SetDisplayName("Envelope for primitives")
	case "file.max.records":
		d.SetDisplayName("Max records per file")
	case name:
		d.SetDefaultValue("gcs-connector-" + strings.ToLower(random.String(4)))

	// Below properties will be grouped into "Error Handling"
	case errorsRetryTimeout:
		d.SetDisplayName("Retry timeout")
	case "kafka.retry.backoff.ms":
		d.SetDisplayName("Retry back-off").
			SetDocumentation("Retry backoff in milliseconds. Useful for performing recovery in case " +
				"of transient exceptions. Maximum value is 86400000 (24 hours)")
	}

	// Importance Patches
	switch d.Definition.Name {
	case formatOutputType:
		d.SetImportance(model.ConfigDefinitionImportanceHigh)
	case configActionReload:
		d.SetImportance(model.ConfigDefinitionImportanceLow)
	case "avro.codec":
		d.SetComponentType(model.ComponentRadioGroup).
			SetDocumentation("The Avro compression codec to be used for Avro output files").
			AddRecommendedValueWithMetadata("null", "uncompressed").
			AddRecommendedValueWithMetadata("deflate", "deflate").
			AddRecommendedValueWithMetadata("snappy", "snappy").
			AddRecommendedValueWithMetadata("bzip2", "bzip2").
			SetDefaultValue("null")
	}

	return d
}

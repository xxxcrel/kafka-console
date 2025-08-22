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

	"github.com/xxxcrel/kafka-console/pkg/connector/model"
)

const (
	errorsRetryTimeout = "errors.retry.timeout"
	name               = "name"
	headerConverter    = "header.converter"
	keyConverter       = "key.converter"
	valueConverter     = "value.converter"
	configActionReload = "config.action.reload"
	tableIncludeList   = "table.include.list"
)

// ConfigPatchCommon is a config patch that applies specific patches on a set of configurations
// that is common across several connectors (e.g. tasks.max).
type ConfigPatchCommon struct {
	ConfigurationKeySelector IncludeExcludeSelector
	ConnectorClassSelector   IncludeExcludeSelector
}

var _ ConfigPatch = (*ConfigPatchAll)(nil)

// NewConfigPatchCommon returns a new ConfigPatch for common connector configurations.
func NewConfigPatchCommon() *ConfigPatchCommon {
	return &ConfigPatchCommon{
		ConfigurationKeySelector: IncludeExcludeSelector{
			Include: regexp.MustCompile(`(tasks.max|key.converter|value.converter|header.converter|config.action.reload)`),
			Exclude: nil,
		},
		ConnectorClassSelector: IncludeExcludeSelector{
			Include: regexp.MustCompile(".*"),
			Exclude: nil,
		},
	}
}

// IsMatch implements the ConfigPatch.IsMatch interface.
func (c *ConfigPatchCommon) IsMatch(configKey, connectorClass string) bool {
	return c.ConfigurationKeySelector.IsMatch(configKey) && c.ConnectorClassSelector.IsMatch(connectorClass)
}

// PatchDefinition implements the ConfigPatch.PatchDefinition interface.
func (*ConfigPatchCommon) PatchDefinition(d model.ConfigDefinition, _ string) model.ConfigDefinition {
	switch d.Definition.Name {
	case "tasks.max":
		d.SetDisplayName("Max tasks").SetImportance(model.ConfigDefinitionImportanceHigh)
	case keyConverter:
		d.SetDisplayName("Redpanda message key format").
			SetDocumentation("Format of the key in the Redpanda topic").
			SetImportance(model.ConfigDefinitionImportanceHigh).
			SetComponentType(model.ComponentRadioGroup).
			AddRecommendedValueWithMetadata("io.confluent.connect.avro.AvroConverter", "AVRO").
			AddRecommendedValueWithMetadata("io.confluent.connect.protobuf.ProtobufConverter", "PROTOBUF").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.json.JsonConverter", "JSON").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.storage.StringConverter", "STRING").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.converters.ByteArrayConverter", "BYTES").
			SetDefaultValue("org.apache.kafka.connect.converters.ByteArrayConverter")
	case "value.converter":
		d.SetDisplayName("Redpanda message value format").
			SetDocumentation("Format of the value in the Redpanda topic").
			SetImportance(model.ConfigDefinitionImportanceHigh).
			SetComponentType(model.ComponentRadioGroup).
			AddRecommendedValueWithMetadata("io.confluent.connect.avro.AvroConverter", "AVRO").
			AddRecommendedValueWithMetadata("io.confluent.connect.protobuf.ProtobufConverter", "PROTOBUF").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.json.JsonConverter", "JSON").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.storage.StringConverter", "STRING").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.converters.ByteArrayConverter", "BYTES").
			SetDefaultValue("org.apache.kafka.connect.converters.ByteArrayConverter")
	case headerConverter:
		d.SetDisplayName("Redpanda message headers format").
			SetDocumentation("Format of the headers in the Redpanda topic").
			SetImportance(model.ConfigDefinitionImportanceLow).
			SetComponentType(model.ComponentRadioGroup).
			AddRecommendedValueWithMetadata("io.confluent.connect.avro.AvroConverter", "AVRO").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.json.JsonConverter", "JSON").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.storage.StringConverter", "STRING").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.converters.ByteArrayConverter", "BYTES").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.storage.SimpleHeaderConverter", "SIMPLE").
			SetDefaultValue("org.apache.kafka.connect.storage.SimpleHeaderConverter")
	case configActionReload:
		d.SetComponentType(model.ComponentRadioGroup).
			AddRecommendedValueWithMetadata("restart", "RESTART").
			AddRecommendedValueWithMetadata("none", "NONE")
	}
	return d
}

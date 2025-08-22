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

// ConfigPatchBigQuery is a config patch that includes changes that shall be applied to the
// BigQuery sink connector.
type ConfigPatchBigQuery struct {
	ConfigurationKeySelector IncludeExcludeSelector
	ConnectorClassSelector   IncludeExcludeSelector
}

var _ ConfigPatch = (*ConfigPatchBigQuery)(nil)

// NewConfigPatchBigQuery returns a new Patch for the BigQuery sink connector.
func NewConfigPatchBigQuery() *ConfigPatchBigQuery {
	return &ConfigPatchBigQuery{
		ConfigurationKeySelector: IncludeExcludeSelector{
			Include: regexp.MustCompile(`.*`),
			Exclude: nil,
		},
		ConnectorClassSelector: IncludeExcludeSelector{
			Include: regexp.MustCompile(`com.wepay.kafka.connect.bigquery\..*`),
			Exclude: nil,
		},
	}
}

// IsMatch implements the ConfigPatch.IsMatch interface.
func (c *ConfigPatchBigQuery) IsMatch(configKey, connectorClass string) bool {
	return c.ConfigurationKeySelector.IsMatch(configKey) && c.ConnectorClassSelector.IsMatch(connectorClass)
}

// PatchDefinition implements the ConfigPatch.PatchDefinition interface.
//
//nolint:cyclop // This function defines/patches a lot of things, but it's easy to comprehend.
func (*ConfigPatchBigQuery) PatchDefinition(d model.ConfigDefinition, _ string) model.ConfigDefinition {
	// Misc patches
	switch d.Definition.Name {
	case valueConverter:
		d.ClearRecommendedValuesWithMetadata().
			AddRecommendedValueWithMetadata("io.confluent.connect.avro.AvroConverter", "AVRO").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.json.JsonConverter", "JSON").
			AddRecommendedValueWithMetadata("com.redpanda.connectors.converter.protobuf.ProtobufConverter", "PROTOBUF").
			SetDefaultValue("org.apache.kafka.connect.json.JsonConverter")
	case "keySource":
		d.Value.Value = "JSON"
	case "keyfile":
		d.SetRequired(true).
			SetDisplayName("Credentials JSON").
			SetDocumentation("The JSON key with BigQuery service account credentials")
	case "sanitizeTopics":
		d.SetDefaultValue("true").
			SetDisplayName("Sanitize topics")
	case "defaultDataset":
		d.SetDisplayName("Default dataset")
	case "autoCreateTables":
		d.SetDisplayName("Auto create tables").
			SetDocumentation("Automatically create BigQuery tables if they don't already exist. If the table does not exist, it is created based on the record schema")
	case "topic2TableMap":
		d.SetDisplayName("Topic to table map")
	case "allowNewBigQueryFields":
		d.SetDisplayName("Allow new BigQuery fields")
	case "allowBigQueryRequiredFieldRelaxation":
		d.SetDisplayName("Allow BigQuery required field relaxation")
	case "upsertEnabled":
		d.SetDisplayName("Upsert enabled")
	case "deleteEnabled":
		d.SetDisplayName("Delete enabled")
	case "kafkaKeyFieldName":
		d.SetDisplayName("Kafka key field name")
	case "bigQueryRetry":
		d.SetDisplayName("BigQuery retry attempts")
	case "bigQueryRetryWait":
		d.SetDisplayName("BigQuery retry attempts interval")
	case "timePartitioningType":
		d.SetDisplayName("Time partitioning type").
			SetComponentType(model.ComponentRadioGroup)
	case name:
		d.SetDefaultValue("bigquery-connector-" + strings.ToLower(random.String(4)))
	}

	// Importance Patches
	switch d.Definition.Name {
	case "keyfile":
		d.SetImportance(model.ConfigDefinitionImportanceHigh)
	case keyConverter,
		"autoCreateTables",
		"kafkaKeyFieldName":
		d.SetImportance(model.ConfigDefinitionImportanceMedium)
	}

	return d
}

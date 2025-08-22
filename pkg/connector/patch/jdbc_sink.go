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

// ConfigPatchJdbcSink is a config patch that includes changes that shall be applied to the
// JDBC sink connector.
type ConfigPatchJdbcSink struct {
	ConfigurationKeySelector IncludeExcludeSelector
	ConnectorClassSelector   IncludeExcludeSelector
}

var _ ConfigPatch = (*ConfigPatchJdbcSink)(nil)

// NewConfigPatchJdbcSink returns a new Patch for the JDBC sink connector.
func NewConfigPatchJdbcSink() *ConfigPatchJdbcSink {
	return &ConfigPatchJdbcSink{
		ConfigurationKeySelector: IncludeExcludeSelector{
			Include: regexp.MustCompile(`.*`),
			Exclude: nil,
		},
		ConnectorClassSelector: IncludeExcludeSelector{
			Include: regexp.MustCompile(`com.redpanda.kafka.connect.jdbc.JdbcSinkConnector`),
			Exclude: nil,
		},
	}
}

// IsMatch implements the ConfigPatch.IsMatch interface.
func (c *ConfigPatchJdbcSink) IsMatch(configKey, connectorClass string) bool {
	return c.ConfigurationKeySelector.IsMatch(configKey) && c.ConnectorClassSelector.IsMatch(connectorClass)
}

// PatchDefinition implements the ConfigPatch.PatchDefinition interface.
func (*ConfigPatchJdbcSink) PatchDefinition(d model.ConfigDefinition, _ string) model.ConfigDefinition {
	// Misc patches
	switch d.Definition.Name {
	case keyConverter:
		d.ClearRecommendedValuesWithMetadata().
			SetImportance(model.ConfigDefinitionImportanceHigh).
			AddRecommendedValueWithMetadata("io.confluent.connect.avro.AvroConverter", "AVRO").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.json.JsonConverter", "JSON").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.storage.StringConverter", "STRING").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.converters.ByteArrayConverter", "BYTES").
			AddRecommendedValueWithMetadata("com.redpanda.connectors.converter.protobuf.ProtobufConverter", "PROTOBUF").
			SetDefaultValue("org.apache.kafka.connect.json.JsonConverter")
	case valueConverter:
		d.ClearRecommendedValuesWithMetadata().
			AddRecommendedValueWithMetadata("io.confluent.connect.avro.AvroConverter", "AVRO").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.json.JsonConverter", "JSON").
			AddRecommendedValueWithMetadata("com.redpanda.connectors.converter.protobuf.ProtobufConverter", "PROTOBUF").
			SetDefaultValue("org.apache.kafka.connect.json.JsonConverter")
	case "insert.mode":
		d.SetDisplayName("Insert Mode").
			SetDocumentation("The insertion mode to use. INSERT - Use standard SQL 'INSERT' statements. MULTI - Use multi-row `INSERT` statements. 'UPSERT' - Use the appropriate upsert semantics for the target database if it is supported by the connector, e.g. 'INSERT .. ON CONFLICT .. DO UPDATE SET ..'. 'UPDATE' - Use the appropriate update semantics for the target database if it is supported by the connector, e.g. 'UPDATE'").
			SetImportance(model.ConfigDefinitionImportanceMedium).
			SetComponentType(model.ComponentRadioGroup).
			AddRecommendedValueWithMetadata("insert", "INSERT").
			AddRecommendedValueWithMetadata("multi", "MULTI").
			AddRecommendedValueWithMetadata("upsert", "UPSERT").
			AddRecommendedValueWithMetadata("update", "UPDATE").
			SetDefaultValue("insert")
	case "pk.mode":
		d.SetDisplayName("Primary Key Mode").
			SetDocumentation("The primary key mode, also refer to 'Primary Key Fields' documentation for interplay. Supported modes are: 'NONE' - No keys utilized. 'KAFKA' - Redpanda coordinates (the topic, partition, and offset) are used as the PK. 'RECORD_KEY' - Field(s) from the record key are used, which may be a primitive or a struct. 'RECORD_VALUE' - Field(s) from the record value are used, which must be a struct").
			SetImportance(model.ConfigDefinitionImportanceMedium).
			SetComponentType(model.ComponentRadioGroup).
			AddRecommendedValueWithMetadata("none", "NONE").
			AddRecommendedValueWithMetadata("kafka", "KAFKA").
			AddRecommendedValueWithMetadata("record_key", "RECORD_KEY").
			AddRecommendedValueWithMetadata("record_value", "RECORD_VALUE").
			SetDefaultValue("none")
	case "dialect.name":
		d.SetComponentType(model.ComponentRadioGroup).
			AddRecommendedValueWithMetadata("", "AUTO").
			AddRecommendedValueWithMetadata("MySqlDatabaseDialect", "MySQL").
			AddRecommendedValueWithMetadata("PostgreSqlDatabaseDialect", "PostgreSQL").
			AddRecommendedValueWithMetadata("SqliteDatabaseDialect", "SQLite").
			AddRecommendedValueWithMetadata("SqlServerDatabaseDialect", "SQL Server").
			SetDefaultValue("").
			SetDocumentation("The name of the database dialect that should be used for this connector. By default. the connector automatically determines the dialect based upon the JDBC connection URL. Use this if you want to override that behavior and use a specific dialect")
	case name:
		d.SetDefaultValue("jdbc-sink-connector-" + strings.ToLower(random.String(4)))
	case "connection.url":
		d.SetDocumentation("Database JDBC connection URL")
	case "connection.user":
		d.SetDisplayName("User").
			SetDocumentation("Name of the database user to be used when connecting to the database")
	case "connection.password":
		d.SetDisplayName("Password").
			SetDocumentation("Password of the database user to be used when connecting to the database")
	case "fields.whitelist":
		d.SetDisplayName("Include Fields").
			SetDocumentation("List of comma-separated record value field names. If empty, all fields from the record value are utilized, otherwise used to filter to the desired fields. Note that `Primary Key Fields` is applied independently in the context of which field(s) form the primary key columns in the destination database, while this configuration is applicable for the other columns")
	}

	// Importance Patches
	switch d.Definition.Name {
	case "pk.mode", "insert.mode":
		d.SetImportance(model.ConfigDefinitionImportanceMedium)
	case "auto.create":
		d.SetImportance(model.ConfigDefinitionImportanceHigh)
	}

	return d
}

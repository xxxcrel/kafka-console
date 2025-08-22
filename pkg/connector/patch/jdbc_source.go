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

// ConfigPatchJdbcSource is a config patch that includes changes that shall be applied to the
// JDBC source connector.
type ConfigPatchJdbcSource struct {
	ConfigurationKeySelector IncludeExcludeSelector
	ConnectorClassSelector   IncludeExcludeSelector
}

var _ ConfigPatch = (*ConfigPatchJdbcSource)(nil)

// NewConfigPatchJdbcSource returns a new Patch for the JDBC source connector.
func NewConfigPatchJdbcSource() *ConfigPatchJdbcSource {
	return &ConfigPatchJdbcSource{
		ConfigurationKeySelector: IncludeExcludeSelector{
			Include: regexp.MustCompile(`.*`),
			Exclude: nil,
		},
		ConnectorClassSelector: IncludeExcludeSelector{
			Include: regexp.MustCompile(`com.redpanda.kafka.connect.jdbc.JdbcSourceConnector`),
			Exclude: nil,
		},
	}
}

// IsMatch implements the ConfigPatch.IsMatch interface.
func (c *ConfigPatchJdbcSource) IsMatch(configKey, connectorClass string) bool {
	return c.ConfigurationKeySelector.IsMatch(configKey) && c.ConnectorClassSelector.IsMatch(connectorClass)
}

// PatchDefinition implements the ConfigPatch.PatchDefinition interface.
func (*ConfigPatchJdbcSource) PatchDefinition(d model.ConfigDefinition, _ string) model.ConfigDefinition {
	// Misc patches
	switch d.Definition.Name {
	case keyConverter:
		d.ClearRecommendedValuesWithMetadata().
			AddRecommendedValueWithMetadata("io.confluent.connect.avro.AvroConverter", "AVRO").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.json.JsonConverter", "JSON").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.storage.StringConverter", "STRING").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.converters.ByteArrayConverter", "BYTES").
			AddRecommendedValueWithMetadata("io.confluent.connect.protobuf.ProtobufConverter", "PROTOBUF").
			SetDefaultValue("org.apache.kafka.connect.json.JsonConverter")
	case valueConverter:
		d.ClearRecommendedValuesWithMetadata().
			AddRecommendedValueWithMetadata("io.confluent.connect.avro.AvroConverter", "AVRO").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.json.JsonConverter", "JSON").
			SetDefaultValue("org.apache.kafka.connect.json.JsonConverter")
	case "mode":
		d.SetComponentType(model.ComponentRadioGroup).
			AddRecommendedValueWithMetadata("bulk", "BULK").
			AddRecommendedValueWithMetadata("incrementing", "INCREMENTING").
			AddRecommendedValueWithMetadata("timestamp", "TIMESTAMP").
			AddRecommendedValueWithMetadata("timestamp+incrementing", "TIMESTAMP+INCREMENTING").
			SetDefaultValue("bulk")
	case "dialect.name":
		d.SetComponentType(model.ComponentRadioGroup).
			AddRecommendedValueWithMetadata("", "AUTO").
			AddRecommendedValueWithMetadata("MySqlDatabaseDialect", "MySQL").
			AddRecommendedValueWithMetadata("PostgreSqlDatabaseDialect", "PostgreSQL").
			AddRecommendedValueWithMetadata("SqliteDatabaseDialect", "SQLite").
			AddRecommendedValueWithMetadata("SqlServerDatabaseDialect", "SQL Server").
			SetDefaultValue("").
			SetDocumentation("The name of the database dialect that should be used for this connector. By default. the connector automatically determines the dialect based upon the JDBC connection URL. Use this if you want to override that behavior and use a specific dialect")
	case "numeric.mapping":
		d.SetDefaultValue("none")
	case "table.blacklist":
		d.SetDocumentation("List of tables to exclude from copying. If specified, Include Tables may not be set").
			SetDisplayName("Exclude Tables")
	case "table.whitelist":
		d.SetDocumentation("List of tables to include in copying. If specified, Exclude Tables may not be set").
			SetDisplayName("Include Tables")
	case "catalog.pattern":
		d.SetDocumentation("Catalog pattern to fetch table metadata from the database. null - (default) means that the catalog name should not be used to narrow the search so that all table metadata would be fetched, regardless of their catalog. \"\" - retrieves those without a catalog")
	case "connection.url":
		d.SetDocumentation("Database JDBC connection URL")
	case "connection.user":
		d.SetDisplayName("User").
			SetDocumentation("Name of the database user to be used when connecting to the database")
	case "connection.password":
		d.SetDisplayName("Password").
			SetDocumentation("Password of the database user to be used when connecting to the database")
	case "validate.non.null", "timestamp.column.name", "incrementing.column.name":
		d.SetVisible(true)
	case name:
		d.SetDefaultValue("jdbc-source-connector-" + strings.ToLower(random.String(4)))
	}

	// Importance Patches
	switch d.Definition.Name {
	case "mode", "poll.interval.ms", "timestamp.delay.interval.ms", "key.converter":
		d.SetImportance(model.ConfigDefinitionImportanceMedium)
	}

	return d
}

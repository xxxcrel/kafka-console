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

// ConfigPatchMongoDB is a config patch that includes changes that shall be applied to the
// MongoDB connectors.
type ConfigPatchMongoDB struct {
	ConfigurationKeySelector IncludeExcludeSelector
	ConnectorClassSelector   IncludeExcludeSelector
}

var _ ConfigPatch = (*ConfigPatchMongoDB)(nil)

const (
	mongoClassSelectorRegexp = `com.mongodb.kafka.connect.Mongo(Source|Sink)Connector`
)

// NewConfigPatchMongoDB returns a new Patch for the MongoDB connectors.
func NewConfigPatchMongoDB() *ConfigPatchMongoDB {
	return &ConfigPatchMongoDB{
		ConfigurationKeySelector: IncludeExcludeSelector{
			Include: regexp.MustCompile(`.*`),
			Exclude: nil,
		},
		ConnectorClassSelector: IncludeExcludeSelector{
			Include: regexp.MustCompile(mongoClassSelectorRegexp),
			Exclude: nil,
		},
	}
}

// IsMatch implements the ConfigPatch.IsMatch interface.
func (c *ConfigPatchMongoDB) IsMatch(configKey, connectorClass string) bool {
	return c.ConfigurationKeySelector.IsMatch(configKey) && c.ConnectorClassSelector.IsMatch(connectorClass)
}

// PatchDefinition implements the ConfigPatch.PatchDefinition interface.
func (*ConfigPatchMongoDB) PatchDefinition(d model.ConfigDefinition, connectorClass string) model.ConfigDefinition {
	d = patchSinkConnector(d, connectorClass)

	// Misc patches
	switch d.Definition.Name {
	case "connection.uri":
		d.SetDefaultValue("mongodb://").
			SetVisible(false).
			SetType(model.ConfigDefinitionTypePassword)
	case keyConverter, valueConverter:
		converterType, _, _ := strings.Cut(d.Definition.Name, ".")
		d.ClearRecommendedValuesWithMetadata().
			AddRecommendedValueWithMetadata("io.confluent.connect.avro.AvroConverter", "AVRO").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.json.JsonConverter", "JSON").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.storage.StringConverter", "STRING").
			AddRecommendedValueWithMetadata("org.apache.kafka.connect.converters.ByteArrayConverter", "BYTES").
			SetDefaultValue("org.apache.kafka.connect.storage.StringConverter")
		if strings.HasSuffix(connectorClass, "SourceConnector") {
			d.SetDocumentation("Format of the " + converterType + " in the Redpanda topic. Use AVRO or JSON for schematic output, STRING for plain JSON or BYTES for BSON")
		}
		if strings.HasSuffix(connectorClass, "SinkConnector") {
			d.AddRecommendedValueWithMetadata("com.redpanda.connectors.converter.protobuf.ProtobufConverter", "PROTOBUF")
		}
	case "output.schema.infer.value":
		d.SetDocumentation("Infer the schema for the value. Each Document will be processed in isolation, which may lead to multiple schema definitions for the data. Only applied when Redpanda message value format is set to AVRO or JSON")
	case "change.stream.full.document",
		"change.stream.full.document.before.change":
		d.SetComponentType(model.ComponentRadioGroup)
	case "startup.mode":
		d.AddRecommendedValueWithMetadata("latest", "LATEST").
			AddRecommendedValueWithMetadata("timestamp", "TIMESTAMP").
			AddRecommendedValueWithMetadata("copy_existing", "COPY_EXISTING").
			SetComponentType(model.ComponentRadioGroup).
			SetDefaultValue("latest")
	case "key.projection.type":
		d.AddRecommendedValueWithMetadata("none", "NONE").
			AddRecommendedValueWithMetadata("allowlist", "ALLOWLIST").
			AddRecommendedValueWithMetadata("blocklist", "BLOCKLIST").
			SetComponentType(model.ComponentRadioGroup).
			SetDefaultValue("none").
			SetDocumentation("The type of key projection to use, either: `AllowList` or `BlockList`")
	case "value.projection.type":
		d.AddRecommendedValueWithMetadata("none", "NONE").
			AddRecommendedValueWithMetadata("allowlist", "ALLOWLIST").
			AddRecommendedValueWithMetadata("blocklist", "BLOCKLIST").
			SetComponentType(model.ComponentRadioGroup).
			SetDefaultValue("none").
			SetDisplayName("The value projection type").
			SetDocumentation("The type of value projection to use, either: `AllowList` or `BlockList`")
	case "change.data.capture.handler":
		d.AddRecommendedValueWithMetadata("", "NONE").
			AddRecommendedValueWithMetadata("com.mongodb.kafka.connect.sink.cdc.mongodb.ChangeStreamHandler", "MongoDB").
			AddRecommendedValueWithMetadata("com.mongodb.kafka.connect.sink.cdc.debezium.mongodb.MongoDbHandler", "Debezium MongoDB").
			AddRecommendedValueWithMetadata("com.mongodb.kafka.connect.sink.cdc.debezium.rdbms.postgres.PostgresHandler", "Debezium Postgres").
			AddRecommendedValueWithMetadata("com.mongodb.kafka.connect.sink.cdc.debezium.rdbms.mysql.MysqlHandler", "Debezium MySQL").
			AddRecommendedValueWithMetadata("com.mongodb.kafka.connect.sink.cdc.qlik.rdbms.RdbmsHandler", "Qlik").
			SetComponentType(model.ComponentRadioGroup).
			SetDefaultValue("").
			SetDocumentation("The CDC handler to use for processing. MongoDB handler requires plain JSON or BSON format")
	case "mongo.errors.tolerance":
		d.AddRecommendedValueWithMetadata("none", "NONE").
			AddRecommendedValueWithMetadata("all", "ALL").
			SetComponentType(model.ComponentRadioGroup).
			SetDocumentation("Behavior for tolerating errors during connector operation. 'NONE' is the default value and signals that any error will result in an immediate connector task failure; 'ALL' changes the behavior to skip over problematic records")
	case "heartbeat.interval.ms":
		d.SetDisplayName("Heartbeat interval milliseconds")
	case name:
		d.SetDefaultValue("mongodb-" + extractType(connectorClass, mongoClassSelectorRegexp) + "-connector-" + strings.ToLower(random.String(4)))
	}

	d = patchImportance(d)

	return d
}

func patchImportance(d model.ConfigDefinition) model.ConfigDefinition {
	// Importance Patches
	switch d.Definition.Name {
	case "topic.prefix",
		"database",
		"collection",
		"startup.mode":
		d.SetImportance(model.ConfigDefinitionImportanceHigh)
	case "change.stream.full.document",
		"change.stream.full.document.before.change",
		"publish.full.document.only",
		"collation":
		d.SetImportance(model.ConfigDefinitionImportanceMedium)
	case "output.schema.key",
		"output.schema.value":
		d.SetImportance(model.ConfigDefinitionImportanceLow)
	}

	return d
}

func patchSinkConnector(d model.ConfigDefinition, connectorClass string) model.ConfigDefinition {
	if isSink(connectorClass) {
		switch d.Definition.Name {
		case "database":
			d.SetDisplayName("MongoDB database name").
				SetDocumentation("The name of an existing MongoDB database to store output files in")
		case "collection":
			d.SetDisplayName("Default MongoDB collection name")
		}
	}

	return d
}

func isSink(connectorClass string) bool {
	return extractType(connectorClass, mongoClassSelectorRegexp) == "sink"
}

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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/xxxcrel/kafka-console/pkg/connector/model"
)

func TestPatchConfigAll(t *testing.T) {
	patcher := NewConfigPatchAll()

	tt := []struct {
		ConnectorClass   string
		ConfigurationKey string
		ShouldMatch      bool

		GivenDisplayName    string
		ExpectedDisplayName string
	}{
		{
			ConnectorClass:   "org.apache.kafka.connect.mirror.MirrorSourceConnector",
			ConfigurationKey: "ssl.protocol",
			ShouldMatch:      true,

			GivenDisplayName:    "SSL protocol",
			ExpectedDisplayName: "SSL protocol",
		},

		{
			ConnectorClass:   "org.apache.kafka.connect.mirror.MirrorSourceConnector",
			ConfigurationKey: "ssl.protocol",
			ShouldMatch:      true,

			GivenDisplayName:    "",
			ExpectedDisplayName: "SSL protocol",
		},

		{
			ConnectorClass:   "org.apache.kafka.connect.mirror.MirrorSourceConnector",
			ConfigurationKey: "ssl.protocol",
			ShouldMatch:      true,

			GivenDisplayName:    "ssl.protocol",
			ExpectedDisplayName: "SSL protocol",
		},

		{
			ConnectorClass:   "org.apache.kafka.connect.mirror.MirrorSourceConnector",
			ConfigurationKey: "refresh.topics.interval.seconds",
			ShouldMatch:      true,

			GivenDisplayName:    "",
			ExpectedDisplayName: "Refresh topics interval seconds",
		},

		{
			ConnectorClass:   "org.apache.kafka.connect.mirror.MirrorSourceConnector",
			ConfigurationKey: "sync.topic.acls.enabled",
			ShouldMatch:      true,

			GivenDisplayName:    "",
			ExpectedDisplayName: "Sync topic ACLs enabled",
		},

		{
			ConnectorClass:   "org.apache.kafka.connect.mirror.MirrorSourceConnector",
			ConfigurationKey: "foo.bar",
			ShouldMatch:      true,

			GivenDisplayName:    "DisplayName should always take precedence",
			ExpectedDisplayName: "DisplayName should always take precedence",
		},

		{
			ConnectorClass:   "org.apache.kafka.connect.mirror.MirrorSourceConnector",
			ConfigurationKey: "segment.ms",
			ShouldMatch:      true,

			GivenDisplayName:    "",
			ExpectedDisplayName: "Segment milliseconds",
		},

		{
			ConnectorClass:   "org.apache.kafka.connect.mirror.MirrorSourceConnector",
			ConfigurationKey: "errors.tolerance",
			ShouldMatch:      true,

			GivenDisplayName:    "",
			ExpectedDisplayName: "Error tolerance",
		},
	}

	for _, tc := range tt {
		name := fmt.Sprintf("%v-%v", tc.ConnectorClass, tc.ConfigurationKey)
		t.Run(name, func(t *testing.T) {
			isMatch := patcher.IsMatch(tc.ConfigurationKey, tc.ConnectorClass)
			require.Equal(t, tc.ShouldMatch, isMatch)
			if !isMatch {
				return
			}

			dummyConfigDefinition := model.ConfigDefinition{
				Definition: model.ConfigDefinitionKey{
					Name:        tc.ConfigurationKey,
					DisplayName: tc.GivenDisplayName,
				},
			}
			patchedDefinition := patcher.PatchDefinition(dummyConfigDefinition, "")
			assert.Equal(t, tc.ExpectedDisplayName, patchedDefinition.Definition.DisplayName)
		})
	}
}

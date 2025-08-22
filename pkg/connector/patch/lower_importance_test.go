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

func TestPatchConfigLowerImportance(t *testing.T) {
	patcher := NewConfigPatchLowerImportance()

	tt := []struct {
		ConnectorClass     string
		ConfigurationKey   string
		ShouldMatch        bool
		ExpectedImportance model.ConfigDefinitionImportance
	}{
		{
			ConnectorClass:     "org.apache.kafka.connect.mirror.MirrorSourceConnector",
			ConfigurationKey:   "ssl.protocol",
			ShouldMatch:        true,
			ExpectedImportance: model.ConfigDefinitionImportanceLow,
		},
		{
			ConnectorClass:   "org.apache.kafka.connect.mirror.MirrorSourceConnector",
			ConfigurationKey: "use.ssl",
			ShouldMatch:      false,
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
				Definition: model.ConfigDefinitionKey{Importance: model.ConfigDefinitionImportanceHigh},
			}
			patchedDefinition := patcher.PatchDefinition(dummyConfigDefinition, "")
			assert.Equal(t, tc.ExpectedImportance, patchedDefinition.Definition.Importance)
		})
	}
}

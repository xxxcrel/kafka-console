// Copyright 2023 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package guide

import "github.com/xxxcrel/kafka-console/pkg/connector/model"

// NewMirrorSourceGuide returns a new guide for MirrorSourceConnector.
func NewMirrorSourceGuide(opts ...Option) Guide {
	var o Options
	for _, opt := range opts {
		opt(&o)
	}

	return &WizardGuide{
		DefaultGuide: DefaultGuide{
			options: o,
		},
		className: "org.apache.kafka.connect.mirror.MirrorSourceConnector",
		wizardSteps: []model.ValidationResponseStep{
			{
				Name: "Regexes of topics to import",
				Groups: []model.ValidationResponseStepGroup{
					{
						// No Group name and description here
						ConfigKeys: []string{"topics"},
					},
				},
			},

			mirrorClusterConnection(),

			{
				Name: "Connector configuration",
				Groups: []model.ValidationResponseStepGroup{
					{
						// No Group name and description here
						ConfigKeys: []string{
							"sync.topic.configs.enabled",
							"sync.topic.configs.interval.seconds",
							"sync.topic.acls.enabled",
							"sync.topic.acls.interval.seconds",
							"topics.exclude",
							"source.cluster.alias",
							"replication.policy.class",
							"replication.factor",
							"refresh.topics.interval.seconds",
							"offset-syncs.topic.location",
							"offset-syncs.topic.replication.factor",
							"config.properties.exclude",
							"producer.override.compression.type",
							"producer.override.max.request.size",
							"consumer.auto.offset.reset",
							"offset.lag.max",
						},
					},
				},
			},

			sizing(),

			reviewAndLaunch(),
		},
	}
}

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

// NewRedpandaAwsS3SinkGuide returns a new guide for Redpanda's AWS S3 sink connector.
func NewRedpandaAwsS3SinkGuide(opts ...Option) Guide {
	var o Options
	for _, opt := range opts {
		opt(&o)
	}

	return &WizardGuide{
		DefaultGuide: DefaultGuide{
			options: o,
		},
		className: "com.redpanda.kafka.connect.s3.S3SinkConnector",
		wizardSteps: []model.ValidationResponseStep{
			topicsToExport(),

			{
				Name: "S3 connection",
				Groups: []model.ValidationResponseStepGroup{
					{
						Name:              "Authentication with AWS access keys",
						Description:       "An access key grants programmatic access to AWS resources",
						DocumentationLink: "https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_create.html",
						ConfigKeys:        []string{"aws.access.key.id", "aws.secret.access.key"},
					},
					{
						Name:       "S3 bucket settings",
						ConfigKeys: []string{"aws.s3.bucket.name", "aws.s3.region"},
					},
				},
			},

			{
				Name: "Connector configuration",
				Groups: []model.ValidationResponseStepGroup{
					{
						// No Group name and description here
						ConfigKeys: append([]string{
							"key.converter",
							"key.converter.schemas.enable",
							"value.converter",
							"value.converter.schemas.enable",
							"format.output.type",
							"avro.codec",

							"file.name.template",
							"file.name.prefix",
							"format.output.fields",
							"format.output.fields.value.encoding",
							"format.output.envelope",
							"file.compression.type",
							"file.max.records",
							"file.flush.interval.ms",
							"aws.s3.bucket.check",
							"aws.s3.part.size.bytes",
							"aws.s3.backoff.delay.ms",
							"aws.s3.backoff.max.delay.ms",
							"aws.s3.backoff.max.retries",
							"errors.tolerance",
						}, dlq()...),
					},
				},
			},

			sizing(),

			reviewAndLaunch(),
		},
	}
}

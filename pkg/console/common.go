// Copyright 2022 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package console

import "github.com/twmb/franz-go/pkg/kgo"

// BrokerRequestError is a helper struct that is used to wrap an error with the respective
// broker metadata that returned this error.
type BrokerRequestError struct {
	BrokerMeta kgo.BrokerMetadata `json:"brokerMetadata"`
	Error      error              `json:"error"`
}

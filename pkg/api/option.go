// Copyright 2022 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package api

import (
	"context"

	"go.uber.org/zap"

	"github.com/xxxcrel/kafka-console/pkg/factory/kafka"
	"github.com/xxxcrel/kafka-console/pkg/factory/schema"
	"github.com/xxxcrel/kafka-console/pkg/license"
)

type options struct {
	license              license.License
	kafkaClientProvider  kafka.ClientFactory
	schemaClientProvider schema.ClientFactory
	logger               *zap.Logger
	cacheNamespaceFn     func(context.Context) (string, error)
}

// Option is a function that applies some configuration to the options struct.
type Option func(*options)

// apply takes an options instance and applies the configuration.
func (opt Option) apply(opts *options) {
	opt(opts)
}

// WithLicense provides the license information which was used to start Redpanda
// Console. It is used only for visibility & logging purposes and not for any
// license enforcing actions.
func WithLicense(license license.License) Option {
	return func(o *options) {
		o.license = license
	}
}

// WithKafkaClientFactory uses the provided ClientFactory for creating new Kafka
// clients in all endpoint handlers.
func WithKafkaClientFactory(factory kafka.ClientFactory) Option {
	return func(o *options) {
		o.kafkaClientProvider = factory
	}
}

// WithSchemaClientFactory uses the provided ClientFactory for creating new
// Schema Registry clients in all endpoint handlers.
func WithSchemaClientFactory(factory schema.ClientFactory) Option {
	return func(o *options) {
		o.schemaClientProvider = factory
	}
}

// WithLogger allows to plug in your own pre-configured zap.Logger. If provided
// we will not try to set up our own logger.
func WithLogger(logger *zap.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

// WithCacheNamespaceFn is an option to set a function that determines the
// namespace for caching objects such as schemas. This should be set in
// multi-tenant environments where users and tenants should be strictly isolated
// from each other.
//
// The function has to return a unique identifier for the currently logged-in
// user that can be used as a namespace for caching. Only within that namespace
// cache objects will be stored and looked-up then.
func WithCacheNamespaceFn(fn func(context.Context) (string, error)) Option {
	return func(o *options) {
		o.cacheNamespaceFn = fn
	}
}

// Copyright 2023 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package config

import (
	"fmt"
	"strings"
)

// ConsoleAPI declares the configuration properties for managing the
// connect/grpc/grpc-gateway API endpoints.
type ConsoleAPI struct {
	// Enabled determines whether any of the connect/grpc/grpc-gateway endpoints
	// will be mounted to the server.
	Enabled bool `yaml:"enabled"`

	// EnabledProcedures is a list of procedure names that shall be allowed.
	// If a procedure is called that is not on this list a descriptive error
	// will be returned. A procedure name has the following format, regardless
	// whether it's called via connect, gRPC or the HTTP interface:
	// "/redpanda.api.dataplane.v1alpha1.UserService/ListUsers".
	// You can use "*" to enable all procedures.
	EnabledProcedures []string `yaml:"enabledProcedures"`
}

// Validate configuration options for the Console topic documentation feature.
func (c *ConsoleAPI) Validate() error {
	if !c.Enabled {
		return nil
	}

	for _, p := range c.EnabledProcedures {
		if p == "*" {
			continue
		}
		if !strings.HasPrefix(p, "/") {
			return fmt.Errorf("every procedure must start with a slash. Entry %q does not start with a slash", p)
		}
	}

	return nil
}

// SetDefaults for ConsoleAPI.
func (c *ConsoleAPI) SetDefaults() {
	c.Enabled = true
	c.EnabledProcedures = []string{"*"}
}

// GetEnabledProcedures returns a map where the key represents the
// enabled procedures. The bool indicates whether all procedures are
// allowed. The bool will be true if at least one string of EnabledProcedures
// contains the wildcard ("*").
func (c *ConsoleAPI) GetEnabledProcedures() (map[string]struct{}, bool) {
	allAllowed := false
	procedures := make(map[string]struct{}, len(c.EnabledProcedures))
	for _, p := range c.EnabledProcedures {
		procedures[p] = struct{}{}
		if p == "*" {
			allAllowed = true
		}
	}

	return procedures, allAllowed
}

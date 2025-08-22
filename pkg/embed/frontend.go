// Copyright 2022 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package embed

import "embed"

// FrontendFiles is a directory with all compiled Frontend assets so that they can be served
// by the Go binary via HTTP.
//
//go:embed all:frontend
var FrontendFiles embed.FS

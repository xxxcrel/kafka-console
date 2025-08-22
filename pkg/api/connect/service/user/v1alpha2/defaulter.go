// Copyright 2024 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package user

import v1alpha2 "github.com/xxxcrel/kafka-console/pkg/protogen/redpanda/api/dataplane/v1alpha2"

// Defaulter updates a given user request with defaults.
type defaulter struct{}

func (*defaulter) applyListUsersRequest(req *v1alpha2.ListUsersRequest) {
	if req.GetPageSize() == 0 {
		req.PageSize = 100
	}
}

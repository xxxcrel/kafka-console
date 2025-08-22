// Copyright 2023 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package acl

import (
	"testing"

	"github.com/stretchr/testify/assert"

	v1alpha2 "github.com/xxxcrel/kafka-console/pkg/protogen/redpanda/api/dataplane/v1alpha2"
)

func TestApplyListACLsRequest(t *testing.T) {
	tests := []struct {
		name     string
		input    *v1alpha2.ListACLsRequest
		expected *v1alpha2.ListACLsRequest
	}{
		{
			name:     "nil filter",
			input:    &v1alpha2.ListACLsRequest{},
			expected: &v1alpha2.ListACLsRequest{Filter: &v1alpha2.ListACLsRequest_Filter{ResourceType: v1alpha2.ACL_RESOURCE_TYPE_ANY, ResourcePatternType: v1alpha2.ACL_RESOURCE_PATTERN_TYPE_ANY, Operation: v1alpha2.ACL_OPERATION_ANY, PermissionType: v1alpha2.ACL_PERMISSION_TYPE_ANY}},
		},
		{
			name:     "unspecified values",
			input:    &v1alpha2.ListACLsRequest{Filter: &v1alpha2.ListACLsRequest_Filter{}},
			expected: &v1alpha2.ListACLsRequest{Filter: &v1alpha2.ListACLsRequest_Filter{ResourceType: v1alpha2.ACL_RESOURCE_TYPE_ANY, ResourcePatternType: v1alpha2.ACL_RESOURCE_PATTERN_TYPE_ANY, Operation: v1alpha2.ACL_OPERATION_ANY, PermissionType: v1alpha2.ACL_PERMISSION_TYPE_ANY}},
		},
		{
			name:     "partially specified values",
			input:    &v1alpha2.ListACLsRequest{Filter: &v1alpha2.ListACLsRequest_Filter{Operation: v1alpha2.ACL_OPERATION_CREATE}},
			expected: &v1alpha2.ListACLsRequest{Filter: &v1alpha2.ListACLsRequest_Filter{ResourceType: v1alpha2.ACL_RESOURCE_TYPE_ANY, ResourcePatternType: v1alpha2.ACL_RESOURCE_PATTERN_TYPE_ANY, Operation: v1alpha2.ACL_OPERATION_CREATE, PermissionType: v1alpha2.ACL_PERMISSION_TYPE_ANY}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &defaulter{}
			d.applyListACLsRequest(tt.input)
			assert.Equal(t, tt.expected, tt.input)
		})
	}
}

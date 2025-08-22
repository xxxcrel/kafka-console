// Copyright 2023 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package user

import (
	"fmt"

	adminapi "github.com/redpanda-data/common-go/rpadmin"

	dataplanev1alpha2 "github.com/xxxcrel/kafka-console/pkg/protogen/redpanda/api/dataplane/v1alpha2"
)

// saslMechanismToRedpandaAdminAPIString converts the SASL Mechanism enum into a string that is understood by
// the Redpanda Admin API.
func saslMechanismToRedpandaAdminAPIString(mechanism dataplanev1alpha2.SASLMechanism) (string, error) {
	switch mechanism {
	case dataplanev1alpha2.SASLMechanism_SASL_MECHANISM_SCRAM_SHA_256:
		return adminapi.ScramSha256, nil
	case dataplanev1alpha2.SASLMechanism_SASL_MECHANISM_SCRAM_SHA_512:
		return adminapi.ScramSha512, nil
	default:
		return "", fmt.Errorf("unable to convert %q to a known string that can be handled by the Redpanda Admin API", mechanism.String())
	}
}

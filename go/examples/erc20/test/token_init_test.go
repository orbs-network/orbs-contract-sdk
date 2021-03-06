// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package test

import (
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/orbs-network/orbs-contract-sdk/go/examples/test"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestTokenInit(t *testing.T) {
	sender, _ := orbs.CreateAccount()

	h := newHarness()
	response := h.deployContract(t, sender)
	require.Len(t, response.OutputEvents, 1, "initial transfer (mint) did not fire during init")

	require.True(t, test.Eventually(1*time.Second, func() bool {
		return h.balanceOf(t, sender) == uint64(1000000000000000000)
	}))
}

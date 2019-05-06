// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package test

import (
	"github.com/orbs-network/orbs-contract-sdk/go/examples/test"
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestTokenTransfer(t *testing.T) {
	user1, _ := orbs.CreateAccount()
	user2, _ := orbs.CreateAccount()

	h := newHarness()
	h.deployContract(t, user1)
	require.EqualValues(t, 1000000000000000000, h.balanceOf(t, user1))

	response, err := h.transfer(t, user1, user2, 5000)
	require.NoError(t, err)
	require.EqualValues(t, "Transfer", response.OutputEvents[0].EventName)

	require.True(t, test.Eventually(1*time.Second, func() bool {
		return h.balanceOf(t, user2) == uint64(5000)
	}))

	invalidUser, _ := orbs.CreateAccount()
	invalidUser.Address = "123"
	response, err = h.transfer(t, user1, invalidUser, 5000)

	require.NoError(t, err)
	require.EqualValuesf(t, codec.EXECUTION_RESULT_ERROR_SMART_CONTRACT, response.ExecutionResult, "executing with invalid user worked")
}

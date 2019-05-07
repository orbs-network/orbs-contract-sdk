// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package test

import (
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/orbs-network/orbs-contract-sdk/go/examples/test"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestSuccessfulTransfer(t *testing.T) {
	user1, _ := orbs.CreateAccount()
	user2, _ := orbs.CreateAccount()


	h := newHarness()
	h.deployContract(t, user1)

	require.EqualValues(t, 1000, h.getBalance(t, user1))
	require.EqualValues(t, 0, h.getBalance(t, user2))

	result, err := h.transfer(t, user1, user2, uint64(15))
	require.NoError(t, err)
	require.EqualValues(t, codec.EXECUTION_RESULT_SUCCESS, result.ExecutionResult)


	require.True(t, test.Eventually(1*time.Second, func() bool {
		return uint64(985) == h.getBalance(t, user1)
	}))

	require.True(t, test.Eventually(1*time.Second, func() bool {
		return uint64(15) == h.getBalance(t, user2)
	}))
}

// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package test

import (
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTransferWithInsufficientFunds(t *testing.T) {
	user1, _ := orbs.CreateAccount()
	user2, _ := orbs.CreateAccount()

	h := newHarness()
	h.deployContract(t, user1)

	require.EqualValues(t, 1000, h.getBalance(t, user1))
	require.EqualValues(t, 0, h.getBalance(t, user2))

	result, err := h.transfer(t, user1, user2, uint64(1500))
	require.NoError(t, err)
	require.EqualValues(t, codec.EXECUTION_RESULT_ERROR_SMART_CONTRACT, result.ExecutionResult)

	require.EqualValues(t, 1000, h.getBalance(t, user1))
	require.EqualValues(t, 0, h.getBalance(t, user2))
}

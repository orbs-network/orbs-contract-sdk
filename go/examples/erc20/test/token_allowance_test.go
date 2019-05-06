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

func TestAllowance(t *testing.T) {
	user1, _ := orbs.CreateAccount()
	user2, _ := orbs.CreateAccount()
	user3, _ := orbs.CreateAccount()

	h := newHarness()
	h.deployContract(t, user1)

	t.Run("SimpleAllowance", func (t *testing.T) {
		require.EqualValues(t, 0, h.allowance(t, user1, user2))

		response, err := h.approve(t, user1, user2, 3000)
		require.NoError(t, err)
		require.EqualValues(t, "Approval", response.OutputEvents[0].EventName)

		require.True(t, test.Eventually(1*time.Second, func() bool {
			return uint64(3000) == h.allowance(t, user1, user2)
		}))
	})

	// test relies on setup done in previous
	t.Run("IncreaseAllowance", func (t *testing.T) {
		response, err := h.increaseAllowance(t, user1, user2, 1000)
		require.NoError(t, err)
		require.EqualValues(t, "Approval", response.OutputEvents[0].EventName)
		require.EqualValues(t, 4000, response.OutputEvents[0].Arguments[2], "event data incorrect")

		require.True(t, test.Eventually(1*time.Second, func() bool {
			return uint64(4000) == h.allowance(t, user1, user2)
		}))

		response, err = h.increaseAllowance(t, user1, user2, 18446744073709551615)
		require.NoError(t, err)
		require.EqualValuesf(t, codec.EXECUTION_RESULT_ERROR_SMART_CONTRACT, response.ExecutionResult, "integer overflow on increaseAllowance")
	})

	t.Run("DecreaseAllowance", func (t *testing.T) {
		response, err := h.decreaseAllowance(t, user1, user2, 18446744073709551615)
		require.NoError(t, err)
		require.EqualValuesf(t, codec.EXECUTION_RESULT_ERROR_SMART_CONTRACT, response.ExecutionResult, "integer overflow on decreaseAllowance")

		response, err = h.decreaseAllowance(t, user1, user2, 1000)
		require.NoError(t, err)
		require.EqualValuesf(t, codec.EXECUTION_RESULT_SUCCESS, response.ExecutionResult, "decrease allowance failed")
		require.EqualValues(t, "Approval", response.OutputEvents[0].EventName)
		require.EqualValues(t, 3000, response.OutputEvents[0].Arguments[2], "event data incorrect")


		require.True(t, test.Eventually(1*time.Second, func() bool {
			return uint64(3000) == h.allowance(t, user1, user2)
		}))
	})

	t.Run("AllowedTransfer", func (t *testing.T) {
		response, err := h.approve(t, user1, user2, 3000)
		require.NoError(t, err)
		require.EqualValues(t, "Approval", response.OutputEvents[0].EventName)

		response, err = h.transferFrom(t, user2, user1, user3, 2000)
		require.NoError(t, err)
		require.EqualValues(t, "Approval", response.OutputEvents[0].EventName)
		require.EqualValues(t, "Transfer", response.OutputEvents[1].EventName)

		response, err = h.transferFrom(t, user2, user1, user3, 2000)
		require.NoError(t, err)
		require.EqualValuesf(t, codec.EXECUTION_RESULT_ERROR_SMART_CONTRACT, response.ExecutionResult, "transferred without approval")
	})
}
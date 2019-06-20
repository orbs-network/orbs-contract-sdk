// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package test

import (
	"encoding/json"
	"github.com/orbs-network/orbs-contract-sdk/go/examples/test"
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPhoneBook(t *testing.T) {
	sender, _ := orbs.CreateAccount()

	h := newHarness()
	h.deployContract(t, sender)

	payload := map[string]interface{}{
		"FirstName": "Huckleberry",
		"LastName": "Finn",
		"Phone": 1234567,
	}
	rawJson, _ := json.Marshal(payload)

	result, err := h.register(t, sender, string(rawJson))
	require.NoError(t, err)
	require.EqualValues(t, codec.EXECUTION_RESULT_SUCCESS, result.ExecutionResult)

	require.True(t, test.Eventually(1*time.Second, func() bool {
		value := h.get(t, sender)
		var m map[string]interface{}
		if err = json.Unmarshal([]byte(value.(string)), &m); err == nil {
			return m["FirstName"] == "Huckleberry" && m["LastName"] == "Finn" && m["Phone"] == 1234567.0
		}

		return false
	}))
}

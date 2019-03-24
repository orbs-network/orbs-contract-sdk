// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package test

import (
	"github.com/orbs-network/orbs-contract-sdk/go/testing/gamma"
	"strings"
	"testing"
)

func TestAllowance(t *testing.T) {
	gammaCli := gamma.Cli().Start()
	defer gammaCli.Stop()

	out := gammaCli.Run("deploy ../erc20.go -name OrbsERC20 -signer user1")
	if !strings.Contains(out, `"ExecutionResult": "SUCCESS"`) {
		t.Fatal("deploy failed")
	}

	t.Run("SimpleAllowance", func (t *testing.T) {
		out = gammaCli.Run("run-query allowance-user1-user2.json -signer user1")
		if !strings.Contains(out, `"Value": "0"`) {
			t.Fatal("initial get failed")
		}

		out = gammaCli.Run("send-tx approve-user2-3000.json -signer user1")
		if !strings.Contains(out, `"EventName": "Approval",`) {
			t.Fatal("simple approve failed")
		}

		out = gammaCli.Run("run-query allowance-user1-user2.json -signer user1")
		if !strings.Contains(out, `"Value": "3000"`) {
			t.Fatal("data validation failed, approve did not set to the right amount")
		}
	})

	// test relies on setup done in previous
	t.Run("IncreaseAllowance", func (t *testing.T) {
		out = gammaCli.Run("send-tx increaseAllowance-user2-1000.json -signer user1")
		if !strings.Contains(out, `"EventName": "Approval"`) {
			t.Fatal("increase allowance failed")
		}
		if !strings.Contains(out, `"Value": "4000"`) {
			t.Fatal("event data incorrect")
		}

		out = gammaCli.Run("run-query allowance-user1-user2.json -signer user1")
		if !strings.Contains(out, `"Value": "4000"`) {
			t.Fatal("data validation failed, approve/increase did not set to the right amount")
		}

		out = gammaCli.Run("send-tx increaseAllowance-user2-64max.json -signer user1")
		if !strings.Contains(out, `"ExecutionResult": "ERROR_SMART_CONTRACT"`) {
			t.Fatal("integer overflow on increaseAllowance")
		}
	})

	t.Run("DecreaseAllowance", func (t *testing.T) {
		out = gammaCli.Run("send-tx decreaseAllowance-user2-64max.json -signer user1")
		if !strings.Contains(out, `"ExecutionResult": "ERROR_SMART_CONTRACT"`) {
			t.Fatal("integer overflow on decreaseAllowance")
		}

		out = gammaCli.Run("send-tx decreaseAllowance-user2-1000.json -signer user1")
		if !strings.Contains(out, `"EventName": "Approval"`) {
			t.Fatal("decrease allowance failed")
		}
		if !strings.Contains(out, `"Value": "3000"`) {
			t.Fatal("event data incorrect")
		}

		out = gammaCli.Run("run-query allowance-user1-user2.json -signer user1")
		if !strings.Contains(out, `"Value": "3000"`) {
			t.Fatal("data validation failed, approve/increase did not set to the right amount")
		}
	})

	t.Run("AllowedTransfer", func (t *testing.T) {
		out = gammaCli.Run("send-tx approve-user2-3000.json -signer user1")
		if !strings.Contains(out, `"EventName": "Approval"`) {
			t.Fatal("approval for transfer failed")
		}

		out = gammaCli.Run("send-tx transferFrom-user1-user3.json -signer user2")
		if !strings.Contains(out, `"EventName": "Approval"`) {
			t.Fatal("approval did not update after transfer")
		}
		if !strings.Contains(out, `"EventName": "Transfer"`) {
			t.Fatal("transfer event missing in transfer action")
		}

		out = gammaCli.Run("send-tx transferFrom-user1-user3.json -signer user2")
		if !strings.Contains(out, `"ExecutionResult": "ERROR_SMART_CONTRACT"`) {
			t.Fatal("transferred without approval")
		}
	})
}
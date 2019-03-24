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

func TestTokenTransfer(t *testing.T) {
	gammaCli := gamma.Cli().Start()
	defer gammaCli.Stop()

	out := gammaCli.Run("deploy ../erc20.go -name OrbsERC20 -signer user1")
	if !strings.Contains(out, `"ExecutionResult": "SUCCESS"`) {
		t.Fatal("deploy failed")
	}

	out = gammaCli.Run("send-tx transfer-user2-5000.json -signer user1")
	if !strings.Contains(out, `"EventName": "Transfer",`) {
		t.Fatal("initial get failed")
	}

	out = gammaCli.Run("run-query balanceOf-user2.json")
	if !strings.Contains(out, `"Value": "5000"`) {
		t.Fatal("funds are not present for user2 after transfer")
	}

	out = gammaCli.Run("send-tx transfer-invalid-user.json")
	if !strings.Contains(out, `"ExecutionResult": "ERROR_SMART_CONTRACT"`) {
		t.Fatal("executing with invalid user worked")
	}
}

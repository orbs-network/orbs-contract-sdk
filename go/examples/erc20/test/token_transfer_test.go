package test

import (
	"github.com/orbs-network/orbs-contract-sdk/go/testing/gamma"
	"strings"
	"testing"
)

func TestTokenTransfer(t *testing.T) {
	gammaCli := gamma.Cli().StartExperimental()
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

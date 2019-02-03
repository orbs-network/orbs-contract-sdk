package test

import (
	"github.com/orbs-network/orbs-contract-sdk/go/testing/gamma"
	"strings"
	"testing"
)

func TestAllow(t *testing.T) {
	gammaCli := gamma.Cli().StartExperimental()
	defer gammaCli.Stop()

	out := gammaCli.Run("deploy ../erc20.go -name OrbsERC20 -signer user1")
	if !strings.Contains(out, `"ExecutionResult": "SUCCESS"`) {
		t.Fatal("deploy failed")
	}

	out = gammaCli.Run("run-query allowance-user1-user2.json -signer user1")
	if !strings.Contains(out, `"Value": "0"`) {
		t.Fatal("initial get failed")
	}

	out = gammaCli.Run("send-tx approve-user2-3000.json -signer user1")
	if !strings.Contains(out, `"EventName": "Approval",`) {
		t.Fatal("initial get failed")
	}

	out = gammaCli.Run("run-query allowance-user1-user2.json -signer user1")
	if !strings.Contains(out, `"Value": "3000"`) {
		t.Fatal("initial get failed")
	}
}


func TestAllowedTransfer(t *testing.T) {
	gammaCli := gamma.Cli().StartExperimental()
	defer gammaCli.Stop()

	out := gammaCli.Run("deploy ../erc20.go -name OrbsERC20 -signer user1")
	if !strings.Contains(out, `"ExecutionResult": "SUCCESS"`) {
		t.Fatal("deploy failed")
	}

	out = gammaCli.Run("send-tx transferFrom-user1-user3.json -signer user2")
	if !strings.Contains(out, `"ExecutionResult": "ERROR_SMART_CONTRACT"`) {
		t.Fatal("transferred without approval")
	}

	out = gammaCli.Run("send-tx approve-user2-3000.json -signer user1")
	if !strings.Contains(out, `"EventName": "Approval"`) {
		t.Fatal("initial get failed")
	}

	out = gammaCli.Run("send-tx transferFrom-user1-user3.json -signer user2")
	if !strings.Contains(out, `"EventName": "Approval"`) {
		t.Fatal("initial get failed")
	}
	if !strings.Contains(out, `"EventName": "Transfer"`) {
		t.Fatal("initial get failed")
	}
}
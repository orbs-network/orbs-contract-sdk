package test

import (
	"github.com/orbs-network/orbs-contract-sdk/go/testing/gamma"
	"strings"
	"testing"
)

func TestTransferWithInsufficientFunds(t *testing.T) {
	gammaCli := gamma.Cli().Start()
	defer gammaCli.Stop()

	out := gammaCli.Run("deploy ../contract.go -name MySimpleToken -signer user1")
	if !strings.Contains(out, `"ExecutionResult": "SUCCESS"`) {
		t.Fatal("deploy failed")
	}

	out = gammaCli.Run("run-query get-user1-balance.json")
	if !strings.Contains(out, `"Value": "1000"`) {
		t.Fatal("initial user1 balance failed")
	}

	out = gammaCli.Run("run-query get-user2-balance.json")
	if !strings.Contains(out, `"Value": "0"`) {
		t.Fatal("initial user2 balance failed")
	}

	out = gammaCli.Run("send-tx transfer-1500-to-user2.json -signer user1")
	if !strings.Contains(out, `"ExecutionResult": "ERROR_SMART_CONTRACT"`) {
		t.Fatal("transfer failed")
	}

	out = gammaCli.Run("run-query get-user1-balance.json")
	if !strings.Contains(out, `"Value": "1000"`) {
		t.Fatal("final user1 balance failed")
	}

	out = gammaCli.Run("run-query get-user2-balance.json")
	if !strings.Contains(out, `"Value": "0"`) {
		t.Fatal("final user2 balance failed")
	}
}

package test

import (
	"github.com/orbs-network/orbs-contract-sdk/go/testing/gamma"
	"strings"
	"testing"
)

func TestTokenInit(t *testing.T) {
	gammaCli := gamma.Cli().Start()
	defer gammaCli.Stop()

	out := gammaCli.Run("deploy ../erc20.go -name OrbsERC20 -signer user1")
	if !strings.Contains(out, `"ExecutionResult": "SUCCESS"`) {
		t.Fatal("deploy failed")
	}

	if !strings.Contains(out, `"EventName": "Transfer",`) {
		t.Fatal("initial transfer (mint) did not fire during init")
	}

	out = gammaCli.Run("run-query balanceOf-user1.json")
	if !strings.Contains(out, `"Value": "1000000000000000000"`) {
		t.Fatal("initial get failed")
	}
}

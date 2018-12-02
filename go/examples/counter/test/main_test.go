package test

import (
	"github.com/orbs-network/orbs-contract-sdk/go/testing/gamma"
	"strings"
	"testing"
)

func TestCounterIncrement(t *testing.T) {
	gammaCli := gamma.Cli().Start()
	defer gammaCli.Stop()

	out := gammaCli.Run("deploy -name MyCounter -code ../contract.go")
	if !strings.Contains(out, `"ExecutionResult": "SUCCESS"`) {
		t.Fatal("deploy failed")
	}

	out = gammaCli.Run("read -i get.json")
	if !strings.Contains(out, `"Value": "0"`) {
		t.Fatal("initial get failed")
	}

	out = gammaCli.Run("send-tx -i add-25.json")
	if !strings.Contains(out, `"ExecutionResult": "SUCCESS"`) {
		t.Fatal("first add 25 failed")
	}

	out = gammaCli.Run("send-tx -i add-25.json")
	if !strings.Contains(out, `"ExecutionResult": "SUCCESS"`) {
		t.Fatal("second add 25 failed")
	}

	out = gammaCli.Run("send-tx -i add-25.json")
	if !strings.Contains(out, `"ExecutionResult": "SUCCESS"`) {
		t.Fatal("third add 25 failed")
	}

	out = gammaCli.Run("read -i get.json")
	if !strings.Contains(out, `"Value": "75"`) {
		t.Fatal("final get failed")
	}
}

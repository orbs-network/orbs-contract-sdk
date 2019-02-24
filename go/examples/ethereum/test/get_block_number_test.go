package test

import (
	"github.com/orbs-network/orbs-contract-sdk/go/testing/ethereum"
	"github.com/orbs-network/orbs-contract-sdk/go/testing/gamma"
	"strings"
	"testing"
)

func TestGetBlockNumberOnEthereum(t *testing.T) {
	gammaCli := gamma.Cli().Start()
	defer gammaCli.Stop()

	truffleCli := truffle.Cli("./EthereumContract")
	truffleCli.Run("exec deploy-new.js") // so we have at least one closed block

	out := gammaCli.Run("deploy ../contract.go -name MyCrossChain")
	if !strings.Contains(out, `"ExecutionResult": "SUCCESS"`) {
		t.Fatal("deploy failed")
	}

	out = gammaCli.Run("run-query get-block-number.json")
	//TODO(v1): add expectations for block number
	t.Log(out)
}

package test

import (
	"github.com/orbs-network/orbs-contract-sdk/go/testing/ethereum"
	"github.com/orbs-network/orbs-contract-sdk/go/testing/gamma"
	"regexp"
	"strings"
	"testing"
)

func TestCallMethodOnEthereum(t *testing.T) {
	gammaCli := gamma.Cli().Start()
	defer gammaCli.Stop()

	truffleCli := truffle.Cli("./EthereumContract")
	out := truffleCli.Run("exec deploy-new.js")
	ethContractAddress := extractAddressFromTruffleOutput(out)

	out = gammaCli.Run("deploy ../contract.go -name MyCrossChain")
	if !strings.Contains(out, `"ExecutionResult": "SUCCESS"`) {
		t.Fatal("deploy failed")
	}

	out = gammaCli.Run("run-query call-hello-world.json -arg1 " + ethContractAddress)
	if !strings.Contains(out, `"Value": "hello world"`) {
		t.Fatal("crosschain call failed")
	}
}

func extractAddressFromTruffleOutput(out string) string {
	re := regexp.MustCompile(`\"ContractAddress\":\s+\"(\w+)\"`)
	res := re.FindStringSubmatch(out)
	return res[1]
}

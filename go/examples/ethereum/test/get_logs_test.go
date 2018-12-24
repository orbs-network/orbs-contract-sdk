package test

import (
	"github.com/orbs-network/orbs-contract-sdk/go/testing/ethereum"
	"github.com/orbs-network/orbs-contract-sdk/go/testing/gamma"
	"regexp"
	"strings"
	"testing"
)

func TestGetLogsOnEthereum(t *testing.T) {
	gammaCli := gamma.Cli().Start()
	defer gammaCli.Stop()

	truffleCli := truffle.Cli("./EthereumContract")
	out := truffleCli.Run("exec emit-event.js")
	ethContractAddress := extractAddressFromTruffleOutput(out)
	ethTxHash := extractTxHashFromTruffleOutput(out)

	out = gammaCli.Run("deploy ../contract.go -name MyCrossChain")
	if !strings.Contains(out, `"ExecutionResult": "SUCCESS"`) {
		t.Fatal("deploy failed")
	}

	out = gammaCli.Run("run-query get-hello-said-log.json -arg1 " + ethContractAddress + " -arg2 " + ethTxHash)
	if !strings.Contains(out, `"Value": "John Snow"`) {
		t.Fatal("crosschain get logs failed")
	}
}

func extractTxHashFromTruffleOutput(out string) string {
	re := regexp.MustCompile(`\"TxHash\":\s+\"(\w+)\"`)
	res := re.FindStringSubmatch(out)
	return res[1]
}

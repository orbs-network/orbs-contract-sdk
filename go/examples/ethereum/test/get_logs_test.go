// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package test

import (
	"github.com/orbs-network/orbs-contract-sdk/go/testing/ethereum"
	"github.com/orbs-network/orbs-contract-sdk/go/testing/gamma"
	"regexp"
	"strings"
	"testing"
)

func TestGetLogsOnEthereum(t *testing.T) {
	t.Skip("Skipped because it uses Ethereum")
	gammaCli := gamma.Cli()

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
	//TODO(v1): add expectations for block number and txIndex
	t.Log(out)
}

func extractTxHashFromTruffleOutput(out string) string {
	re := regexp.MustCompile(`\"TxHash\":\s+\"(\w+)\"`)
	res := re.FindStringSubmatch(out)
	return res[1]
}

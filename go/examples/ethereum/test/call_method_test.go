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

func TestCallMethodOnEthereum(t *testing.T) {
	t.Skip("skipped because it uses Ethreum")
	gammaCli := gamma.Cli()

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

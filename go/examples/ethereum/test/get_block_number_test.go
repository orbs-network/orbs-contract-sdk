// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package test

import (
	"github.com/orbs-network/orbs-contract-sdk/go/testing/ethereum"
	"github.com/orbs-network/orbs-contract-sdk/go/testing/gamma"
	"strings"
	"testing"
)

func TestGetBlockNumberOnEthereum(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping on CI")
	}
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

// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package test

import (
	"github.com/orbs-network/orbs-contract-sdk/go/testing/gamma"
	"strings"
	"testing"
)

func TestCounterIncrement(t *testing.T) {
	gammaCli := gamma.Cli().Start()
	defer gammaCli.Stop()

	out := gammaCli.Run("deploy ../contract.go -name MyCounter")
	if !strings.Contains(out, `"ExecutionResult": "SUCCESS"`) {
		t.Fatal("deploy failed")
	}

	out = gammaCli.Run("run-query get.json")
	if !strings.Contains(out, `"Value": "0"`) {
		t.Fatal("initial get failed")
	}

	out = gammaCli.Run("send-tx add-25.json")
	if !strings.Contains(out, `"ExecutionResult": "SUCCESS"`) {
		t.Fatal("first add 25 failed")
	}

	out = gammaCli.Run("send-tx add-25.json")
	if !strings.Contains(out, `"ExecutionResult": "SUCCESS"`) {
		t.Fatal("second add 25 failed")
	}

	out = gammaCli.Run("send-tx add-25.json")
	if !strings.Contains(out, `"ExecutionResult": "SUCCESS"`) {
		t.Fatal("third add 25 failed")
	}

	out = gammaCli.Run("run-query get.json")
	if !strings.Contains(out, `"Value": "75"`) {
		t.Fatal("final get failed")
	}
}

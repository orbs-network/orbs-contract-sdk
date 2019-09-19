// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package main

import (
	"fmt"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/address"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/state"
)

var PUBLIC = sdk.Export(transfer, getBalance)
var SYSTEM = sdk.Export(_init)

const TOTAL_SUPPLY = 1000

func _init() {
	ownerAddress := address.GetSignerAddress()
	state.WriteUint64(ownerAddress, TOTAL_SUPPLY)
}

func transfer(amount uint64, targetAddress []byte) {
	// sender
	callerAddress := address.GetCallerAddress()
	callerBalance := state.ReadUint64(callerAddress)
	if callerBalance < amount {
		panic(fmt.Sprintf("transfer of %d failed since balance is only %d", amount, callerBalance))
	}
	state.WriteUint64(callerAddress, callerBalance-amount)

	// recipient
	address.ValidateAddress(targetAddress)
	targetBalance := state.ReadUint64(targetAddress)
	state.WriteUint64(targetAddress, targetBalance+amount)
}

func getBalance(targetAddress []byte) uint64 {
	address.ValidateAddress(targetAddress)
	return state.ReadUint64(targetAddress)
}

package main

import (
	"github.com/orbs-network/orbs-contract-sdk/go/sdk"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/ethereum"
)

var PUBLIC = sdk.Export(readValueFromLog)
var SYSTEM = sdk.Export(_init)

type event struct {
	value string
}

func _init() {
}

func readValueFromLog(address string, abi string, txid string, eventName string) string {
	var event event
	ethereum.GetTransactionLog(address, abi, txid, eventName, &event)
	return event.value
}

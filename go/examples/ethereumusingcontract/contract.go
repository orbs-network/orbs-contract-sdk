package main

import (
	"github.com/orbs-network/orbs-contract-sdk/go/sdk"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/ethereum"
)

var PUBLIC = sdk.Export(readValueFromLog)
var SYSTEM = sdk.Export(_init)

type event struct {
	Value string
}

func _init() {
}

func readValueFromLog(address string, abi string, txid string, eventName string) string {
	var event event
	ethereum.GetTransactionLog(address, abi, txid, eventName, &event)
	return event.Value
}

func callEthereumMethod(address string, abi string, methodName string, args ...interface{}) string {
	var event event
	ethereum.CallMethod(address, abi, methodName, &event, args...)
	return event.Value
}
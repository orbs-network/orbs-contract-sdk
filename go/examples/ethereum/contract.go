package main

import (
	"github.com/orbs-network/orbs-contract-sdk/go/sdk"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/ethereum"
	"math/big"
	"strings"
)

var PUBLIC = sdk.Export(readValueFromLog, callEthereumHelloWorld, getEthereumHelloSaidLog)
var SYSTEM = sdk.Export(_init)

type event struct {
	Value string
}

func _init() {
}

// TODO(talkol): this should not take arguments for simplicity, see new function used for gamma test
func readValueFromLog(address string, abi string, txid string, eventName string) string {
	var event event
	ethereum.GetTransactionLog(address, abi, txid, eventName, &event)
	return event.Value
}

// TODO(talkol): this should not take arguments for simplicity, see new function used for gamma test
func callEthereumMethod(address string, abi string, methodName string, args ...interface{}) string {
	var event event
	ethereum.CallMethod(address, abi, methodName, &event, args...)
	return event.Value
}

func callEthereumHelloWorld(ethContractAddress string) string {
	jsonAbi := `
	[
    {
      "constant": true,
      "inputs": [],
      "name": "sayHello",
      "outputs": [
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "stateMutability": "pure",
      "type": "function"
    }
  ]
	`
	var output string
	ethereum.CallMethod(ethContractAddress, jsonAbi, "sayHello", &output)
	return output
}

type HelloSaidEthereumEvent struct {
	Name   [10]byte
	Amount *big.Int
}

func getEthereumHelloSaidLog(ethContractAddress string, ethTxHash string) string {
	jsonAbi := `
	[
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "name",
          "type": "bytes10"
        },
        {
          "indexed": false,
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "HelloSaid",
      "type": "event"
    }
  ]
	`
	event := HelloSaidEthereumEvent{}
	ethereum.GetTransactionLog(ethContractAddress, jsonAbi, ethTxHash, "HelloSaid", &event)
	return nullTermString(event.Name[:])
}

func nullTermString(cstr []byte) string {
	return strings.TrimRight(string(cstr), "\x00")
}

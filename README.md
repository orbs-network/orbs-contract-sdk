# Orbs Smart Contract SDK

Orbs Smart Contract SDK is a framework for building  applications over Orbs blockchain.
Here you will find everything you need to develop smart contracts for the Orbs blockchain
> **Note** the orbs smart contract SDK is stable, but still under development. There may be breaking changes.

## Table of Contents

* [Overview](#Overview)
* [Quick start](#Quick-Start)
* [Deploying your first contract](#Deploying-your-first-contract)
* [Next steps](#Next-steps)
&nbsp;

## Overview

Smart contracts are the basic building block of blockchain which make the decentralized applications that run on top.</br>
The smart contract SDK allows you to develop, test and deploy contracts so that users can interact with them on-chain using transactions.
To make the development process as easy and productive as possible, Orbs relies on familiar programming languages with established toolchains such as [Go](https://en.wikipedia.org/wiki/Go_(programming_language)).</br>
> The smart contract SDK is  written in Go language. </br>
Please [contact us](FeatureRequest@orbs.com ) for information about additional implementation in another language.
&nbsp;

## Quick Start
&nbsp;

### Prerequisites

- For mac - [brew](https://brew.sh/) should be installed.
- [Go](https://golang.org/doc/install) 1.10+ is installed. 
   
   > Verify installation by typing in terminal `go version`<br/> 
   >The full guide to install Go can be found [here](https://golang.org/doc/install)). 

- Please Make sure [Docker](https://docs.docker.com/docker-for-mac/install/) is installed on your machine.

### Installation 

1. Download the Orbs smart contract SDK by typing in terminal `go get -u github.com/orbs-network/orbs-contract-sdk/...`</br>
   > If you're new to Go, it is recommended to download the SDK to the Go workspace.

   > The SDK includes [examples](https://github.com/orbs-network/orbs-contract-sdk/tree/master/go/examples) - those
     will be download to  ~/github.com/orbs-network/orbs-contract-sdk/go/examples on your computer.
   
2. Install [Gamma]((GAMMA.md))- the local orbs blockchain and CLI tool. Simply type in the terminal
```
  brew install orbs-network/devtools/gamma-cli

```
> Verify by typing in the terminal: `gamma-cli version`

## Deploying your first contract

### 1. Write a simple contract

Let's write a simple counter contract.
This will be our code [`counter.go`](https://github.com/orbs-network/orbs-contract-sdk/blob/master/go/examples/counter/counter.go)
```go
package main

import (
    "github.com/orbs-network/orbs-contract-sdk/go/sdk"
    "github.com/orbs-network/orbs-contract-sdk/go/sdk/state"
)

var PUBLIC = sdk.Export(add, get)
var SYSTEM = sdk.Export(_init)

func _init() {
    state.WriteUint64ByKey("count", 0)
}

func add(amount uint64) {
    count := state.ReadUint64ByKey("count")
    count += amount
    state.WriteUint64ByKey("count", count)
}

func get() uint64 {
    return state.ReadUint64ByKey("count")
}
```

### 2. Start the Gamma server local blockchain 

Type in the terminal:
```
    gamma-cli start-local
```
> A message that Orbs Gamma blockchain is running with the port it is listening to.
> To stop the gamma server, type command `gamma-cli stop-local`

### 2. Deploy the contract

To deploy the counter contract type in the terminal:
```
gamma-cli deploy -name MyCounter -code counter.go
```

> For a successful deploy, response should contain **`"ExecutionResult": "SUCCESS"`**.

> Output example:
```go
{
  "RequestStatus": "COMPLETED",
  "TxId": "7Y4urVmKvunYsxh7kKhUoQ72XjSJcdkBxxzBcauC9icC9gzMy8mPDcg",
  "ExecutionResult": "SUCCESS",
  "OutputArguments": [],
  "TransactionStatus": "COMMITTED",
  "BlockHeight": "1869",
  "BlockTimestamp": "2018-12-05T13:05:51.347Z"
}
```
### 3. Send a transaction to increment the counter

Write the contract's method arguments in a json file named [`add-25.json`](https://github.com/orbs-network/orbs-contract-sdk/blob/master/go/examples/counter/test/add-25.json). See example: 

```json
{
  "ContractName": "MyCounter",
  "MethodName": "add", 
  "Arguments": [
    {
      "Type": "uint64",
      "Value": "25"
    }
  ]
}
```

To increment the counter by 75, for example, please type in the terminal:
```
// sending 
gamma-cli send-tx -i add-25.json -signer user1
gamma-cli send-tx -i add-25.json -signer user1
gamma-cli send-tx -i add-25.json -signer user1
```
>  `send-tx` is used when the smart contract method may change contract's state. 
   The transaction will be added to the blockchain under consensus.
>  `user 1`, from orbs-test-keys.json file, represent the private key that signed the transaction.

### 4. Read the counter value

Write the contract's method arguments in a json file named [`get.json`](https://github.com/orbs-network/orbs-contract-sdk/blob/master/go/examples/counter/test/get.json):

```json
{
  "ContractName": "MyCounter",
  "MethodName": "get",
  "Arguments": []
}
```

To read the counter value from contract's state type in terminal:
```
gamma-cli read -i get.json
```
&nbsp;

## Next steps

- Explore more contract examples in Go  - https://github.com/orbs-network/orbs-contract-sdk/tree/master/go/examples.
- Explore Gamma - local Orbs blockchain - [link](https://github.com/orbs-network/orbs-contract-sdk/blob/master/GAMMA.md) 
    >you can also type in the terminal `gamma-cli help`
- Explore the API of the SDK - [link](https://github.com/orbs-network/orbs-contract-sdk/tree/master/go/sdk)
- client SDK Go - [more information](https://github.com/orbs-network/orbs-client-sdk-go)


## License

MIT

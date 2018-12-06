# Orbs Smart Contract SDK

Orbs Smart Contract Software development kit to empower developers to build, test and deploy smart contracts for their applications over Orbs blockchain.
Here you will find everything you need to develop smart contracts for the Orbs blockchain
> Please note the smart contract sdk development is in progress. 

## Table of Contents

* [Overview](#Overview)
* [Quick start](#Quick-Start)
* [Deploying your first contract](#Deploying-your-first-contract)
* [Next steps](#Next-steps)

&nbsp;

## Overview

Smart contracts are the basic building block of blockchain which make the decentralized applications that run on top.<enter>
The smart contract SDK allows you to develop, test and deploy contracts so that users can interact with them on-chain using transactions.
To make the development process as easy and productive as possible, Orbs relies on familiar programming languages with established toolchains such as [Go](https://en.wikipedia.org/wiki/Go_(programming_language)).<enter>
> Orbs Smart Contract SDK is currently written in Go language,  it could be easily implemented in additional programming languages.<enter>
Please [contact us](FeatureRequest@orbs.com ) for information about additional implementation in another language.
&nbsp;

Orbs smart contracts SDK includes the following capabilities:
-  Orbs smart contracts reading data from Ethereum.
-  [TBD]  TODO:
&nbsp;

## Quick Start
&nbsp;

### Prerequisites
 
1. [Go](https://golang.org/doc/install) 1.10+ is installed. 
   
   > Verify installation with `go version`<br/> 
   > Make sure  [GOPATH](https://github.com/golang/go/wiki/SettingGOPATH) variable is set to the location of your Go workspace.<enter> 
      Use command `echo $GOPATH` .<br/> 
   >The full guide to install Go can be found [here]((https://golang.org/doc/install)) 

2. Please Make sure [Docker](https://docs.docker.com/docker-for-mac/install/) is installed on your machine.

### Installation 

1. Download the Orbs smart contract SDK to your desired location by typing in terminal `go get -u github.com/orbs-network/orbs-contract-sdk/...`
>If you're new to Go, it is recommended to download the SDK to the Go workspace.
> The SDK includes both source code and code example, examples can be found ~/github.com/orbs-network/orbs-contract-sdk/go/examples or [online](https://github.com/orbs-network/orbs-contract-sdk/tree/master/go/examples).

2. Install [Gamma]((GAMMA.md))  - the local orbs blockchain and CLI - simply type this command in your terminal using [brew](https://brew.sh/):
```
         brew install orbs-network/devtools/gamma-cli

```
> Verify by running the command  `gamma-cli version`


## Deploying your first contract

### 1. Write a simple counter contract

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

### 2. Start the Gamma local blockchain 

Type in the terminal:
```go
    gamma-cli start local
```
> you should get a message that Orbs Gamma blockchain is running with the port it is listening to.
> to stop the gamma server, type command `gamma-cli stop-local`

### 2. Deploy the contract

Type in the terminal the command:
```
gamma-cli deploy -name MyCounter -code counter.go
```

> For a successful deploy,  response should contain **`"ExecutionResult": "SUCCESS"`**.
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

The smart contract arguments for are provided in a JSON format.
* The JSON file includes the contract name and input arguments.

Write the contract's method arguments in a json file named `add-25.json` 

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

To increment the counter by 75 type in the terminal the commands:

```
gamma-cli send-tx -i add-25.json
gamma-cli send-tx -i add-25.json
gamma-cli send-tx -i add-25.json
```
>  'send-tx' is used when the smart contract method  may change the the contract state. The transaction will be added to the blockchain under consensus.


### 4. Read the counter value

Write the contract's method arguments in a json file named `get.json`

```json
{
  "ContractName": "MyCounter",
  "MethodName": "get",
  "Arguments": []
}
```

To read the counter value from the smart contract state type in terminal:
```
gamma-cli read -i get.json
```

&nbsp;

## Next steps

- Explore more contract examples in Go - https://github.com/orbs-network/orbs-contract-sdk/tree/master/go/examples.
- Explore Gamma - local Orbs blockchain - [link](https://github.com/orbs-network/orbs-contract-sdk/blob/master/GAMMA.md) 
    >you can also type in the terminal `gamma-cli help`
- Explore the API of the SDK - [link](https://github.com/orbs-network/orbs-contract-sdk/tree/master/go/sdk)
- client SDK Go - [more information](https://github.com/orbs-network/orbs-client-sdk-go)


## License

MIT

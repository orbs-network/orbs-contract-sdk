# Orbs Smart Contract SDK

Orbs smart contract software development kit allows you to develop, test and deploy contracts on the Orbs network so that users can interact with them on-chain using transactions.
> Here you may find everything you need to develop smart contracts for the Orbs blockchain

* [Overview](#Overview)
* [Getting started](#Getting-started)
* [Getting started - TLDR](#Getting-started---TLDR)
* [Deploying your first contract](#Deploying-your-first-contract)
* [Next steps](#Next-steps)

&nbsp;
## Overview

Smart contracts are the basic building block of blockchain which make the decentralized applications that run on top. The smart contract SDK allows you to develop, test and deploy contracts so that users can interact with them on-chain using transactions.

To make the development process as easy and productive as possible, Orbs relies on familiar programming languages with established toolchains such as [Go](https://en.wikipedia.org/wiki/Go_(programming_language)) and [JavaScript](https://en.wikipedia.org/wiki/JavaScript).

&nbsp;
## Getting started

In order to start developing, you will need the following:

1. Choose your preferred language - such as Go or JavaScript
2. Install the standard dev tools for this language on your machine - IDE, compiler
3. Download the source code and examples in the contract SDK - this repo 
4. Install Gamma - the local Orbs blockchain instance for testing your code

### 1. Choose your preferred language

Go is a low-level open source language developed by Google for extreme performance. It is very popular in the blockchain space and was used to implement the node core of popular blockchains like Ethereum and Hyperledger. It rivals the performance of C++ but is substantially simpler and easier to learn.

JavaScript is the world's most popular programming language. It's a high-level open language designed for isolation and productivity. JavaScript shines in environments where anyone can deploy code and this code must be well isolated to maintain security. JavaScript support in Orbs is currently under development.

### 2. Install the standard dev tools

##### Go developers:

1. Install Go with [brew](https://brew.sh/) by typing in terminal `brew install go`
2. Verify the installation by typing in terminal `go version`
3. Go creates a workspace for you to work in, usually in your home directory at `~/go/src`
4. For more details see the official [go installation guide](https://golang.org/doc/install) and [how to write go code](https://golang.org/doc/code.html) guide

### 3. Download the contract SDK

##### Go developers:

1. Run in terminal `go get -u github.com/orbs-network/orbs-contract-sdk`
2. The SDK will be downloaded to your workspace at `~/go/src/github.com/orbs-network/orbs-contract-sdk`
3. If you're new to Go, keep the SDK in the workspace since Go is particular about the location of source files
4. Choose an [IDE](https://golang.org/doc/editors.html) and open it in `~/go/src/github.com/orbs-network/orbs-contract-sdk/go`

### 4. Install Gamma

##### Go developers:

1. Install the CLI with [brew](https://brew.sh/) by typing in terminal `brew install orbs-network/devtools/gamma-cli`
2. Verify the installation by typing in terminal `gamma-cli version`
3. Start Gamma server with `gamma-cli start-local` and stop with `gamma-cli stop-local`
4. For more details see the full Gamma [documentation](GAMMA.md)

&nbsp;
## Getting started - TLDR

Or just do all 4 steps above at once

##### Go developers:

```
brew install go
go get -u github.com/orbs-network/orbs-contract-sdk
cd ~/go/src/github.com/orbs-network/orbs-contract-sdk/go
brew cask install atom
apm install go-plus
atom ~/go/src/github.com/orbs-network/orbs-contract-sdk/go
brew install orbs-network/devtools/gamma-cli
```

&nbsp;
## Deploying your first contract

### 1. Write a simple counter contract

##### Go developers:

This will be our code `counter.go`

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

### 2. Deploy the contract

Use `gamma-cli` command line tool to deploy

```
gamma-cli start-local
gamma-cli deploy -name MyCounter -code counter.go
```

### 3. Send a transaction to increment the counter

This will be our command `add-25.json` 

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

Use `gamma-cli` command line tool to send it a few times

```
gamma-cli send-tx -i add-25.json
gamma-cli send-tx -i add-25.json
gamma-cli send-tx -i add-25.json
```

### 4. Read the counter value

This will be our command `get.json`

```json
{
  "ContractName": "MyCounter",
  "MethodName": "get",
  "Arguments": []
}
```

Use `gamma-cli` command line tool to send it

```
gamma-cli read -i get.json
```

&nbsp;
## Next steps

### 1. Explore more contract examples

##### Go developers:

https://github.com/orbs-network/orbs-contract-sdk/tree/master/go/examples

### 2. Explore the API of the SDK

##### Go developers:

https://github.com/orbs-network/orbs-contract-sdk/tree/master/go/sdk

### 3. Explore Gamma - local Orbs blockchain

```
gamma-cli help
```

### 4. Explore the client SDK

##### Go developers:

https://github.com/orbs-network/orbs-client-sdk-go

&nbsp;
## License

MIT

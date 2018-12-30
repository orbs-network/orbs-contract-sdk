# Orbs Smart Contract SDK

The Orbs smart contract SDK is a framework for building decentralized applications over the Orbs blockchain. These applications are made of smart contracts written in the [Go](https://en.wikipedia.org/wiki/Go_(programming_language)) programming language.

> Note: While the Orbs smart contract SDK is stable, it is still under active development; there may be breaking changes.

Support for additional programming languages like JavaScript is under way, [contact us](FeatureRequest@orbs.com) for more information.

&nbsp;

## Table of contents

* [Quick start](#quick-start)
* [Deploying your first contract](#deploying-your-first-contract)
* [Next steps](#next-steps)
* [Detailed documentation](https://orbs.gitbook.io)

&nbsp;

## Quick start

### Prerequisites (Mac)

* Make sure [brew](https://brew.sh/) is available on your machine.

* Make sure [Go language](https://golang.org/doc/install) 1.10+ is installed on your machine.
   
    > Verify the installation by running in terminal `go version`

* To test contracts locally, make sure [Docker](https://docs.docker.com/docker-for-mac/install/) is installed on your machine.

### Installation 

1. Download the Orbs smart contract SDK by running in terminal:

    ```
    go get -u github.com/orbs-network/orbs-contract-sdk/...
    ```

   > It will be downloaded to your Go workspace, typically `~/go/src/github.com/orbs-network/orbs-contract-sdk`<br>[Example contracts](https://github.com/orbs-network/orbs-contract-sdk/tree/master/go/examples) will be at `~/go/src/github.com/orbs-network/orbs-contract-sdk/go/examples`

2. Install [Gamma](GAMMA.md), a personal Orbs blockchain running locally, by running in terminal:

    ```
    brew install orbs-network/devtools/gamma-cli
    ```
    
    > To verify the installation, run in terminal `gamma-cli version`

&nbsp;

## Deploying your first contract

### 1. Write a simple contract

Let's write a simple example that implements a counter. This will be our code [`counter.go`](https://github.com/orbs-network/orbs-contract-sdk/blob/master/go/examples/counter/counter.go)

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

### 2. Start Gamma server

We'll test our contract on Gamma server - a locally running blockchain. Start it from terminal:

```
gamma-cli start-local
```

### 3. Deploy the contract

To deploy the counter contract, run in terminal:

```
gamma-cli deploy counter.go -name MyCounter
```

> If the deploy is successful, you'll see a response similar to this:

```json
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

### 4. Send a transaction to increment the counter

Write the transaction details in a JSON file named [`add-25.json`](https://github.com/orbs-network/orbs-contract-sdk/blob/master/go/examples/counter/test/add-25.json)

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

To increment the counter by 75, let's send this transaction 3 times from terminal:

```
gamma-cli send-tx add-25.json -signer user1
gamma-cli send-tx add-25.json -signer user1
gamma-cli send-tx add-25.json -signer user1
```

> Note: The transaction will be signed by `user1`, an example account found in `orbs-test-keys.json`

### 5. Read the counter value

Write the query details in a JSON file named [`get.json`](https://github.com/orbs-network/orbs-contract-sdk/blob/master/go/examples/counter/test/get.json)

```json
{
  "ContractName": "MyCounter",
  "MethodName": "get",
  "Arguments": []
}
```

This query will read the counter value from the contract's state. Send it from terminal:

```
gamma-cli run-query get.json
```

> Note: Transactions that change state require consensus by several nodes. Reading state with queries is a simpler action that doesn't require consensus.

### 6. Stop Gamma server

Since we're done testing, the server is no longer needed. Let's stop it from terminal:

```
gamma-cli stop-local
```

&nbsp;

## Next steps

* Explore more examples of contracts [here](https://github.com/orbs-network/orbs-contract-sdk/tree/master/go/examples).

* Read more about Gamma, the local Orbs blockchain, [here](https://github.com/orbs-network/orbs-contract-sdk/blob/master/GAMMA.md).

    > You can also run in terminal `gamma-cli help`
    
* Explore the API of the SDK [here](https://github.com/orbs-network/orbs-contract-sdk/tree/master/go/sdk).

* After your contracts are deployed to test net or main net, build clients that access them using the [Orbs Client SDK](https://github.com/orbs-network/orbs-client-sdk-go).

&nbsp;

## Detailed documentation

The detailed documentation website for Orbs Contract SDK is available here:

https://orbs.gitbook.io

&nbsp;

## License

MIT

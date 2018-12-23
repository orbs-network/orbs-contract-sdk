# Gamma - Personal Orbs Blockchain for Developers 

Gamma is a personal Orbs blockchain that allows developers to easily test, run and deploy smart contracts.

`Gamma server` - runs an in-memory virtual chain on top of an Orbs blockchain with several nodes on your local machine.

`gamma-cli` - command line tool for developers to interact with a Gamma server instance running on their machine. 

&nbsp;

## Getting started

### Prerequisites (Mac)

* Make sure [brew](https://brew.sh/) is available on your machine.

* Make sure [Docker](https://docs.docker.com/docker-for-mac/install/) is installed on your machine.

* If you're planning to develop your own smart contracts in Go, install the [Orbs Smart Contract SDK](https://github.com/orbs-network/orbs-contract-sdk#installation).

&nbsp;

## Installation

1. To install the command line tool, run the following command in terminal:

    ```
    brew install orbs-network/devtools/gamma-cli
    ```
    
    > To verify the installation, run in terminal `gamma-cli version`
    
2. Gamma server will automatically be installed the first time you start it with `gamma-cli start-local`

&nbsp;

## Starting and stopping Gamma server 

* Start Gamma server by running in terminal:

    ```
    gamma-cli start-local
    ```

* When finished working with the server, stop it by running in terminal:

    ```
    gamma-cli stop-local
    ```
    
    > Note: The local blockchain instance is running in-memory. The next time you start the instance, all contracts and state will disappear from memory and you will need to deploy them again.

&nbsp;

## Commands

```
Usage:

gamma-cli COMMAND [OPTIONS]

Commands:

  start-local      start a local Orbs personal blockchain instance listening on port
                   options: -port <PORT>
                   example: gamma-cli start-local -port 8080

  stop-local       stop a locally running Orbs personal blockchain instance

  gen-test-keys    generate a new batch of 10 test keys and store in orbs-test-keys.json (default filename)
                   options: -keys [OUTPUT_FILE]
                   example: gamma-cli gen-test-keys -keys orbs-test-keys.json

  deploy           deploy a smart contract with the code specified in the source file <CODE_FILE>
                   options: <CODE_FILE> -name [CONTRACT_NAME] -signer [ID_FROM_KEYS_JSON]
                   example: gamma-cli deploy MyToken.go -signer user1
                            gamma-cli deploy contract.go -name MyToken

  send-tx          sign and send the transaction specified in the JSON file <INPUT_FILE>
                   options: <INPUT_FILE> -arg# [OVERRIDE_ARG_#] -signer [ID_FROM_KEYS_JSON]
                   example: gamma-cli send-tx transfer.json -signer user1
                            gamma-cli send-tx transfer.json -arg2 b3d1caa2b3680e2c8feffa269c207c553fbbc828

  run-query        read state or run a read-only contract method as specified in the JSON file <INPUT_FILE>
                   options: <INPUT_FILE> -arg# [OVERRIDE_ARG_#] -signer [ID_FROM_KEYS_JSON]
                   example: gamma-cli run-query get-balance.json -signer user1
                            gamma-cli run-query get-balance.json -arg1 b3d1caa2b3680e2c8feffa269c207c553fbbc828

  get-status       get the current status of a sent transaction with txid <TX_ID> (from send-tx response)
                   options: <TX_ID>
                   example: gamma-cli get-status nXAmGL2peGvXkrDxC2cFaZwhykfMGFGj1DUJ9eDFRdSnNgCpQ69MQz

  tx-proof         get cryptographic proof for transaction receipt with txid <TX_ID> (from send-tx response)
                   options: <TX_ID>
                   example: gamma-cli tx-proof nXAmGL2peGvXkrDxC2cFaZwhykfMGFGj1DUJ9eDFRdSnNgCpQ69MQz

  upgrade-server   upgrade to the latest version of Gamma server

  version          print gamma-cli and Gamma server versions

  help             print this help screen


Options:

  -config string
    	path to config file (default "orbs-gamma-config.json")
  -env string
    	environment from config file containing server connection details (default "local")
  -keys string
    	name of the json file containing test keys (default "orbs-test-keys.json")
  -name string
    	name of the smart contract being deployed
  -port int
    	listening port for Gamma server (default "8080")
  -signer string
    	id of the signing key from the test key json (default "user1")
  -wait
    	wait until Gamma server is ready and listening

Multiple environments (eg. local and testnet) can be defined in orbs-gamma-config.json configuration file.
See https://github.com/orbs-network/orbs-contract-sdk for more info.
```

&nbsp;

## Upgrading to latest versions

* Upgrade to the latest version of `gamma-cli` by running in terminal:

    ```
    brew upgrade gamma-cli
    ```

* Upgrade to the latest version of Gamma server by running in terminal:

    ```
    gamma-cli upgrade-server
    ```

&nbsp;

## Advanced use

### Starting and stopping the server

The server runs locally and listens on port 8080 by default, although a custom port can be given as argument. The local server instance is a full blockchain network made from several nodes which communicate using the actual consensus protocol. 

When transactions are not sent, the nodes will keep closing empty blocks. It is therefore recommended to stop the server when it's not needed with the `gamma-cli stop-local` command.

### Test keys and accounts

When first launching Gamma server a batch of 10 testing accounts (public keys, private keys and addresses) will automatically be created and saved under `orbs-test-keys.json` in the local directory. Please note that these keys are for testing only and should not be used in secure production environments.

To replace the keys with a new batch of 10 accounts run `gamma-cli gen-test-keys` in terminal.

Every account in the file is given an ID. The default IDs are `user1` to `user10`. You can change the IDs by editing the file directly.

### Deploying smart contracts

To deploy a smart contract run the command `gamma-cli deploy` and provide the contract name. You will also need to provide the source code for the contract.

Contracts are immutable. If you want to update the code for a contract, deploy it again under a different name.

Note that Gamma server is an in-memory blockchain. When you stop the instance with `gamma-cli stop-local` all contracts will disappear.

### Sending transactions and calling contracts

Send transactions to contracts (write operations) by running `gamma-cli send-tx` and call contracts (read operations) by running `gamma-cli read`.

Both commands rely on an input JSON file which contains the actual details of the smart contract call:

```json
{
  "ContractName": "CounterExample",
  "MethodName": "add",
  "Arguments": [
    {
      "Type": "uint64",
      "Value": "25"
    }
  ]
}
``` 

The **contract name** should be the name chosen during deployment of the contract with `gamma-cli deploy`. The method name should be one of the exported methods found in the contract source code.

The array of arguments should be given according to the declaration of the method in contract source code. The following primitive argument types are supported:

* `uint32` - Number with 32 bit precision. Matches the `uint32` type in go contracts and `number` type in js contracts. Provided in the JSON as a string of a decimal number.

* `uint64` - Number with 64 bit precision. Matches the `uint64` type in go contracts and `BigInt` type in js contracts. Provided in the JSON as a string of a decimal number.

* `string` - String. Matches the `string` type in go contracts and `string` type in js contracts. Provided in the JSON as a UTF8 string.

* `bytes` - An array of bytes (blob). Matches the `[]byte` type in go contracts and `Uint8Array` type in js contracts. Provided in the JSON as a hex string.

In addition to these primitives, `gamma-cli` supports several aliases for convenience:

* `gamma:keys-file-address` - An account address by ID. The value should be an account ID from `orbs-test-keys.json`. An alias for the type `bytes` which means it matches the `[]byte` type in go contracts and `Uint8Array` type in js contracts. Used to pass the account address in raw form.

### Working with multiple environments

The command line tool supports multiple environments such as local and test net. To configure multiple environments create a file named `orbs-gamma-config.json` in the local directory, with the following format:

```json
{
  "Environments": {
    "local": {
      "VirtualChain": 42,
      "Endpoints": ["localhost"]
    },
    "testnet1": {
      "VirtualChain": 90043,
      "Endpoints": ["http://192.168.1.1", "http://192.168.2.2:8080"]
    },
    "testnet2": {
      "VirtualChain": 3007,
      "Endpoints": ["http://10.1.1.122", "https://node-example.com", "http://another.io:8081"]
    }
  }
}
```

You can choose the active environment by passing the `-env` command line argument to every command. For example:

```
gamma-cli send-tx transfer.json -signer user1 -env testnet2
```

If a config file does not exist, the default environment is `local` with virtual chain `42` and the endpoint `localhost`.

&nbsp;

## Building the tools from source

All developer tools (Gamma server and `gamma-cli`) are provided as binaries for convenience only.

If you're interested in building the tools by yourself from source, find Gamma server as part of the node core in https://github.com/orbs-network/orbs-network-go and `gamma-cli` as part of the client SDK in https://github.com/orbs-network/orbs-client-sdk-go.

&nbsp;

## Detailed documentation

The detailed documentation website for Gamma is available here:

https://orbs.gitbook.io

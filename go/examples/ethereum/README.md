# Ethereum Cross Chain Example

This simple example is designed to show you the basics of cross-chain access to Ethereum.

The smart contract supports two main methods:

* `callEthereumHelloWorld` - takes an Ethereum contract address and calls a method.

* `getEthereumHelloSaidLog` - takes an Ethereum contract address and TxHash and gets logs from the receipt.

## Testing this example on Gamma

Gamma is a local Orbs blockchain instance for smart contract developers. Use Gamma to test smart contracts locally before deploying them to test net.

1. Make sure `gamma-cli` and Gamma are [installed](../../../GAMMA.md).

2. Make sure a Go dev environment is [installed](../../../README.md).

3. Make sure Ethereum truffle is [installed](https://truffleframework.com/docs/truffle/getting-started/installation). 

4. Make sure Ethereum ganache is [installed](https://truffleframework.com/docs/ganache/quickstart) and running on port **7545**. 

5. Run the following in terminal:

    ```
    cd ~/go/src/github.com/orbs-network/orbs-contract-sdk/go/examples/ethereum/test
    go test
    ```
    
6. If you wish to run the `gamma-cli` commands manually instead (and not through the test), run the following in terminal:

    Call method scenario:

    ```
    cd ~/go/src/github.com/orbs-network/orbs-contract-sdk/go/examples/ethereum/test
    cd EthereumContract
    # ganache should be running on port 7545
    truffle exec deploy-new.js
    # copy ContractAddress from response, eg: 0x66c8bC6e162e45Da2Fc3337cF2164CA5E43CA4c5
    cd ..
    gamma-cli start-local
    gamma-cli deploy ../contract.go -name MyCrossChain
    gamma-cli run-query call-hello-world.json -arg1 0x66c8bC6e162e45Da2Fc3337cF2164CA5E43CA4c5
    gamma-cli stop-local
    ```
    
    Get logs scenario:
    
    ```
    cd ~/go/src/github.com/orbs-network/orbs-contract-sdk/go/examples/ethereum/test
    cd EthereumContract
    # ganache should be running on port 7545
    truffle exec emit-event.js
    # copy ContractAddress from response, eg: 0x66c8bC6e162e45Da2Fc3337cF2164CA5E43CA4c5
    # copy TxHash from the response, eg: 0x734c2bb544c90d7b178cdfa18e60d8d002c4e9158e716000f67dbaed72d1a093
    cd ..
    gamma-cli start-local
    gamma-cli deploy ../contract.go -name MyCrossChain
    gamma-cli run-query get-hello-said-log.json -arg1 0x66c8bC6e162e45Da2Fc3337cF2164CA5E43CA4c5 -arg2 0x734c2bb544c90d7b178cdfa18e60d8d002c4e9158e716000f67dbaed72d1a093
    gamma-cli stop-local
    ```
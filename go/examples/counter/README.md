# Counter Example

This simple example is designed to show you how to read and write state variables.

The counter smart contract stores a counter variable on state and supports two main methods:

* `get` - returns the current value of the counter.

* `add` - takes a value as argument and adds it to the counter.

## Testing this example on Gamma

Gamma is a local Orbs blockchain instance for smart contract developers. Use Gamma to test smart contracts locally before deploying them to test net.

1. Make sure `gamma-cli` and Gamma are [installed](../../../GAMMA.md).

2. Make sure a Go dev environment is [installed](../../../README.md).

3. Run the following in terminal:

    ```
    cd ~/go/src/github.com/orbs-network/orbs-contract-sdk/go/examples/counter/test
    go test
    ```
    
4. If you wish to run the `gamma-cli` commands manually instead (and not through the test), run the following in terminal:

    ```
    cd ~/go/src/github.com/orbs-network/orbs-contract-sdk/go/examples/counter/test
    gamma-cli start-local
    gamma-cli deploy ../contract.go -name MyCounter
    gamma-cli run-query get.json
    gamma-cli send-tx add-25.json
    gamma-cli send-tx add-25.json
    gamma-cli send-tx add-25.json
    gamma-cli run-query get.json
    gamma-cli stop-local
    ```
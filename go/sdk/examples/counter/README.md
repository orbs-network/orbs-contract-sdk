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
    gamma-cli start-local -wait
    go test
    ```

## End to end testing

All examples have end to end tests written in Go using [Orbs client SDK for GO](https://github.com/orbs-network/orbs-client-sdk-go/).
`tests/harness.go` introduces a very simple wrapper for the contract and is used in the end to end test.
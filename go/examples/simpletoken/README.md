# Simple Token Example

This simple example is designed to show you the basics of creating an address-based ledger.

The smart contract supports two main methods:

* `getBalance` - takes an address as argument and returns its balance.

* `transfer` - takes an amount and recipient address as arguments and attempts to transfer from the signer's account.

An initial supply of 1000 tokens is created during deploy and given to the address of the deploy signer.

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
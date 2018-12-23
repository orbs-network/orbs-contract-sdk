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
    cd ~/go/src/github.com/orbs-network/orbs-contract-sdk/go/examples/simpletoken/test
    go test
    ```
    
4. If you wish to run the `gamma-cli` commands manually instead (and not through the test), run the following in terminal:

    A working transfer scenario:

    ```
    cd ~/go/src/github.com/orbs-network/orbs-contract-sdk/go/examples/simpletoken/test
    gamma-cli start-local
    gamma-cli deploy ../contract.go -name MySimpleToken -signer user1
    gamma-cli run-query get-user1-balance.json
    gamma-cli run-query get-user2-balance.json
    gamma-cli send-tx transfer-15-to-user2.json -signer user1
    gamma-cli run-query get-user1-balance.json
    gamma-cli run-query get-user2-balance.json
    gamma-cli stop-local
    ```
    
    A failure when trying to send with insufficient funds:
    
    ```
    cd ~/go/src/github.com/orbs-network/orbs-contract-sdk/go/examples/simpletoken/test
    gamma-cli start-local
    gamma-cli deploy ../contract.go -name MySimpleToken -signer user1
    gamma-cli run-query get-user1-balance.json
    gamma-cli run-query get-user2-balance.json
    gamma-cli send-tx transfer-1500-to-user2.json -signer user1
    gamma-cli run-query get-user1-balance.json
    gamma-cli run-query get-user2-balance.json
    gamma-cli stop-local
    ```
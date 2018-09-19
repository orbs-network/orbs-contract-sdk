# Gamma - local ORBS blockchain for developers
>Version  0.5 (alpha)


## Overview
Gamma is a local ORBS blockchain to empower developers to easily and efficiently deploy, run & test smart contracts.<enter>
Gamma runs an in-memory virtual chain on top of an ORBS blockchain with N nodes on your local machine. 
Gamma-cli -  the command line interface is deigned to help you to interact with the virtual chain. 


## Getting Started... 

### Requirements
- Go 1.10.X installed 
- Mac or Linux (Windows support coming soon)

### Installation  
The installation of a personal ORBS and the command line interface is done by running the following command: 
- TODO: Add the command 

* To validate that the installation of Gamma- [TODO]
```go
DROR\ ODED\ ITAMAR PLEASE ADD A CODE EXMAPLE HERE ONCE RUNNING IT
```

### Let's start with examples 

2 contracts examples are provided to quickly get started:
> * [Counter contract](https://github.com/orbs-network/orbs-contract-sdk/tree/master/go/examples/counter "Counter Contract") - designed to show you how to read and write state variables.<ENTER>
> * [Fun token contract] (https://github.com/orbs-network/orbs-contract-sdk/tree/master/go/examples/tokens/fun_token "fun token contract")- designed to show you a basic token functionality.

**Steps to deploy the example contracts**  

* **Step 1**: Open the terminal & restart and start local ORBS blockchain instance <enter>
  `$ gamma-cli -start`. You should get a message "Your personal ORBS blockchain is ready for use"
* **Step 2** : Deploy your contract `$ gamma-cli deploy [contract file pathn] ` , you should get a message "Contract [file name] was deployed successfully".<enter>
          Please note that the code was compiled - part of the deployment process to save time.

> Code to start the local virtual chain and deploy the 2 contracts: 
``` 
gamma-cli start
gamma-cli deploy /examples/tokens/fun_token/fun_token.go
gamma-cli deploy /examples/counter/counter/fun_token.go
``` 

**Steps to test using Call or Send:**
* Use `call` when you want to access a smart contract method that reads from your state variables. In this case, the read will be done with no need to run the consensus. 
* Use `send` when you want to acess a smart contract method that also writes to the blockchain. This means that a condenses should be reached.

>Please note that the smart contract's arguments expected in `send` or `run` should be written in a jason format. 
> Example of jason  files can be downloaded here: [Counter jason ](https://github.com/orbs-network/orbs-contract-sdk/tree/master/go/examples/counter/tests) and [fun token jason]( https://github.com/orbs-network/orbs-contract-sdk/tree/master/go/examples/tokens/fun_token/tests )
```go
DROR\ ODED PLEASE ADD A CODE EXMAPLE HERE ONCE RUNNING IT
```
**Steps to test using the test files:**

We provided you corresponding test files for the examples.
> Test files can be downloaded: [Counter tests ](https://github.com/orbs-network/orbs-contract-sdk/tree/master/go/examples/counter/tests) and [fun token tests]( https://github.com/orbs-network/orbs-contract-sdk/tree/master/go/examples/tokens/fun_token/tests )
* **Step 1**: In the terminal `$ go [test file path]`. 
* **Step 2**: you should in the terminal the expected test results and actual results, including an indication of "yes" or "no" if the test passed. 

``` 
gamma-cli start
gamma-cli deploy /examples/tokens/fun_token/fun_token.go
gamma-cli deploy /examples/counter/counter/fun_token.go
``` 
-- TODO: ADD screenshots            
<p align="center">
  <img src="tbd?raw=true")
</p>

*

# Deploy & test your own contract  

>Make sure the ORBS blockchain is on. If not please use the `start` command.<enter>
 Please use a GO file, the `deploy`- also includes compilation with GO v1.10.x.
 
* Step 1: deploy your contract, using `deploy`
* Step 2: Test your contract using `run` or `$ GO test_file_name.go` 

# Gamma CLI

## Command line

`$ gamma-cli  <options>`

### Options 

* `start`  - start a local virtual chain over ORBS blockchain network, running on 3 nodes. 
* `Stop`   - stops the virtual chain. Unlike in a Mainnet or Testnet, the smart contract state variables are deleted. 
* `deploy` - compile the smart contract with go v10.0 and deploy it on the personal orbs blockchain on your machine. 
* `run`    - gets as arguments `call` or `send`. Use 
* `genKeys`- generates a new pair public and private key to sign on the transactions you send or you contract sends. 
             The keys are stored on your computer on a file named ORBS.KEYS.
* `help`   - information of all the commands that gamma-cli supports. 

>To ease the work with Gamma, part of the installation a pair of sK & pK are generated to sign the transactions.

---

## Project status

#### Gamma v0.5 (alpha) feature list

- Connecting to an in-memory ORBS blockchain with 3 nodes.
- Examples of basic contracts to run & test: token contract & counter contract. 
- Test jason with an example on how to test your owen smart contract easily (just copy and adjust). 
- Log files to assist with debugging


### Gamma v1 - coming next...
- Virtual chain configuration settings
- Support in additional Consensus algorithm.
- Block explorer APIs

## Licence  
MIT

#### TODO:
- How to use the logs , 
- work with the public sdk

---

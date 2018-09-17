#Kitsat - personal ORBS blockchain for developers

- Version  0.5 (alpha)

<p align="center">
  <img src="tbd?raw=true")
</p>

##Overview
Kitsat is a personal ORBS blockchain to empower developers to easily and efficiently deploy, run & test smart contracts.<enter>
Kitsat runs an in-memory ORBS blockchain with N nodes on your local machine. 
The Command line interface is deisnged to help you an interact with the blockchain network. 


##Getting Started... 

###Requirements
- Go 1.10.X installed 
- Mac or Linux (Windows support coming soon)

### Installation  
The installation of a personal ORBS and the command line interface is done by running the following command: 
- TODO: Add the command 

* To validate that the installation of Kitsat- [TODO]

###Let's start with examples 

2 contracts examples are provided to quickly get started:
>   *[Counter contract](https://github.com/orbs-network/orbs-contract-sdk/tree/master/go/examples/counter "Counter Contract") - 
designed to show you how to read and write state variables.
<ENTER>
*[MyToken contract](https://github.com/orbs-network/orbs-contract-sdk/tree/master/go/examples/MyToken "Counter contract".

**Steps to deploy the example contracts**  
* **Step 1**: Open the terminal & restart and start local ORBS blockchain instance <enter>
  `$ kitsat-cli -start`. You should get a message "Your personal ORBS blockchain is ready for use"
* **Step 2** : Deploy your contract `$ kitsat-cli deploy [contract file pathn] ` , you should get a message "Contract [file name] was deployed successfully".<enter>
          Please note that the code was compiled - part of the deployment process to save time.

**Steps to test the counter & FunToken contracts:**
* **Step 1**: In kitsat-cli `$ go [test file path]`. 
* **Step 2**: you should in the terminal the expected test results and actual results, including an indication of "yes" or "no" if the test passed. 

-- TODO: ADD screenshots            
<p align="center">
  <img src="tbd?raw=true")
</p>

*

#Deploy & test your own contract  

-TODO the test file should be 

#Kitsat CLI

##Commands

start , stop , deploy, run [call, send transcation]

Kitsat run call [jason contract name + arrgum,ents ]


##Project status

####Kitsat v0.5 (alpha) feature list
- Connecting to an in-memory ORBS blockchain with 3 nodes.
- Support Benchmark consensus algorithm (Lean Helix is planned to be realised soon)
- Examples of basic contracts to run & test: token contract & counter contract. 
- Test jason with an example on how to test your owen smart contract easily (just copy and adjust). 
- Log files to assist with debugging



### Kitsat v1 -coming next...
- Support in additional Consensus algorithm: Lean Helix  
- Virtual chain configuration settings
- block explorer APIs


## Licence  
TODO

#### TODO:
- How to use the logs , 

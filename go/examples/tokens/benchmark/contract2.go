package main

import (
	"fmt"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk"
)


var CONTRACT = sdk.ContractInfo{
	Name: "BenchmarkERC20Token",
	Permission: sdk.PERMISSION_SCOPE_SERVICE,
	Methods: map[string]sdk.MethodInfo{
		METHOD_INIT.Name :METHOD_INIT,
		METHOD_getName.Name: METHOD_getName,
		METHOD_getSymbol.Name: METHOD_getSymbol,
		METHOD_getDecimals.Name: METHOD_getDecimals,
		METHOD_totalSupply.Name: METHOD_totalSupply,
		METHOD_balanceOf.Name: METHOD_balanceOf,
		METHOD_allowance.Name: METHOD_allowance,
		METHOD_transfer.Name: METHOD_transfer,
		METHOD_approve.Name: METHOD_approve,
		METHOD_transferFrom.Name: METHOD_transferFrom,
	},
	InitSingleton: newContract,
}

func newContract(base *sdk.BaseContract) sdk.ContractInstance{
	return &contract{base}
}

type contract struct{ *sdk.BaseContract }

// name of the token
const TokenName = "ORBS"
const TokenSymbol = "ORB"
const TokenDecimals = 18


// init the contract
// set amount of coins during deployment
var METHOD_INIT = sdk.MethodInfo{
	Name: "_init",
	External: false,
	Access: sdk.ACCESS_SCOPE_READ_WRITE,
	Implementation: (*contract)._init,
}

func (c *contract) _init(ctx sdk.Context) error {
	err := c.State.WriteUint64ByKey(ctx, "totalSupply", uint64(0x100))
	if err != nil{
		return err
	}
	err = c.State.WriteUint64ByKey(ctx, "owner", uint64(0x100))
	if err != nil{
		return err
	}
	err = c.State.WriteStringByKey(ctx, "TokenName", TokenName)
	if err != nil{
		return err
	}
	err = c.State.WriteStringByKey(ctx, "TokenSymbol", TokenSymbol)
	if err != nil{
		return err
	}
	err = c.State.WriteUint64ByKey(ctx, "TokenDecimals", TokenDecimals)
	return err
}


// get the tokens name
var METHOD_getName = sdk.MethodInfo{
	Name: "getName",
	External: true,
	Access: sdk.ACCESS_SCOPE_READ_ONLY,
	Implementation: (*contract).getName,
}

// read from the contracts state
func (c *contract) getName(ctx sdk.Context) (string, error){
	return c.State.ReadStringByKey(ctx, "TokenName")
}

// get the token's symbol
var METHOD_getSymbol = sdk.MethodInfo{
	Name: "getSymbol",
	External: true,
	Access: sdk.ACCESS_SCOPE_READ_ONLY,
	Implementation: (*contract).getSymbol,
}

// read from the contract's state
func (c *contract) getSymbol(ctx sdk.Context) (string, error){
	return c.State.ReadStringByKey(ctx, "TokenSymbol")
}

// get the token's decimals
var METHOD_getDecimals = sdk.MethodInfo{
	Name: "getDecimals",
	External: true,
	Access: sdk.ACCESS_SCOPE_READ_ONLY,
	Implementation: (*contract).getDecimals,
}

// read from the contract's state
func (c *contract) getDecimals(ctx sdk.Context) (uint64, error){
	return c.State.ReadUint64ByKey(ctx, "TokenDecimals")
}


// returns the total token supply that was set in the deployment (_init) of the contract
var METHOD_totalSupply = sdk.MethodInfo{
	Name: "totalSupply",
	External: true,
	Access: sdk.ACCESS_SCOPE_READ_ONLY,
	Implementation: (*contract).totalSupply,
}

// read from the contract's state the total supply of the tokens
func (c *contract) totalSupply(ctx sdk.Context) (uint64 ,error) {
	return c.State.ReadUint64ByKey(ctx, "totalSupply")
}

// returns the total balance of some address
var METHOD_balanceOf = sdk.MethodInfo{
	Name: "balanceOf",
	External: true,
	Access: sdk.ACCESS_SCOPE_READ_ONLY,
	Implementation: (*contract).balanceOf,
}

// read from the contracts state the balance held by the given address
func (c *contract) balanceOf(ctx sdk.Context, address string) (uint64, error) {
	output, err := c.State.ReadUint64ByKey(ctx, string(address))

	if err != nil{
		return 0, err
	}
	return output, nil
}


// allowance returns how much can an address spender spend from the wallet of another address
var METHOD_allowance = sdk.MethodInfo{
	Name: "allowance",
	External: true,
	Access: sdk.ACCESS_SCOPE_READ_ONLY,
	Implementation: (*contract).allowance,
}

// read from the key "allowance-owner-spender"
func (c *contract) allowance(ctx sdk.Context, owner string, spender string) (uint64, error){
	output, err := c.State.ReadUint64ByKey(ctx, "allowance-"+string(owner)+"-"+string(spender))
	//output, err = c.State.ReadUint64ByKey(ctx, "allowance-"+string(owner)+"-"+string(spender))
	if err != nil{
		return 0, err
	}
	return output, nil
}

// transfers tokens from the owner's address to some other occount
var METHOD_transfer = sdk.MethodInfo{
	Name: "transfer",
	External: true,
	Access: sdk.ACCESS_SCOPE_READ_WRITE,
	Implementation: (*contract).transfer,
}

// TODO must be called by the owner of the tokens, have a message owner parameter
func (c *contract) transfer(ctx sdk.Context, owner string, recepient string, amount uint64) (uint32, error){
	//make sure sender has sufficient funds
	balance, err := c.State.ReadUint64ByKey(ctx, string(owner))
	if err != nil{
		return 1, err
	}
	if balance < amount{
		return 1, fmt.Errorf("not enough balance to complete transaction. balance: %d, amount: %d", balance, amount)
	}

	// get recepient's balance
	rBalance, err := c.State.ReadUint64ByKey(ctx, string(recepient))
	if err != nil{
		return 1, err
	}

	// decrease sender's balance
	err = c.State.WriteUint64ByKey(ctx, string(owner), balance - amount)
	if err != nil{
		return 1, fmt.Errorf("could not complete one of the writing operations, %v", err)
	}
	fmt.Printf("string owner: %s, balance - amount = %d, err: %v\n", string(owner), balance - amount, err)

	// increase recepient's balance
	err = c.State.WriteUint64ByKey(ctx, string(recepient), rBalance + amount)
	if err != nil{
		// TODO decreased the senders balance but not increase recepient's balance
		return 1, fmt.Errorf("could not complete one of the writing operations, %v", err)
	}
	return 0, nil
}

// approve an account to send from the owner's tokens to the recepient tokens
var METHOD_approve = sdk.MethodInfo{
	Name: "approve",
	External: true,
	Access: sdk.ACCESS_SCOPE_READ_WRITE,
	Implementation: (*contract).approve,
}

// make sure the owner has sufficeint funds then write to state the amount approved
func (c *contract) approve(ctx sdk.Context, owner string, sender string, amount uint64) (uint32, error){
	// make sure owner has sufficient funds
	balance, err := c.State.ReadUint64ByKey(ctx, string(owner))
	if err != nil{
		return 1, err
	}

	// return an error and dont change any state
	if balance < amount{
		return 1, fmt.Errorf("the owner does not have suffiecient funds, want %d, have %d", amount, balance)
	}

	// read the current approved funds and write the addition of the amount
	currentApproved, err := c.State.ReadUint64ByKey(ctx, "allowance-"+string(owner)+"-"+string(sender))
	if err != nil{
		return 0, err
	}

	err = c.State.WriteUint64ByKey(ctx, "allowance-"+string(owner)+"-"+string(sender), currentApproved + amount)
	if err != nil{
		return 1, err
	}

	return 0, nil
}

// transfer the approved amount from the sender to the recepient
var METHOD_transferFrom = sdk.MethodInfo{
	Name: "transferFrom",
	External: true,
	Access: sdk.ACCESS_SCOPE_READ_WRITE,
	Implementation: (*contract).transferFrom,
}

// read from the allowance of the msg.sender, lower it by amount and change the recepient balance by that amount
func (c *contract) transferFrom(ctx sdk.Context, owner string, sender string, amount uint64, recepient string)(uint32, error){
	// read the amount of allowance the msg.sender has
	allowance, err := c.State.ReadUint64ByKey(ctx, "allowance-"+string(owner)+"-"+string(sender))
	if err != nil{
		return 1, nil
	}

	// if the msg.sender does not have sufficient funds cancel the transaction
	if amount > allowance {
		return 1, fmt.Errorf("the sender does not have suffiecent allowance, allowance: %d, amount: %d", allowance, amount)
	}

	// read the recepients current balance from the contract state
	balance, err := c.State.ReadUint64ByKey(ctx, string(recepient))
	if err != nil {
		return 1, err
	}

	// update the recepeints balance
	err = c.State.WriteUint64ByKey(ctx, string(recepient), balance + amount)
	if err != nil {
		return 1, err
	}

	// update the msg.sender allowance
	err = c.State.WriteUint64ByKey(ctx, "allowance-"+string(owner)+"-"+string(sender), allowance - amount)
	if err != nil {
		// TODO so what should you do?
		// the recepient's amount was increased but the allowance was not decreased
		return 0, err
	}

	// transaction was done succesfully
	return 0, nil
}
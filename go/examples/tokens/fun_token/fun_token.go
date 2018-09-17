package main

import "github.com/orbs-network/orbs-contract-sdk/go/sdk"

var CONTRACT = sdk.ContractInfo{
	Name:       "FunToken",
	Permission: sdk.PERMISSION_SCOPE_SERVICE,
	Methods: map[string]sdk.MethodInfo{
		METHOD_INIT.Name:  			METHOD_INIT,
		METHOD_MINT.Name:			METHOD_MINT,
		METHOD_BALANCEOF.Name:  	METHOD_BALANCEOF,
		METHOD_TRANSFER.Name:   	METHOD_TRANSFER,
		METHOD_TOTALSUPPLY.Name:	METHOD_TOTALSUPPLY,
	},
	InitSingleton: newContract,
}

func newContract(base *sdk.BaseContract) sdk.ContractInstance {
	return &contract{base}
}

type contract struct{ *sdk.BaseContract }

///////////////////////////////////////////////////////////////////////////

var METHOD_INIT = sdk.MethodInfo{
	Name:           "_init",
	External:       false,
	Access:         sdk.ACCESS_SCOPE_READ_WRITE,
	Implementation: (*contract)._init,
}

func (c *contract) _init(ctx sdk.Context) error {
	c.State.WriteStringByKey(ctx,"name","Fun_Token")
	c.State.WriteStringByKey(ctx,"symbol","SMP")
	return c.State.WriteUint64ByKey(ctx, "totalSupply", 1000000000)
}

///////////////////////////////////////////////////////////////////////////

var METHOD_MINT = sdk.MethodInfo{
	Name:			"mint",
	External:		false,
	Access:			sdk.ACCESS_SCOPE_READ_WRITE,
	Implementation:	(*contract).mint,
}

func (c *contract) mint(ctx sdk.Context, address Ripmd160Sha256, value uint64) error {
	// TODO: add limitation according to current total supply
	return c.State.WriteUint64ByAddress(ctx,address,value)
}

///////////////////////////////////////////////////////////////////////////

var METHOD_BALANCEOF = sdk.MethodInfo{
	Name:           "balanceOf",
	External:       false,
	Access:         sdk.ACCESS_SCOPE_READ_WRITE,
	Implementation: (*contract).balanceOf,
}

func (c *contract) balanceOf(ctx sdk.Context, address Ripmd160Sha256) error {
	return c.State.ReadUint64ByAddress(ctx, address)
}

///////////////////////////////////////////////////////////////////////////

var METHOD_TRANSFER = sdk.MethodInfo{
	Name:           "transfer",
	External:       false,
	Access:         sdk.ACCESS_SCOPE_READ_WRITE,
	Implementation: (*contract).transfer,
}

func (c *contract) transfer(ctx sdk.Context, from Ripmd160Sha256, to Ripmd160Sha256, value uint32) error {
	fromAmount := c.State.ReadUint64ByAddress(ctx,from)
	toAmount := c.State.RedaUint64ByAddress(ctx,to)

	if fromAmount < value{
		return nil
	}

	if toAmount + value < totalSupply {
		return nil
	}

	c.State.WriteUint64ByAddress(ctx,from, fromAmount - value)
	return c.State.WriteUint64ByAddress(to,toAmount + value)
}

///////////////////////////////////////////////////////////////////////////

var METHOD_TOTALSUPPLY = sdk.MethodInfo{
	Name:           "totalSupply",
	External:       false,
	Access:         sdk.ACCESS_SCOPE_READ_WRITE,
	Implementation: (*contract).totalSupply,
}


func (c *contract) totalSupply(ctx sdk.Context){
	return c.State.ReadUint64ByKey(ctx,"totalSupply")
}



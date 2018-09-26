package fun_token

import (
	"github.com/orbs-network/orbs-contract-sdk/go/sdk"
)


// Define contract info, it is similar to interface
var CONTRACT = sdk.ContractInfo{

	// Contract name TODO: we had a discussion about it (should it be the filename or what)
	Name:       "FunToken",

	// Permission scope can be either SERVICE or SYSTEM, for non-system contracts use PERMISSION_SCOPE_SERVICE
	Permission: sdk.PERMISSION_SCOPE_SERVICE,

	// Map of the contract methods, the visibility modes will be defined later
	Methods: map[string]sdk.MethodInfo{
		// Key format: 		METHOD_(FUNCTION_NAME).Name
		// Value format:	METHOD_(FUNCTION_NAME)
		METHOD_INIT.Name:        METHOD_INIT,
		METHOD_MINT.Name:        METHOD_MINT,
		METHOD_BALANCEOF.Name:   METHOD_BALANCEOF,
		METHOD_TRANSFER.Name:    METHOD_TRANSFER,
		METHOD_TOTALSUPPLY.Name: METHOD_TOTALSUPPLY,
	},
	// Define as singleton TODO: why?
	InitSingleton: newContract,
}

// TODO: ctor?
func newContract(base *sdk.BaseContract) sdk.ContractInstance {
	return &contract{base}
}

// TODO: why?
type contract struct{ *sdk.BaseContract }

///////////////////////////////////////////////////////////////////////////

// Define method info for _init function
// The _init function will be called automatically once
// (the deploy system call look for _init method) TODO: Is that right??
// when the contract will be deployed to the blockchain
var METHOD_INIT = sdk.MethodInfo{

	// Name of the method is _init
	Name:           "_init",

	// External property is false for this method, it can be called only from within the contract
	External:       false,

	// Access level is READ_WRITE because this method both reads and writes to the blockchain
	Access:         sdk.ACCESS_SCOPE_READ_WRITE,

	// A pointer to the function TODO: any better explanations?
	Implementation: (*contract)._init,
}

// Implementation of the _init function - stores token variables to the state of the contract
// Each function receives the sdk.Context as a parameter and returns an error
// Parameters:
// 	(1) sdk.Context
// Returns:
// 	(1) an error in case of an exception
func (c *contract) _init(ctx sdk.Context) error {

	// Save a key with the name "name" and string value "Fun_Token" in the state of the contract
	c.State.WriteStringByKey(ctx, "name", "Fun_Token")

	// Save a key with the name "symbol" and string value "FUN"
	c.State.WriteStringByKey(ctx, "symbol", "FUN")

	// Save a key with the name "totalSupply" and uint64 value "FUN" and return
	return c.State.WriteUint64ByKey(ctx, "totalSupply", 1000000000)
}

///////////////////////////////////////////////////////////////////////////

// Define method info for mint function
var METHOD_MINT = sdk.MethodInfo{

	// Name of the method is mint
	Name:           "mint",

	// External property is true for this method, it can be called from within the contract
	// and also from other contracts
	External:       true,

	// Access level is READ_WRITE because this method both reads and writes to the blockchain
	Access:         sdk.ACCESS_SCOPE_READ_WRITE,

	// A pointer to the function TODO: any better explanations?
	Implementation: (*contract).mint,
}

// Implementation of the mint function.
// This function mint a given value of tokens once into a given address.
// Parameters:
// 	(1) sdk.Context
// 	(2) address - a byte array which represents an address
// 	(3) value - an uint64 which represents the total supply of tokens
// Return:
// 	(1) an error in case of an exception
func (c *contract) mint(ctx sdk.Context, address []byte, value uint64) error {
	// TODO: add logic only the owner can call this method
	// TODO: add logic if all tokens were minted already

	// Mint the value number of tokens to the given address
	return c.State.WriteUint64ByAddress(ctx, address, value)
}

///////////////////////////////////////////////////////////////////////////

// Define method info for balanceOf function
var METHOD_BALANCEOF = sdk.MethodInfo{

	// Name of the method is balanceOf
	Name:           "balanceOf",

	// External property is true for this method, it can be called from within the contract
	// and also from other contracts
	External:       true,

	// Access level is READ_ONLY because this method only reads from the blockchain
	Access:         sdk.ACCESS_SCOPE_READ_ONLY,

	// A pointer to the function TODO: any better explanations?
	Implementation: (*contract).balanceOf,
}

// Implementation of the balanceOf function.
// This function returns the number of tokens of a given address.
// Parameters:
// 	(1) sdk.Context
// 	(2) address - a byte array which represents an address
// Return:
// 	(1) The number of tokens in a uint64 value
// 	(2) an error in case of an exception
func (c *contract) balanceOf(ctx sdk.Context, address []byte) (uint64,error) {
	return c.State.ReadUint64ByAddress(ctx, address)
}

///////////////////////////////////////////////////////////////////////////

// Define method info for transfer function
var METHOD_TRANSFER = sdk.MethodInfo{

	// Name of the method is transfer
	Name:           "transfer",

	// External property is true for this method, it can be called from within the contract
	// and also from other contracts
	External:       true,

	// Access level is READ_WRITE because this method both reads and writes to the blockchain
	Access:         sdk.ACCESS_SCOPE_READ_WRITE,

	// A pointer to the function TODO: any better explanations?
	Implementation: (*contract).transfer,
}

// Implementation of the balanceOf function.
// This function returns the number of tokens of a given address.
// Parameters:
// 	(1) sdk.Context
// 	(2) address - a byte array which represents an address
// Return:
// 	(1) The number of tokens in a uint64 value
// 	(2) an error in case of an exception
func (c *contract) transfer(ctx sdk.Context, from string, to string, value uint64) error {
	fromAmount, err := c.State.ReadUint64ByKey(ctx, from)

	if err != nil {
		return err
	}

	toAmount, err := c.State.ReadUint64ByKey(ctx, to)

	if err != nil {
		return err
	}

	//totalSupply, err := c.State.ReadUint64ByKey(ctx,"totalSupply")

	if err != nil {
		return err
	}

	//if fromAmount < value {
	//	return nil
	//}
	//
	//if toAmount+value < totalSupply {
	//	return nil
	//}

	c.State.WriteUint64ByKey(ctx, from, fromAmount-value)
	return c.State.WriteUint64ByKey(ctx, to, toAmount+value)
}

///////////////////////////////////////////////////////////////////////////

// Define method info for totalSupply function
var METHOD_TOTALSUPPLY = sdk.MethodInfo{

	// Name of the method is totalSupply
	Name:           "totalSupply",

	// External property is true for this method, it can be called from within the contract
	// and also from other contracts
	External:       true,

	// Access level is READ_ONLY because this method only reads from the blockchain
	Access:         sdk.ACCESS_SCOPE_READ_ONLY,

	// A pointer to the function TODO: any better explanations?
	Implementation: (*contract).totalSupply,
}

// Implementation of the totalSupply function.
// This function returns the total number of tokens //
// TODO: maybe I should be more accurate (total supply vs total circulation)
// Parameters:
// 	(1) sdk.Context
// Return:
// 	(1) The total supply of tokens in an uint64 value
// 	(2) an error in case of an exception
func (c *contract) totalSupply(ctx sdk.Context) (uint64, error){
	return c.State.ReadUint64ByKey(ctx, "totalSupply")
}

///////////////////////////////////////////////////////////////////////////

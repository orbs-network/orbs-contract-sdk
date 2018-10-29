package main

import (
	"errors"
	"fmt"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk"
)

var CONTRACT = sdk.ContractInfo{
	Name:       "Singular",
	Permission: sdk.PERMISSION_SCOPE_SERVICE,
	Methods: map[string]sdk.MethodInfo{
		METHOD_INIT.Name: METHOD_INIT,
		METHOD_SET.Name:  METHOD_SET,
		METHOD_CAS.Name:  METHOD_CAS,
		METHOD_GET.Name:  METHOD_GET,
	},
	InitSingleton: newContract,
}

const STATE_VALUE_KEY = "value"

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
	return c.State.WriteStringByKey(ctx, STATE_VALUE_KEY, "nil")
}

///////////////////////////////////////////////////////////////////////////

var METHOD_SET = sdk.MethodInfo{
	Name:           "set",
	External:       true,
	Access:         sdk.ACCESS_SCOPE_READ_WRITE,
	Implementation: (*contract).set,
}

func (c *contract) set(ctx sdk.Context, value string) error {
	return c.State.WriteStringByKey(ctx, STATE_VALUE_KEY, value)
}

///////////////////////////////////////////////////////////////////////////

var METHOD_CAS = sdk.MethodInfo{
	Name:           "cas",
	External:       true,
	Access:         sdk.ACCESS_SCOPE_READ_WRITE,
	Implementation: (*contract).cas,
}

func (c *contract) cas(ctx sdk.Context, oldValue string, newValue string) error {
	stateValue, err := c.State.ReadStringByKey(ctx, STATE_VALUE_KEY)

	if err != nil {
		return err
	}

	if stateValue == oldValue {
		return c.State.WriteStringByKey(ctx, STATE_VALUE_KEY, newValue)
	} else {
		return errors.New(fmt.Sprintf("CAS Failed: Current value (%s) did not match provided (%s)", stateValue, oldValue))
	}
}

///////////////////////////////////////////////////////////////////////////

var METHOD_GET = sdk.MethodInfo{
	Name:           "get",
	External:       true,
	Access:         sdk.ACCESS_SCOPE_READ_ONLY,
	Implementation: (*contract).get,
}

func (c *contract) get(ctx sdk.Context) (string, error) {
	return c.State.ReadStringByKey(ctx, STATE_VALUE_KEY)
}

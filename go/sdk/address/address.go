// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package address

import (
	"fmt"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/context"
)

func ValidateAddress(address []byte) {
	if len(address) != 20 {
		panic(fmt.Sprintf("valid address length is %d bytes, received %d bytes", 20, len(address)))
	}
}

func GetSignerAddress() []byte {
	contextId, handler, permissionScope := context.GetContext()
	return handler.SdkAddressGetSignerAddress(contextId, permissionScope)
}

func GetCallerAddress() []byte {
	contextId, handler, permissionScope := context.GetContext()
	return handler.SdkAddressGetCallerAddress(contextId, permissionScope)
}

func GetOwnAddress() []byte {
	contextId, handler, permissionScope := context.GetContext()
	return handler.SdkAddressGetOwnAddress(contextId, permissionScope)
}

func GetContractAddress(contractName string) []byte {
	contextId, handler, permissionScope := context.GetContext()
	return handler.SdkAddressGetContractAddress(contextId, permissionScope, contractName)
}

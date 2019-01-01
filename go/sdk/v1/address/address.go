package address

import (
	"fmt"
	"github.com/orbs-network/orbs-contract-sdk/go/context"
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

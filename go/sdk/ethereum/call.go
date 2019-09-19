// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package ethereum

import "github.com/orbs-network/orbs-contract-sdk/go/context"

// ethContractAddress: string like "0x66c8bC6e162e45Da2Fc3337cF2164CA5E43CA4c5"
func CallMethod(ethContractAddress string, jsonAbi string, methodName string, out interface{}, args ...interface{}) {
	contextId, handler, permissionScope := context.GetContext()
	handler.SdkEthereumCallMethod(contextId, permissionScope, ethContractAddress, jsonAbi, 0, methodName, out, args...)
}

// ethContractAddress: string like "0x66c8bC6e162e45Da2Fc3337cF2164CA5E43CA4c5"
func CallMethodAtBlock(ethBlockNumber uint64, ethContractAddress string, jsonAbi string, methodName string, out interface{}, args ...interface{}) {
	contextId, handler, permissionScope := context.GetContext()
	handler.SdkEthereumCallMethod(contextId, permissionScope, ethContractAddress, jsonAbi, ethBlockNumber, methodName, out, args...)
}

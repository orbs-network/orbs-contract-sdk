// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package ethereum

import "github.com/orbs-network/orbs-contract-sdk/go/sdk/context"

// returns the current Ethereum block number - with some safety margin to avoid forks (about 25 minutes)
func GetBlockNumber() uint64 {
	contextId, handler, permissionScope := context.GetContext()
	return handler.SdkEthereumGetBlockNumber(contextId, permissionScope)
}

// returns the Ethereum block number of a timestamp - make sure its older than the safety margin to avoid forks (about 25 minutes)
func GetBlockNumberByTime(ethBlockTimestamp uint64) uint64 {
	contextId, handler, permissionScope := context.GetContext()
	return handler.SdkEthereumGetBlockNumberByTime(contextId, permissionScope, ethBlockTimestamp)
}

// returns the current Ethereum block's timestamp - with some safety margin to avoid forks (about 25 minutes)
func GetBlockTime() uint64 {
	contextId, handler, permissionScope := context.GetContext()
	return handler.SdkEthereumGetBlockTime(contextId, permissionScope)
}

// returns the timestamp of the Ethereum block number - make sure its older than the safety margin to avoid forks (about 25 minutes)
func GetBlockTimeByNumber(ethBlockNumber uint64) uint64 {
	contextId, handler, permissionScope := context.GetContext()
	return handler.SdkEthereumGetBlockTimeByNumber(contextId, permissionScope, ethBlockNumber)
}

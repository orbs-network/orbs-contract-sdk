// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package env

import "github.com/orbs-network/orbs-contract-sdk/go/context"

func GetBlockHeight() uint64 {
	contextId, handler, permissionScope := context.GetContext()
	return handler.SdkEnvGetBlockHeight(contextId, permissionScope)
}

func GetBlockTimestamp() uint64 {
	contextId, handler, permissionScope := context.GetContext()
	return handler.SdkEnvGetBlockTimestamp(contextId, permissionScope)
}

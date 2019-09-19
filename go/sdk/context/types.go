// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package context

type ContextId []byte

type PermissionScope uint16

const (
	PERMISSION_SCOPE_SYSTEM  PermissionScope = 1
	PERMISSION_SCOPE_SERVICE PermissionScope = 2
)

type ContractInfo struct {
	PublicMethods []interface{}
	SystemMethods []interface{}
	EventsMethods []interface{}
	Permission    PermissionScope
}

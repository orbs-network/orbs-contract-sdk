// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package state

import (
	"encoding/binary"
	"github.com/orbs-network/orbs-contract-sdk/go/context"
)

func ReadBytes(key []byte) []byte {
	contextId, handler, permissionScope := context.GetContext()
	return handler.SdkStateReadBytes(contextId, permissionScope, key)
}

func ReadString(key []byte) string {
	bytes := ReadBytes(key)
	return string(bytes)
}

func ReadUint64(key []byte) uint64 {
	bytes := ReadBytes(key)
	if len(bytes) < 8 {
		return 0
	}
	return binary.LittleEndian.Uint64(bytes)
}

func ReadUint32(key []byte) uint32 {
	bytes := ReadBytes(key)
	if len(bytes) < 4 {
		return 0
	}
	return binary.LittleEndian.Uint32(bytes)
}

func ReadBytes20(key []byte) (out [20]byte) {
	bytes := ReadBytes(key)
	if len(bytes) != 20 {
		return
	}
	copy(out[:], bytes)
	return
}

func ReadBytes32(key []byte) (out [32]byte) {
	bytes := ReadBytes(key)
	if len(bytes) != 32 {
		return
	}
	copy(out[:], bytes)
	return
}

// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package state

import (
	"encoding/binary"
	"github.com/orbs-network/orbs-contract-sdk/go/context"
	"math/big"
)

func WriteBytes(key []byte, value []byte) {
	contextId, handler, permissionScope := context.GetContext()
	handler.SdkStateWriteBytes(contextId, permissionScope, key, value)
}

func WriteString(key []byte, value string) {
	bytes := []byte(value)
	WriteBytes(key, bytes)
}

func WriteUint64(key []byte, value uint64) {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, value)
	WriteBytes(key, bytes)
}

func WriteUint32(key []byte, value uint32) {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, value)
	WriteBytes(key, bytes)
}

func WriteBytes20(key []byte, value [20]byte) {
	WriteBytes(key, value[:])
}

func WriteBytes32(key []byte, value [32]byte) {
	WriteBytes(key, value[:])
}

func WriteBool(key []byte, value bool) {
	bytes := make([]byte, 1)
	if value {
		bytes[0] = 1
	}
	WriteBytes(key, bytes)
}

func WriteBigInt(key []byte, value *big.Int) {
	actual := [32]byte{}
	if value != nil {
		b := value.Bytes()
		copy(actual[32-len(b):], b)
	}
	WriteBytes(key, actual[:])
}

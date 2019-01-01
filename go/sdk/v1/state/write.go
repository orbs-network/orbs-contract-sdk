package state

import (
	"encoding/binary"
	"github.com/orbs-network/orbs-contract-sdk/go/context"
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

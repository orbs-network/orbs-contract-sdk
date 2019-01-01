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

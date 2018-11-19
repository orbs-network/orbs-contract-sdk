package state

import (
	"github.com/orbs-network/orbs-contract-sdk/go/context"
	"encoding/binary"
)

func ReadBytesByAddress(address []byte) []byte {
	contextId, handler, permissionScope := context.GetContext()
	return handler.SdkStateReadBytesByAddress(contextId, permissionScope, address)
}

func ReadBytesByKey(key string) []byte {
	address := keyToAddress(key)
	return ReadBytesByAddress(address)
}

func ReadStringByAddress(address []byte) string {
	bytes := ReadBytesByAddress(address)
	return string(bytes)
}

func ReadStringByKey(key string) string {
	address := keyToAddress(key)
	return ReadStringByAddress(address)
}

func ReadUint64ByAddress(address []byte) uint64 {
	bytes := ReadBytesByAddress(address)
	if len(bytes) < 8 {
		return 0
	}
	return binary.LittleEndian.Uint64(bytes)
}

func ReadUint64ByKey(key string) uint64 {
	address := keyToAddress(key)
	return ReadUint64ByAddress(address)
}

func ReadUint32ByAddress(address []byte) uint32 {
	bytes := ReadBytesByAddress(address)
	if len(bytes) < 4 {
		return 0
	}
	return binary.LittleEndian.Uint32(bytes)
}

func ReadUint32ByKey(key string) uint32 {
	address := keyToAddress(key)
	return ReadUint32ByAddress(address)
}

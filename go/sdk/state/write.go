package state

import (
	"github.com/orbs-network/orbs-contract-sdk/go/context"
	"encoding/binary"
)

func WriteBytesByAddress(address []byte, value []byte) {
	contextId, handler, permissionScope := context.GetContext()
	handler.SdkStateWriteBytesByAddress(contextId, permissionScope, address, value)
}

func WriteBytesByKey(key string, value []byte) {
	address := keyToAddress(key)
	WriteBytesByAddress(address, value)
}

func WriteStringByAddress(address []byte, value string) {
	bytes := []byte(value)
	WriteBytesByAddress(address, bytes)
}

func WriteStringByKey(key string, value string) {
	address := keyToAddress(key)
	WriteStringByAddress(address, value)
}

func WriteUint64ByAddress(address []byte, value uint64) {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, value)
	WriteBytesByAddress(address, bytes)
}

func WriteUint64ByKey(key string, value uint64) {
	address := keyToAddress(key)
	WriteUint64ByAddress(address, value)
}

func WriteUint32ByAddress(address []byte, value uint32) {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, value)
	WriteBytesByAddress(address, bytes)
}

func WriteUint32ByKey(key string, value uint32) {
	address := keyToAddress(key)
	WriteUint32ByAddress(address, value)
}

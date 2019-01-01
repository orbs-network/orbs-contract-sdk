package state

import (
	"crypto/sha256"
)

func calc20ByteHash(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[12:]
}

func keyToAddress(key string) []byte {
	return calc20ByteHash([]byte(key))
}

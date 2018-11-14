package state

import (
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
)

func calcRipmd160Sha256(data []byte) []byte {
	hash := sha256.Sum256(data)
	r := ripemd160.New()
	r.Write(hash[:])
	return r.Sum(nil)
}

func keyToAddress(key string) []byte {
	return calcRipmd160Sha256([]byte(key))
}

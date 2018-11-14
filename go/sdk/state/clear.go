package state

func ClearByAddress(address []byte) {
	WriteBytesByAddress(address, []byte{})
}

func ClearByKey(key string) {
	address := keyToAddress(key)
	ClearByAddress(address)
}
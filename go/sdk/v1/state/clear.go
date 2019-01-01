package state

func Clear(key []byte) {
	WriteBytes(key, []byte{})
}

package state

func Rename(oldKey []byte, newKey []byte) {
	value := ReadBytes(oldKey)
	Clear(oldKey)
	WriteBytes(newKey, value)
}
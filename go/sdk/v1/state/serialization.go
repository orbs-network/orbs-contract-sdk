package state

import (
	"fmt"
	"reflect"
)

// Takes struct as parameter:
// SerializeStruct("key", Value{})
func SerializeStruct(compositeKey string, item interface{}) error {
	meta := reflect.ValueOf(item)

	for i := 0; i < meta.NumField(); i++ {
		f := meta.Field(i)
		key := Key(compositeKey, "$", meta.Type().Field(i).Name)

		switch v := f.Kind(); v {
		case reflect.String:
			WriteString(key, f.String())
		case reflect.Uint64:
			WriteUint64(key, f.Interface().(uint64))
		case reflect.Uint32:
			WriteUint32(key, f.Interface().(uint32))
		case reflect.Slice:
			WriteBytes(key, f.Interface().([]byte))
		default:
			return fmt.Errorf("failed to serialize key %s with type %s", key, v)
		}
	}

	return nil
}

// Takes pointer as parameter:
// DeserializeStruct("key", &Value{})
func DeserializeStruct(compositeKey string, value interface{}) error {
	meta := reflect.ValueOf(value).Elem()
	for i := 0; i < meta.NumField(); i++ {
		f := meta.Field(i)
		key := Key(compositeKey, "$", meta.Type().Field(i).Name)

		fValue := meta.Field(i)

		switch v := f.Kind(); v {
		case reflect.String:
			fValue.Set(reflect.ValueOf(ReadString(key)))
		case reflect.Uint64:
			fValue.Set(reflect.ValueOf(ReadUint64(key)))
		case reflect.Uint32:
			fValue.Set(reflect.ValueOf(ReadUint32(key)))
		case reflect.Slice:
			bytes := ReadBytes(key)
			if len(bytes) > 0 { // to preserve require.EqualValues checks
				fValue.Set(reflect.ValueOf(bytes))
			}
		default:
			return fmt.Errorf("failed to deserialize key %s with type %s", key, v)
		}
	}

	return nil
}

// Takes struct as parameter:
// DeleteStruct("key", Value{})
func DeleteStruct(compositeKey string, value interface{}) {
	meta := reflect.ValueOf(value)
	for i := 0; i < meta.NumField(); i++ {
		key := Key(compositeKey, "$"+meta.Type().Field(i).Name)
		Clear(key)
	}
}

// Takes struct as parameter:
// RenameStruct("key", Value{})
func RenameStruct(oldCompositeKey, newCompositeKey string, value interface{}) {
	meta := reflect.ValueOf(value)
	for i := 0; i < meta.NumField(); i++ {
		oldKey := Key(oldCompositeKey, "$"+meta.Type().Field(i).Name)
		newKey := Key(newCompositeKey, "$"+meta.Type().Field(i).Name)
		Rename(oldKey, newKey)
	}
}

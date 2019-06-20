package state

import (
	"github.com/pkg/errors"
	"reflect"
)

type Serializer func(compositeKey string, item interface{}) error
type Deserializer func(compositeKey string, item interface{}) error
type Deleter func(compositeKey string, item interface{})

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
			return errors.Errorf("failed to serialize key %s with type %s", key, v)
		}
	}

	return nil
}

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
			fValue.Set(reflect.ValueOf(ReadBytes(key)))
		default:
			return errors.Errorf("failed to deserialize key %s with type %s", key, v)
		}
	}

	return nil
}

func DeleteStruct(compositeKey string, value interface{}) {
	meta := reflect.ValueOf(value)
	for i := 0; i < meta.NumField(); i++ {
		key := Key(compositeKey, "$"+meta.Type().Field(i).Name)
		Clear(key)
	}
}

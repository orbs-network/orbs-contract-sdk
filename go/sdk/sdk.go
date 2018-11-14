package sdk

type ProcessorType uint16

const (
	PROCESSOR_TYPE_NATIVE ProcessorType = 1
)

func Export(funcs ...interface{}) []interface{} {
	return funcs
}
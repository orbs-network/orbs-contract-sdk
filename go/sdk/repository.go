package sdk

type ContractInfo struct {
	Name          string
	Permission    ExecutionPermissionScope
	Methods       map[string]MethodInfo
	InitSingleton func(*BaseContract) ContractInstance
}

type MethodInfo struct {
	Name           string
	External       bool
	Access         ExecutionAccessScope
	Implementation interface{}
}

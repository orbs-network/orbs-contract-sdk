package context

type ContextId uint32 // TODO: change to []byte (in protos too)

type PermissionScope uint16

const (
	PERMISSION_SCOPE_SYSTEM  PermissionScope = 1
	PERMISSION_SCOPE_SERVICE PermissionScope = 2
)

type ContractInfo struct {
	PublicMethods []interface{}
	SystemMethods []interface{}
	EventsMethods []interface{}
	Permission    PermissionScope
}

package sdk

type ExecutionPermissionScope uint16

const (
	PERMISSION_SCOPE_SYSTEM ExecutionPermissionScope = 1
	PERMISSION_SCOPE_SERVICE ExecutionPermissionScope = 2
)

type ExecutionAccessScope uint16

const (
	ACCESS_SCOPE_READ_ONLY ExecutionAccessScope = 1
	ACCESS_SCOPE_READ_WRITE ExecutionAccessScope = 2
)

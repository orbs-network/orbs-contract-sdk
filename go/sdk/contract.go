package sdk

// Contract receiver for repository contracts (instantiated only on repository init = system init)
// TODO: consider merging the Contract receiver into the contract Context so we don't have two separate mechanisms

type ContractInstance interface {
	// _init(ctx Context) error
}

type BaseContract struct {
	State   StateSdk
	Service ServiceSdk
	Address AddressSdk
}

func NewBaseContract(
	state StateSdk,
	service ServiceSdk,
	address AddressSdk,
) *BaseContract {

	return &BaseContract{
		State:   state,
		Service: service,
		Address: address,
	}
}

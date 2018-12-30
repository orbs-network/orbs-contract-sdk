package main

import (
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/state"
	. "github.com/orbs-network/orbs-contract-sdk/go/testing/unit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInit(t *testing.T) {
	ownerAddress := AnAddress()

	InSystemScope(ownerAddress, nil, func(m Mockery) {
		_init()

		require.EqualValues(t, TOTAL_SUPPLY, state.ReadUint64ByAddress(ownerAddress))
	})
}

func TestTransfer_NotEnoughBalance(t *testing.T) {
	caller := AnAddress()
	recipient := AnAddress()

	InServiceScope(nil, caller, func(m Mockery) {
		state.WriteUint64ByAddress(caller, 10)
		require.Panics(t, func() {
			transfer(20, recipient)
		})
	})
}

func TestTransfer_NoSourceAddress(t *testing.T) {
	caller := AnAddress()
	recipient := AnAddress()

	InServiceScope(nil, caller, func(m Mockery) {
		require.Panics(t, func() {
			transfer(20, recipient)
		})
	})
}

func TestTransfer_Success(t *testing.T) {
	recipient := AnAddress()
	caller := AnAddress()

	InServiceScope(nil, caller, func(m Mockery) {
		state.WriteUint64ByAddress(caller, 20)
		transfer(20, recipient)

		require.EqualValues(t, 0, state.ReadUint64ByAddress(caller))
		require.EqualValues(t, 20, state.ReadUint64ByAddress(recipient))
	})
}

func TestGetBalance_NoAddress(t *testing.T) {
	InServiceScope(nil, nil, func(m Mockery) {
		require.Zero(t, getBalance(AnAddress()))
	})
}

func TestGetBalance_Success(t *testing.T) {
	address := AnAddress()
	balance := uint64(42)
	InServiceScope(nil, nil, func(m Mockery) {
		state.WriteUint64ByAddress(address, balance)
		require.EqualValues(t, balance, getBalance(address))
	})
}

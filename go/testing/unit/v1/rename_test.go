package v1

import (
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
	. "github.com/orbs-network/orbs-contract-sdk/go/testing/unit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRename(t *testing.T) {
	caller := AnAddress()

	InServiceScope(nil, caller, func(m Mockery) {
		state.WriteString([]byte("artist"), "David Bowie")

		require.EqualValues(t, "David Bowie", state.ReadString([]byte("artist")))
		require.Empty(t, state.ReadString([]byte("performer")))

		state.Rename([]byte("artist"), []byte("performer"))

		require.EqualValues(t, "David Bowie", state.ReadString([]byte("performer")))
		require.Empty(t, state.ReadString([]byte("artist")))
	})
}
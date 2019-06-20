package v1

import (
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKey(t *testing.T) {
	require.EqualValues(t, []byte("hello"), state.Key("hello"))
	require.EqualValues(t, []byte("hello$again"), state.Key("hello", "$", "again"))
}


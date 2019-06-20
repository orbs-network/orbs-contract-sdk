package test

import (
	"fmt"
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/orbs-network/orbs-contract-sdk/go/examples/test"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
	"time"
)

type harness struct {
	client *orbs.OrbsClient
	contractName string
}


func newHarness() *harness {
	return &harness{
		client: orbs.NewClient(test.GetGammaEndpoint(), 42, codec.NETWORK_TYPE_TEST_NET),
		contractName: fmt.Sprintf("MyPhoneBook%d", time.Now().UnixNano()),
	}
}

func (h *harness) deployContract(t *testing.T, sender *orbs.OrbsAccount) {
	contractSource, err := ioutil.ReadFile("../contract.go")
	require.NoError(t, err)

	deployTx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey,
		"_Deployments", "deployService", h.contractName, uint32(1), contractSource)
	require.NoError(t, err)

	deployResponse, err := h.client.SendTransaction(deployTx)
	require.NoError(t, err)

	require.EqualValues(t, codec.EXECUTION_RESULT_SUCCESS, deployResponse.ExecutionResult)
}

func (h *harness) get(t *testing.T, sender *orbs.OrbsAccount) interface{} {
	query, err := h.client.CreateQuery(sender.PublicKey, h.contractName, "get", sender.Address)
	require.NoError(t, err)

	queryResponse, err := h.client.SendQuery(query)
	require.NoError(t, err)

	return queryResponse.OutputArguments[0]
}

func (h *harness) register(t *testing.T, sender *orbs.OrbsAccount, payload string) (*codec.SendTransactionResponse, error) {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.contractName, "register", payload)
	require.NoError(t, err)

	return h.client.SendTransaction(tx)
}
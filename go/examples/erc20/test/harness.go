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
	client       *orbs.OrbsClient
	contractName string
}

func newHarness() *harness {
	return &harness{
		client:       orbs.NewClient(test.GetGammaEndpoint(), 42, codec.NETWORK_TYPE_TEST_NET),
		contractName: fmt.Sprintf("MyERC20%d", time.Now().UnixNano()),
	}
}

func (h *harness) deployContract(t *testing.T, sender *orbs.OrbsAccount) *codec.SendTransactionResponse {
	contractSource, err := ioutil.ReadFile("../erc20.go")
	require.NoError(t, err)

	deployTx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey,
		"_Deployments", "deployService", h.contractName, uint32(1), contractSource)
	require.NoError(t, err)

	deployResponse, err := h.client.SendTransaction(deployTx)
	require.NoError(t, err)

	require.EqualValues(t, codec.EXECUTION_RESULT_SUCCESS, deployResponse.ExecutionResult)

	return deployResponse
}

func (h *harness) balanceOf(t *testing.T, sender *orbs.OrbsAccount) interface{} {
	query, err := h.client.CreateQuery(sender.PublicKey, h.contractName, "balanceOf", sender.AddressAsBytes())
	require.NoError(t, err)

	queryResponse, err := h.client.SendQuery(query)
	require.NoError(t, err)

	return queryResponse.OutputArguments[0]
}

func (h *harness) allowance(t *testing.T, sender *orbs.OrbsAccount, receiver *orbs.OrbsAccount) interface{} {
	query, err := h.client.CreateQuery(sender.PublicKey, h.contractName, "allowance", sender.AddressAsBytes(), receiver.AddressAsBytes())
	require.NoError(t, err)

	queryResponse, err := h.client.SendQuery(query)
	require.NoError(t, err)

	return queryResponse.OutputArguments[0]
}

func (h *harness) allowedTransfer(t *testing.T, sender *orbs.OrbsAccount, receiver *orbs.OrbsAccount) interface{} {
	query, err := h.client.CreateQuery(sender.PublicKey, h.contractName, "allowedTransfer", sender.AddressAsBytes(), receiver.AddressAsBytes())
	require.NoError(t, err)

	queryResponse, err := h.client.SendQuery(query)
	require.NoError(t, err)

	return queryResponse.OutputArguments[0]
}

func (h *harness) transfer(t *testing.T, sender *orbs.OrbsAccount, receiver *orbs.OrbsAccount, sum uint64) (*codec.SendTransactionResponse, error) {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.contractName, "transfer", receiver.AddressAsBytes(), sum)
	require.NoError(t, err)

	return h.client.SendTransaction(tx)
}

func (h *harness) transferFrom(t *testing.T, signer *orbs.OrbsAccount, sender *orbs.OrbsAccount, receiver *orbs.OrbsAccount, sum uint64) (*codec.SendTransactionResponse, error) {
	tx, _, err := h.client.CreateTransaction(signer.PublicKey, signer.PrivateKey, h.contractName, "transferFrom", sender.AddressAsBytes(), receiver.AddressAsBytes(), sum)
	require.NoError(t, err)

	return h.client.SendTransaction(tx)
}

func (h *harness) approve(t *testing.T, sender *orbs.OrbsAccount, receiver *orbs.OrbsAccount, sum uint64) (*codec.SendTransactionResponse, error) {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.contractName, "approve", receiver.AddressAsBytes(), sum)
	require.NoError(t, err)

	return h.client.SendTransaction(tx)
}

func (h *harness) increaseAllowance(t *testing.T, sender *orbs.OrbsAccount, receiver *orbs.OrbsAccount, sum uint64) (*codec.SendTransactionResponse, error) {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.contractName, "increaseAllowance", receiver.AddressAsBytes(), sum)
	require.NoError(t, err)

	return h.client.SendTransaction(tx)
}

func (h *harness) decreaseAllowance(t *testing.T, sender *orbs.OrbsAccount, receiver *orbs.OrbsAccount, sum uint64) (*codec.SendTransactionResponse, error) {
	tx, _, err := h.client.CreateTransaction(sender.PublicKey, sender.PrivateKey, h.contractName, "decreaseAllowance", receiver.AddressAsBytes(), sum)
	require.NoError(t, err)

	return h.client.SendTransaction(tx)
}

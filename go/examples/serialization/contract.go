package main

import (
	"encoding/hex"
	"encoding/json"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/address"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
)

var PUBLIC = sdk.Export(register, get)
var SYSTEM = sdk.Export(_init)

type PhonebookEntry struct {
	FirstName string
	LastName string

	Phone uint64
	OrbsAddress string
}

func _init() {

}

func register(payload string) {
	entry := PhonebookEntry{}
	if err := json.Unmarshal([]byte(payload), &entry); err != nil {
		panic(err)
	}

	if entry.FirstName == "" || entry.LastName == "" || entry.Phone == 0 {
		panic("one of required fields is empty")
	}

	address := hex.EncodeToString(address.GetSignerAddress())
	entry.OrbsAddress = address

	state.SerializeStruct(address, entry)
}

func get(address string) string {
	entry := PhonebookEntry{}
	if err := state.DeserializeStruct(address, &entry); err != nil {
		panic(err)
	}

	rawJson, _ := json.Marshal(entry)
	return string(rawJson)
}

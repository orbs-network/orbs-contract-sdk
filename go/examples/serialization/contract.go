package main

import (
	"encoding/hex"
	"encoding/json"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/address"
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1/state"
	"strings"
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

func toAddress(input string) string {
	if len(input) > 40 {
		input = input[2:42]
	}
	return strings.ToLower(input)
}

func register(payload string) {
	entry := PhonebookEntry{}
	if err := json.Unmarshal([]byte(payload), &entry); err != nil {
		panic(err)
	}

	if entry.FirstName == "" || entry.LastName == "" || entry.Phone == 0 {
		panic("one of required fields is empty")
	}

	address := toAddress(hex.EncodeToString(address.GetSignerAddress()))
	entry.OrbsAddress = address

	state.SerializeStruct(address, entry)
}

func get(address string) string {
	address = toAddress(address)

	entry := PhonebookEntry{}
	if err := state.DeserializeStruct(address, &entry); err != nil {
		panic(err)
	}

	rawJson, _ := json.Marshal(entry)
	return string(rawJson)
}

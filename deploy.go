package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/orbs-network/orbs-network-go/devtools/jsonapi"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var contracts = []string{
	"go/examples/counter",
}

func ClientBinary() []string {
	ciBinaryPath := "./cli"
	if _, err := os.Stat(ciBinaryPath); err != nil {
		fmt.Println("Couldn't locate cli binary, follow GitHub to re-install project")
		os.Exit(2)
	}

	return []string{ciBinaryPath}
}

func runCommand(command []string) string {
	cmd := exec.Command(command[0], command[1:]...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	_ := cmd.Run()

	fmt.Println("jsonapi exec command:", command)
	fmt.Println("command stdout:", stdout.String())
	fmt.Println("command stderr:", stderr.String())

	return stdout.String()
}

func getContractAsString(contractPath string) (string, error) {
	contractAsBytes, err := ioutil.ReadFile(contractPath + "/counter.go")
	if err != nil {
		return "", err
	}

	return string(contractAsBytes), nil
}

func getKeys() ([]string, error) {
	keysFileContent, err := ioutil.ReadFile(".orbsKeys")
	if err != nil {
		return []string{}, err
	}

	return strings.Split(string(keysFileContent), "\n"), nil
}

func generateDeployJSON(contractAsString string) {
	transferJSON := &jsonapi.Transaction{
		ContractName: "BenchmarkToken",
		MethodName:   "transfer",
		Arguments: []jsonapi.MethodArgument{
			{Name: "amount", Type: protocol.METHOD_ARGUMENT_TYPE_STRING_VALUE, StringValue: contractAsString},
		},
	}

	jsonBytes, _ := json.Marshal(&transferJSON)
	return string(jsonBytes)
}

func main() {
	keys, err := getKeys()
	publicKey := keys[0]
	privateKey := keys[1]

	if err != nil {
		fmt.Println("Could not find your client keys, try to re-generate them by running ./generate-test-keys.sh from your workspace")
		os.Exit(2)
	}

	baseCommand := ClientBinary()

	for _, contract := range contracts {
		deployCommand := append(baseCommand,
			"-send-transaction", generateDeployJSON(),
			"-public-key", publicKey,
			"-private-key", privateKey)

		//runCommand()
		fmt.Println(getContractAsString(contract))
	}
}

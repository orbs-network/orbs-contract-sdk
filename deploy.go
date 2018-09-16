package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

//// Generate calls to the cli to deploy these contracts
//func runCommand(command []string, t *testing.T) string {
//	cmd := exec.Command(command[0], command[1:]...)
//	var stdout, stderr bytes.Buffer
//	cmd.Stdout = &stdout
//	cmd.Stderr = &stderr
//	err := cmd.Run()
//
//	fmt.Println("jsonapi exec command:", command)
//	fmt.Println("command stdout:", stdout.String())
//	fmt.Println("command stderr:", stderr.String())
//
//	require.NoError(t, err, "jsonapi cli command should not fail")
//
//	return stdout.String()
//}

func getContractAsString(contractPath string) (string, error) {
	contractAsBytes, err := ioutil.ReadFile(contractPath + "/contract.go")
	if err != nil {
		return "", err
	}

	return string(contractAsBytes), nil
}

func main() {
	for _, contract := range contracts {
		fmt.Println(getContractAsString(contract))
	}
}

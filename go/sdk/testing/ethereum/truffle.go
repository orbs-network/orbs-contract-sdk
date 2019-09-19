// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package truffle

import (
	"fmt"
	"os/exec"
	"strings"
)

type cli struct {
	projectPath string
}

func Cli(projectPath string) *cli {
	return &cli{
		projectPath: projectPath,
	}
}

func (c *cli) Run(args string) string {
	fmt.Printf("*** RUNNING: truffle %s\n", args)

	argsArr := strings.Split(args, " ")
	cmd := exec.Command("truffle", argsArr...)
	cmd.Dir = c.projectPath
	out, err := cmd.CombinedOutput()

	fmt.Printf("*** OUTPUT:\n%s\n", string(out))

	if err != nil {
		fmt.Printf("Make sure Ethereum truffle is installed, found in your $PATH and working in terminal.\nTry installing with 'yarn global add truffle' or see instructions in https://truffleframework.com/docs/truffle/getting-started/installation\n\n")
		panic(fmt.Sprintf("truffle failed: %s", err.Error()))
	}

	return string(out)
}

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

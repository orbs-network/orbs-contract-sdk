package gamma

import (
	"fmt"
	"os/exec"
	"strings"
)

type cli struct {
	port      string
	isStarted bool
}

func Cli() *cli {
	return &cli{}
}

func (c *cli) SetPort(port int) *cli {
	c.port = fmt.Sprintf("%d", port)
	return c
}

func (c *cli) Run(args string) string {
	if len(c.port) > 0 && len(args) > 0 {
		args = args + " -port " + c.port
	}

	if !c.isStarted {
		c.Start()
	}

	fmt.Printf("*** RUNNING: gamma-cli %s\n", args)

	argsArr := strings.Split(args, " ")
	out, err := exec.Command("gamma-cli", argsArr...).CombinedOutput()

	fmt.Printf("*** OUTPUT:\n%s\n", string(out))

	if err != nil {
		fmt.Printf("Make sure gamma-cli is installed, found in your $PATH and working in terminal.\nSee instructions in https://github.com/orbs-network/orbs-contract-sdk/blob/master/GAMMA.md\n\n")
		panic(fmt.Sprintf("gamma-cli failed: %s", err.Error()))
	}

	return string(out)
}

func (c *cli) Start() *cli {
	if c.isStarted {
		return c
	}
	c.isStarted = true
	c.Run("start-local -wait")
	return c
}

func (c *cli) Stop() {
	if !c.isStarted {
		return
	}
	c.Run("stop-local")
	c.isStarted = false
}

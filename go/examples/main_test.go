package examples

import (
	"github.com/orbs-network/orbs-contract-sdk/go/testing/gamma"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	gamma := gamma.Cli().Start()
	defer gamma.Stop()

	time.Sleep(5*time.Second)

	exitCode := m.Run()

	os.Exit(exitCode)
}
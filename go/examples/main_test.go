package examples_test

import (
	"github.com/orbs-network/orbs-contract-sdk/go/testing/gamma"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	gamma.Cli().Start()
	// it's dirty but there is no other way
	// defer gamma.Stop()

	time.Sleep(5*time.Second)

	exitCode := m.Run()

	os.Exit(exitCode)
}
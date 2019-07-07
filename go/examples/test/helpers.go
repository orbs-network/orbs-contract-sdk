package test

import (
	"fmt"
	"os"
	"time"
)

func GetGammaEndpoint() string {
	if localIP := os.Getenv("LOCAL_IP"); localIP != "" {
		return fmt.Sprintf("http://%s:8080", localIP)
	}

	return "http://localhost:8080"
}

const eventuallyIterations = 50

func Eventually(timeout time.Duration, f func() bool) bool {
	for i := 0; i < eventuallyIterations; i++ {
		if f() {
			return true
		}
		time.Sleep(timeout / eventuallyIterations)
	}
	return false
}

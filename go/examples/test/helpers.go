package test

import (
	"fmt"
	"os"
)

func GetGammaEndpoint() string {
	if localIP := os.Getenv("LOCAL_IP"); localIP != "" {
		return fmt.Sprintf("http://%s:8080", localIP)
	}

	return "http://localhost:8080"
}

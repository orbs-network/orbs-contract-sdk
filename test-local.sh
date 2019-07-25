#!/bin/bash -e

# This test requires a global setup in the shape of launching a Docker container before we run go test.

# Shut down any existing Gamma server container
echo "Clearing up any existing gamma servers.."
docker ps -a | grep orbsnetwork\/gamma | awk '{print $1}' | xargs docker rm -fv

echo "Starting up gamma server.."
docker run -d -p "8080:8080" orbsnetwork/gamma:experimental

echo "Gamma server started and listening on port 8080"

sleep 5

echo "Running SDK tests..."
go test -count=1 -v ./go/testing/...
echo "*********************"
echo "SDK tests passed!"

echo "Running context tests..."
go test -count=1 -v ./go/context/...
echo "*********************"
echo "Context tests passed!"
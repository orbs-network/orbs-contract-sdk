#!/bin/bash -e

echo "Sleeping 5 seconds to allow Gamma server to wake up..."
sleep 5

echo "Running SDK tests..."
go test -count=1 -v ./go/testing/...
echo "*********************"
echo "SDK tests passed!"

echo "Running context tests..."
go test -count=1 -v ./go/context/...
echo "*********************"
echo "Context tests passed!"
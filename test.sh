#!/bin/bash -e

echo "Sleeping 5 seconds to allow Gamma server to wake up..."
sleep 5

cd go/sdk
echo "Running SDK tests..."
go test -count=1 -v ./testing/...
echo "*********************"
echo "SDK tests passed!"

echo "Running context tests..."
go test -count=1 -v ./context/...
echo "*********************"
echo "Context tests passed!"
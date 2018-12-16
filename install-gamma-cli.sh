#!/usr/bin/env bash

go get -u github.com/orbs-network/orbs-client-sdk-go
cd $GOPATH
cd src/github.com/orbs-network/orbs-client-sdk-go/gammacli
rm -rf ./_bin
mkdir -p ./_bin
go build -o ./_bin/gamma-cli
chmod +x ./_bin/gamma-cli

# Let's try and run it.
./bin/gamma-cli




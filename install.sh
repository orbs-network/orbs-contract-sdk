#!/usr/bin/env bash

cd `echo $GOPATH`

echo "Installing Orbs Network Go, Please wait.."
go get github.com/orbs-network/orbs-network-go
cd src/github.com/orbs-network/orbs-network-go && sh git-submodule-checkout.sh

echo "Installing the Developer SDK, Please wait.."
go get github.com/orbs-network/orbs-contract-sdk
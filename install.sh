#!/usr/bin/env bash

cd `echo $GOPATH`

echo "Installing Orbs Network Go, Please wait.."
go get github.com/orbs-network/orbs-network-go
cd src/github.com/orbs-network/orbs-network-go
git fetch --all && git checkout feature/gamma-cli
rm -rf vendor
./git-submodule-checkout.sh

echo "Installing the Developer SDK, Please wait.."
go get github.com/orbs-network/orbs-contract-sdk 2> /dev/null

echo "Compiling gamma-server binary.."
rm -f ./gamma-server ./gamma-cli
go build -o ./gamma-server devtools/sambusac/main/main.go
echo "Compiling gamma-cli binary.."
go build -o ./gamma-cli devtools/jsonapi/main/main.go

cd `echo $GOPATH`
cd src/github.com/orbs-network/orbs-contract-sdk
ln -s ../orbs-network-go/gamma-cli gamma-cli
ln -s ../orbs-network-go/gamma-server gamma-server

# Create global symlinks
cd `echo $GOPATH`
sudo ln -s $GOPATH/src/github.com/orbs-network/orbs-network-go/gamma-cli /usr/local/bin/gamma-cli
sudo ln -s $GOPATH/src/github.com/orbs-network/orbs-network-go/gamma-server /usr/local/bin/gamma-server

cd `echo $GOPATH`
cd src/github.com/orbs-network/orbs-contract-sdk

echo "Generating test keys for gamma-cli to use"
./generate_test_keys.sh

echo "Workspace created successfully!"
echo "You can begin experimenting with Orbs Network"
echo "The workspace is located under"
echo "$GOPATH/src/github.com/orbs-network/orbs-contract-sdk"

./gamma-cli

exit 0
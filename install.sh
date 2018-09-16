#!/usr/bin/env bash

cd `echo $GOPATH`

echo "Installing Orbs Network Go, Please wait.."
go get github.com/orbs-network/orbs-network-go
cd src/github.com/orbs-network/orbs-network-go
sh git-submodule-checkout.sh
git fetch --all && git checkout feature/deployable-go

echo "Compiling Sambusac.."
rm -f ./sambusac ./cli
go build -o ./sambusac devtools/sambusac/main/main.go
echo "Compiling Orbs CLI.."
go build -o ./cli devtools/jsonapi/main/main.go

echo "Installing the Developer SDK, Please wait.."
go get github.com/orbs-network/orbs-contract-sdk 2> /dev/null

cd `echo $GOPATH`
cd src/github.com/orbs-network/orbs-contract-sdk
ln -s ../orbs-network-go/cli cli
ln -s ../orbs-network-go/sambusac sambusac

echo "Workspace created successfully!"
echo "You can begin experimenting with Orbs Network"
echo "The workspace is located under"
echo "$GOPATH/src/github.com/orbs-network/orbs-contract-sdk"

exit 0
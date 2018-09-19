#!/usr/bin/env bash

echo "Generating testing keys.."
./gamma-cli genKeys > .orbsKeys

echo "Keys generated and saved under the following path"
echo "$GOPATH/src/github.com/orbs-contract-sdk/.orbsKeys"

exit 0
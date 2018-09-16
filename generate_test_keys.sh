#!/usr/bin/env bash

echo "Generating testing keys.."
./cli -generate-test-keys > .orbsKeys

echo "Keys generated and saved under the following path"
echo "$GOPATH/src/github.com/orbs-contract-sdk/.orbsKeys"

exit 0
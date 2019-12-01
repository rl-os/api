#!/bin/bash

set -e

goimports -d $(find . -type f -name '*.go' -not -path "./vendor/*")

go vet ./...
./bin/golangci-lint run | grep -v go | grep -v SA1019 | cat

# shellcheck disable=SC2046
gocyclo -over 55 $(find . -iname '*.go' -type f | grep -v /vendor/)
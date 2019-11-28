#!/bin/bash
bash scripts/message.sh

set -e

echo "  >  Checking if there is any missing dependencies..."
go mod download

echo "  >  Cleaning build cache"
go clean

echo "  >  Installing unnessosory packages"
go get -v golang.org/x/lint/golint
go get -v github.com/fzipp/gocyclo
go get -v golang.org/x/tools/cmd/goimports
wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v1.15.0

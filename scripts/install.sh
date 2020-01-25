#!/usr/bin/env bash
bash scripts/message.sh

set -e

echo "  >  Checking if there is any missing dependencies..."
go mod download
go get -u github.com/amacneil/dbmate

echo "  >  Installing unnessosory packages"
go get -u golang.org/x/lint/golint
go get -u github.com/fzipp/gocyclo
go get -u golang.org/x/tools/cmd/goimports
wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v1.15.0

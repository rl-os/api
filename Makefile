# --- Makefile ----

# setup enviroment
SHELL := /bin/bash
export GO111MODULE := on
export PATH := bin:$(PATH)

PACKAGE				   = ayako
VERSION       	       ?= $(shell git describe --tags --always --match="v*" 2> /dev/null || cat $(CURDIR)/.version 2> /dev/null || echo v1.0.0)
COMMIT                 ?= $(shell git rev-parse HEAD)
GIT_BRANCH             := $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null)
GIT_BRANCH_CLEAN       := $(shell echo $(GIT_BRANCH) | sed -e "s/[^[:alnum:]]/-/g")
BUILDTIME              := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

DOCKER_REGISTRY_DOMAIN ?= docker.pkg.github.com
DOCKER_REGISTRY_PATH   ?= deissh/osu-lazer
DOCKER_IMAGE           ?= $(DOCKER_REGISTRY_PATH)/$(PACKAGE):$(VERSION)
DOCKER_IMAGE_DOMAIN    ?= $(DOCKER_REGISTRY_DOMAIN)/$(DOCKER_IMAGE)

# ===============================================
DOCKER                  = docker
DOCKERBUILD            := $(DOCKER) build
DOCKERPUSH             := $(DOCKER) push
# ===============================================
GOCMD                   = go
GOLINT				   := $(shell which golint)
GOIMPORT			   := $(shell which goimports)
GOFMT				   := $(shell which gofmt)
GOCYCLO				   := $(shell which gocyclo)
GOGENERATE 			   := $(GOCMD) generate
GOBUILD				   := $(GOCMD) build
GOCLEAN				   := $(GOCMD) clean
GOTEST				   := $(GOCMD) test
GOMOD				   := $(GOCMD) mod
GOGET				   := $(GOCMD) get
GOLIST				   := $(GOCMD) list
GOVET				   := $(GOCMD) vet
# ===============================================
# build setting
CMD_DIR                ?= $(PWD)/cmd
BIN_DIR                ?= $(PWD)/bin
BUILD_CMDS             := $(shell $(GOLIST) $(CMD_DIR)/...)
GOFILES				   := $(shell find . -name "*.go" -type f)

GOLDFLAGS 			   += -X main.Version=$(VERSION)
GOLDFLAGS 			   += -X main.Commit=$(COMMIT)
GOLDFLAGS  			   += -X main.Branch=$(GIT_BRANCH_CLEAN)
GOLDFLAGS 			   += -X main.BuildTimestamp=$(BUILDTIME)
GOBUILDFLAGS            = -ldflags "-s -w $(GOLDFLAGS)"

# ===============================================
.DEFAULT_GOAL          := help
MAKE_ENV 			   += PACKAGE VERSION DOCKER_IMAGE DOCKER_IMAGE_DOMAIN
SHELL_EXPORT		   := $(foreach v,$(MAKE_ENV),$(v)='$($(v))' )
# check requeres commands
EXECUTABLES             = wget
_                      := $(foreach exec,$(EXECUTABLES), $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH")))
# pass all args
args                    = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`
# ===============================================

## Clear bin folder
.PHONY: clear
clear:
	@echo -e "\e[1;34m> Cleaning stage\e[0m"
	@rm -rf bin

## Set impotants flags and then build all commands in cmd folder
.PHONY: build
build: clear generate
	@echo -e "\e[1;34m> Building stage\e[0m"
	$(GOBUILD) $(GOBUILDFLAGS) -o $(BIN_DIR)/server $(CMD_DIR)

## Build all commands in cmd folder as prod-like
.PHONY: build-prod
build-prod: clear generate
	@echo -e "\e[1;34m> Building stage\e[0m"
	CGO_ENABLED=0 $(GOBUILD) $(GOBUILDFLAGS) -a -installsuffix nocgo -o $(BIN_DIR)/server $(CMD_DIR)

## Run all checks
.PHONY: lint
lint: vet fmt-check
	@for PKG in $(PACKAGES); do golint -set_exit_status $$PKG || exit 1; done;
	@$(GOIMPORT) -d $(shell find . -type f -name '*.go' -not -path "./vendor/*")
	@$(GOCYCLO) -over 55 $(shell find . -iname '*.go' -type f | grep -v /vendor/)

## Install missing dependencies
## Also setup unnessosory packages that using in CI
.PHONY: install
install:
	@echo -e "\e[1;34m> Checking if there is any missing dependencies...\e[0m"
	$(GOMOD) download
	@echo -e "\e[1;34m> Installing generators and etc packages\e[0m"
	$(GOGET) -u github.com/amacneil/dbmate
	$(GOGET) -u github.com/google/wire/cmd/wire
	$(GOGET) -u github.com/golang/mock/mockgen@latest
	$(GOGET) -u github.com/hexdigest/gowrap/cmd/gowrap
	@echo -e "\e[1;34m> Installing unnessosory packages\e[0m"
	$(GOGET) -u golang.org/x/lint/golint
	$(GOGET) -u github.com/fzipp/gocyclo
	$(GOGET) -u golang.org/x/tools/cmd/goimports
	wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v1.15.0


## This help message
## Which can also be multiline
.PHONY: help
help: message
	@printf " Usage\n";

	@awk '{ \
			if ($$0 ~ /^.PHONY: [a-zA-Z\-\_0-9]+$$/) { \
				helpCommand = substr($$0, index($$0, ":") + 2); \
				if (helpMessage) { \
					printf " \033[36m%-20s\033[0m %s\n", \
						helpCommand, helpMessage; \
					helpMessage = ""; \
				} \
			} else if ($$0 ~ /^[a-zA-Z\-\_0-9.]+:/) { \
				helpCommand = substr($$0, 0, index($$0, ":")); \
				if (helpMessage) { \
					printf "  \033[36m%-20s\033[0m %s\n", \
						helpCommand, helpMessage; \
					helpMessage = ""; \
				} \
			} else if ($$0 ~ /^##/) { \
				if (helpMessage) { \
					helpMessage = helpMessage"\n                      "substr($$0, 3); \
				} else { \
					helpMessage = substr($$0, 3); \
				} \
			} else { \
				if (helpMessage) { \
					print "\n\n                      "helpMessage"\n" \
				} \
				helpMessage = ""; \
			} \
		}' \
		$(MAKEFILE_LIST)

## -- Database --
## DBMate wrap with custom ENV and disabled schema dump

## Migrate all files and commit changes
## Require CONFIG__DATABASE__DSN
.PHONY: db-migrate
db-migrate:
	@dbmate --wait --env CONFIG__DATABASE__DSN \
	--no-dump-schema -d migrations -s schema.sql \
	up

## Rollback last changes
## Require CONFIG__DATABASE__DSN
.PHONY: db-rollback
db-rollback:
	@dbmate --wait --env CONFIG__DATABASE__DSN \
	--no-dump-schema -d migrations -s schema.sql \
	up

## Validate migrations
## Require CONFIG__DATABASE__DSN
.PHONY: db-status
db-status:
	@dbmate --wait --env CONFIG__DATABASE__DSN \
	--no-dump-schema -d migrations -s schema.sql \
	status


## -- Docker --

## Build docker container
.PHONY: docker-build
docker-build:
	@$(DOCKERBUILD) -t $(DOCKER_IMAGE_DOMAIN) -f Dockerfile ..

## Push docker container to registry
.PHONY: docker-push
docker-push:
	@$(DOCKERPUSH) $(DOCKER_IMAGE_DOMAIN)

## -- Utils --

.PHONY: message
message:
	@clear
	@echo -e "\033[36m\n                                        \033[0m _                   _         \033[36m\n        _.====.._                    \033[0m   /_\    _  _   __ _  | |__  ___ \033[36m\n     ,:.         ~-_                 \033[0m  / _ \  | || | / _\` | | / / / _ \ \033[36m\n         \`\        ~-_               \033[0m /_/ \_\  \_, | \__,_| |_\_\ \___/\033[36m\n           |          \`.              \033[0m         |__/                    \033[36m\n         ,/             ~-_                        \033[0mOAuth and API server\033[36m\n-..__..-''                 ~~--..__...----...--.....---.....--....---..\033[0m\n\n                                                      https://risu.life\n                                    https://github.com/deissh/rl\n\n                                                      2019-2020, deissh\n\n"


## -- Go commands --

## Run go generate on current project
.PHONY: generate
generate:
	@echo -e "\e[1;34m> Go generate\e[0m"
	@$(GOGENERATE) ./...

## Run go vet on current project
.PHONY: vet
vet:
	@echo -e "\e[1;34m> Go vet\e[0m"
	$(GOVET) ./...

## Run go fmt on current project
.PHONY: fmt
fmt:
	@echo -e "\e[1;34m> Go fmt\e[0m"
	$(GOFMT) -s -w $(GOFILES)

## Run check fmt and show message if not all ok
.PHONY: fmt-check
fmt-check:
	@echo -e "\e[1;34m> Go fmt-check\e[0m"
	@diff=$$($(GOFMT) -s -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

## Run go test on current project
.PHONY: test
test:
	@echo -e "\e[1;34m> Go test\e[0m"
	$(GOTEST) -v -race -coverprofile=coverage.txt -cover -coverpkg=$(shell go list ./... | grep -v mocks | tr '\n' ',') ./...

## Generate coverage report and then show
.PHONY: view-covered
view-covered:
	@echo -e "\e[1;34m> Opening coverage\e[0m"
	$(GOCMD) tool cover -html=coverage.txt

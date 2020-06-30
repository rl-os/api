# --- Makefile ----

# setup enviroment
SHELL := /bin/bash
export GO111MODULE := on
export PATH := bin:$(PATH)

PACKAGE				   = api
VERSION       	       ?= $(shell git describe --tags --always --dirty --match="v*" 2> /dev/null || cat $(CURDIR)/.version 2> /dev/null || echo v1.0.0)
COMMIT                 ?= $(shell git rev-parse HEAD)
GIT_BRANCH             := $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null)
GIT_BRANCH_CLEAN       := $(shell echo $(GIT_BRANCH) | sed -e "s/[^[:alnum:]]/-/g")
BUILDTIME              := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

DOCKER_REGISTRY_DOMAIN ?= docker.pkg.github.com
DOCKER_REGISTRY_PATH   ?= deissh
DOCKER_IMAGE           ?= $(DOCKER_REGISTRY_PATH)/$(PACKAGE):$(VERSION)
DOCKER_IMAGE_DOMAIN    ?= $(DOCKER_REGISTRY_DOMAIN)/$(DOCKER_IMAGE)

# ===============================================
DOCKERCMD				= docker
GOCMD                   = go
GOLINT				   := $(shell which golint)
GOIMPORT			   := $(shell which goimports)
GOFMT				   := $(shell which gofmt)
GOCYCLO				   := $(shell which gocyclo)
GOBUILD				   := $(GOCMD) build
GOCLEAN				   := $(GOCMD) clean
GOTEST				   := $(GOCMD) test
GOMOD				   := $(GOCMD) mod
GOGET				   := $(GOCMD) get
GOLIST				   := $(GOCMD) list
GOVET				   := $(GOCMD) vet
# ===============================================
# build setting
CMD_DIR                ?= ./cmd
BIN_DIR                ?= $(PWD)/bin
BUILD_CMDS             := $(shell $(GOLIST) ./$(CMD_DIR)/...)
GOFILES				   := $(shell find . -name "*.go" -type f)

GOLDFLAGS += -X main.Version=$(VERSION)
GOLDFLAGS += -X main.Commit=$(COMMIT)
GOLDFLAGS += -X main.Branch=$(GIT_BRANCH_CLEAN)
GOLDFLAGS += -X main.BuildTimestamp=$(BUILDTIME)
GOBUILDFLAGS = -ldflags "-s -w $(GOLDFLAGS)"

# ===============================================
GO = go
.DEFAULT_GOAL := help
MAKE_ENV += PACKAGE VERSION DOCKER_IMAGE DOCKER_IMAGE_DOMAIN
SHELL_EXPORT := $(foreach v,$(MAKE_ENV),$(v)='$($(v))' )
# check requeres commands
EXECUTABLES = wget
_ := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH")))
# ===============================================

## Clear bin folder
.PHONY: clear
clear:
	@rm -rf bin

## Set impotants flags and then build all commands in cmd folder
.PHONY: build
build: clear
	@(cd cmd ; for CMD in *; do $(GOBUILD) $(GOBUILDFLAGS) -o $(BIN_DIR)/$$CMD $$CMD/*.go && echo "$$CMD build done" || exit 1; done;)

## Build all commands in cmd folder as prod-like
.PHONY: build-prod
build-prod: clear
	@(cd cmd ; for CMD in *; do CGO_ENABLED=0 $(GOBUILD) $(GOBUILDFLAGS) -a -installsuffix nocgo -o $(BIN_DIR)/$$CMD $$CMD/*.go && echo "$$CMD build done" || exit 1; done;)

.PHONY: docker-build
docker-build: message test
	$(DOCKERCMD) build --tag $(DOCKER_IMAGE_DOMAIN) --file Dockerfile .


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
	@echo ">  Checking if there is any missing dependencies..."
	@$(GOMOD) download
	@$(GOGET) -u github.com/amacneil/dbmate
	@echo ">  Installing unnessosory packages"
	$(GOGET) -u golang.org/x/lint/golint
	$(GOGET) -u github.com/fzipp/gocyclo
	$(GOGET) -u golang.org/x/tools/cmd/goimports
	$(GOGET) -u github.com/git-chglog/git-chglog/cmd/git-chglog
	wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v1.15.0


## This help message
## Which can also be multiline
.PHONY: help
help: message
	@printf "Usage\n";

	@awk '{ \
			if ($$0 ~ /^.PHONY: [a-zA-Z\-\_0-9]+$$/) { \
				helpCommand = substr($$0, index($$0, ":") + 2); \
				if (helpMessage) { \
					printf "\033[36m%-20s\033[0m %s\n", \
						helpCommand, helpMessage; \
					helpMessage = ""; \
				} \
			} else if ($$0 ~ /^[a-zA-Z\-\_0-9.]+:/) { \
				helpCommand = substr($$0, 0, index($$0, ":")); \
				if (helpMessage) { \
					printf "\033[36m%-20s\033[0m %s\n", \
						helpCommand, helpMessage; \
					helpMessage = ""; \
				} \
			} else if ($$0 ~ /^##/) { \
				if (helpMessage) { \
					helpMessage = helpMessage"\n                     "substr($$0, 3); \
				} else { \
					helpMessage = substr($$0, 3); \
				} \
			} else { \
				if (helpMessage) { \
					print "\n                     "helpMessage"\n" \
				} \
				helpMessage = ""; \
			} \
		}' \
		$(MAKEFILE_LIST)


## -- Utils --

## Generate CHANGELOG.md based on git
.PHONY: chglog
chglog:
	git-chglog -o CHANGELOG.md $(shell git describe --tags $(shell git rev-list --tags --max-count=1))

.PHONY: message
message:
	@echo -e "\n        _.====.._                                 ___                _ \n     ,:.         ~-_                             / _ \   ___  _  _  | |\n         \`\        ~-_                          | (_) | (_-< | || | |_|\n           |          \`.                         \___/  /__/  \_,_| (_)\n         ,/             ~-_                                lazer server\n-..__..-''                 ~~--..__...----...--.....---.....--....---...\n\n                                   GitHub: github.com/deissh/rl\n                                                     2019-2020, deissh\n\n"


## -- Go commands --

## Run go vet on current project
.PHONY: vet
vet:
	$(GOVET) ./...

## Run go fmt on current project
.PHONY: fmt
fmt:
	$(GOFMT) -s -w $(GOFILES)

## Run check fmt and show message if not all ok
.PHONY: fmt-check
fmt-check:
	@diff=$$($(GOFMT) -s -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

## Run go test on current project
.PHONY: test
test:
	$(GOTEST) ./...

## Generate coverage report and then show
.PHONY: view-covered
view-covered:
	$(GOTEST) -coverprofile=cover.out ./...
	$(GOCMD) tool cover -html=cover.out

.PHONY: load-config
load-config:
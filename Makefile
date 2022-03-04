# go
GO_CMD = go
GO_BUILD = $(GO_CMD) build
GO_MOD = $(GO_CMD) mod
GO_GET = $(GO_CMD) get
GO_MOD_DOWNLOAD = $(GO_MOD) download
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
BINARY_DIR=bin
BINARY_NAME=server


# git
#GIT_COMMIT_HASH:=$(shell git rev-parse HEAD)

# workspace
ROOT = $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))


test:
	echo $(GOOS) $(GOARCH)
install:
	$(GO_CMD) mod tidy
	cd service/gorm_gen && packr2 clean &&	packr2
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO_BUILD)  .
	chmod +x x
	cp -f x /usr/local/bin/x
.PHONY:install
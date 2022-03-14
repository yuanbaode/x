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


IMPORT_PATH		:=github.com/yuanbaode/x/cmd
BUILD_TIME		:=$(shell date "+%F %T")
COMMIT_ID       :=$(shell git rev-parse HEAD)
GO_VERSION      :=$(shell $(GO_CMD) version)
VERSION			:=$(shell git describe --tags)
BUILD_USER		:=$(shell whoami)
PACKAGES_URL	:=$(github.com/YouCD/esDump)
FLAG			:="-X '${IMPORT_PATH}.buildTime=${BUILD_TIME}' -X '${IMPORT_PATH}.commitID=${COMMIT_ID}' -X '${IMPORT_PATH}.goVersion=${GO_VERSION}' -X '${IMPORT_PATH}.goVersion=${GO_VERSION}' -X '${IMPORT_PATH}.Version=${VERSION}' -X '${IMPORT_PATH}.buildUser=${BUILD_USER}'"

# workspace
ROOT = $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))


test:
	echo $(GOOS) $(GOARCH)
install:
	$(GO_CMD) mod tidy
	cd service/gorm_gen && packr2 clean &&	packr2
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH)  $(GO_BUILD) -ldflags $(FLAG) .
	chmod +x x
	cp -f x /usr/local/bin/x
.PHONY:install
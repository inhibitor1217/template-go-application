# Project directory structure
MODULE_NAME := $(shell go list -m)
PROJECT_PATH := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
PROJECT_NAME := $(shell basename $(PROJECT_PATH))
export PATH := ${PATH}:${GOPATH}/bin

# Artifacts
TARGET_DIR ?= ${PROJECT_PATH}/target
TARGET_BIN_DIR ?= ${TARGET_DIR}/bin
GENERATED_SRC_DIR ?= ${PROJECT_PATH}/generated

# Application environment
STAGE ?= development
VERSION := $(shell git describe --exact-match --tags HEAD 2>/dev/null || git rev-parse --abbrev-ref HEAD)

# Go environment
GOVERSION := $(shell go version | awk '{print $$3}')
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
# GOPRIVATE ?= github.com/your-private-repo
LDFLAGS := -ldflags=""

## ------------- TARGETS -------------

env:
	@echo "PROJECT_PATH:\t${PROJECT_PATH}"
	@echo "PROJECT_NAME:\t${PROJECT_NAME}"
	@echo "MODULE_NAME:\t${MODULE_NAME}"
	@echo "GOVERSION:\t${GOVERSION}"
	@echo "GOOS:\t\t${GOOS}"
	@echo "GOARCH:\t\t${GOARCH}"
	@echo "STAGE:\t\t${STAGE}"
	@echo "VERSION:\t${VERSION}"

init:
	go mod download
# If you have private repos, you can use this command to download them
#	GOPRIVATE=${GOPRIVATE} go mod download

build: init
	GOOS=${GOOS} \
	GOARCH=${GOARCH} \
	go build ${LDFLAGS} \
	-o ${TARGET_BIN_DIR}/${PROJECT_NAME}.${GOOS}.${GOARCH} \
	${PROJECT_PATH}/cmd

run:
	${TARGET_BIN_DIR}/${PROJECT_NAME}.${GOOS}.${GOARCH}

dev: build run

clean: clean-gen clean-target

clean-gen:
	rm -rf ${GENERATED_SRC_DIR}

clean-target:
	rm -rf ${TARGET_DIR}

test: build
	go clean -testcache
	go test `go list ./... | grep -v ./generated`

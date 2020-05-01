SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

#ifeq ($(origin .RECIPEPREFIX), undefined)
#    $(error This version of Make does not support .RECIPEPREFIX. Please use GNU Make 4.0 or later)
#endif
#.RECIPEPREFIX = >

.DEFAULT_GOAL := help

TARGET=./target
VERSION=

GRAY=\033[1;90m
MAGENTA=\033[1;35m
RESET_COLOR=\033[0m

## build: build the application
build: clean fmt lint
	@echo -e "${GRAY}>> ⚙️\t${MAGENTA}Building...${RESET_COLOR}"
	@if [ ! -d ${TARGET} ] ; then mkdir -p ${TARGET} ; fi
	@go mod tidy -v
	@go build -o ${TARGET} -v -a -trimpath ./...
.PHONY: build

## test: execute tests of all packages
test: lint
	@echo -e "${GRAY}>> 🧪\t${MAGENTA}Testing...${RESET_COLOR}"
	@go test -v -count=1 -race -trimpath ./...
.PHONY: test

## fmt: format all Go source files
fmt:
	@echo -e "${GRAY}>> ✏️\t${MAGENTA}Formatting...${RESET_COLOR}"
	@go fmt ./...
.PHONY: fmt

## lint: examine Go source code and report suspicious constructs
lint:
	@echo -e "${GRAY}>> 🔎\t${MAGENTA}Linting...${RESET_COLOR}"
	@go vet ./...
.PHONY: lint

## clean: clean the binary
clean:
	@echo -e "${GRAY}>> 🧹\t${MAGENTA}Cleaning...${RESET_COLOR}"
	@if [ -d ${TARGET} ] ; then rm -rv ${TARGET} ; fi
.PHONY: clean

## run: run main.go
run: build
	@echo -e "${GRAY}>> 👟\t${MAGENTA}Running...${RESET_COLOR}"
	@go run -race -trimpath ./cmd/gogrow
.PHONY: run

## mod-download: download modules to local cache
mod-download:
	@echo -e "${GRAY}>> ⬇️\t${MAGENTA}Downloading modules...${RESET_COLOR}"
	@go mod download
.PHONY: mod-download

## build-release [VERSION]: build the application for release
build-release: clean lint test
	if [ -z ${VERSION} ]; then echo "'VERSION' is not set"; exit 1; fi

	#darwin/amd64
	make GOOS='darwin' GOARCH='amd64' TARGET=./target/release/darwin/amd64 build
	pushd ./target/release/darwin/amd64
	tar czf ./../../gogrow-${VERSION}-darwin_amd64.tar.gz ./gogrow
	popd

	#linux/amd64
	make GOOS='linux' GOARCH='amd64' TARGET=./target/release/linux/amd64 build
	pushd ./target/release/linux/amd64
	tar czf ./../../gogrow-${VERSION}-linux_amd64.tar.gz ./gogrow
	popd

	#linux/arm
	make GOOS='linux' GOARCH='arm' TARGET=./target/release/linux/arm build
	pushd ./target/release/linux/arm
	tar czf ./../../gogrow-${VERSION}-linux_arm.tar.gz ./gogrow
	popd

	#linux/arm64
	make GOOS='linux' GOARCH='arm64' TARGET=./target/release/linux/arm64 build
	pushd ./target/release/linux/arm64
	tar czf ./../../gogrow-${VERSION}-linux_arm64.tar.gz ./gogrow
	popd

	#windows/amd64
	make GOOS='windows' GOARCH='amd64' TARGET=./target/release/windows/amd64 build
	pushd ./target/release/windows/amd64
	zip -r ./../../gogrow-${VERSION}-windows_amd64.zip ./gogrow.exe
	popd
.PHONY: build-release

## help: prints this help message
help:
	@echo -e "Usage: make [target] ...\n"
	@echo -e "Targets:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
.PHONY: help

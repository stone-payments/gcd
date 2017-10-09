PROJECT := gcd
BIN_DIR := $(GOPATH)/bin
RELEASE_DIR := ./bin

VERSION ?= v0.0.1
PLATFORM ?= linux
ARCH ?= amd64
RELEASE_PATH := $(RELEASE_DIR)/$(PROJECT)

PATH_MAIN_PACKAGE := ./
GOMETALINTER := $(BIN_DIR)/gometalinter
PKGS := $(shell go list ./... | grep -v /vendor)

.PHONY: test lint build clear

build: lint test clear $(RELEASE_PATH)

build-image:
	@echo "---> Building the project using Dockerfile"
	@docker build -t guiferpa/$(PROJECT):$(VERSION) .

clear:
	@echo "---> Cleaning up directory"
	@rm -rf $(RELEASE_DIR)

test:
	@echo "---> Testing"
	@go test -v -cover

$(RELEASE_PATH):
	@echo "---> Building the project"
	@mkdir -p $(RELEASE_DIR)
	@CGO_ENABLED=0 GOOS=$(PLATFORM) GOARCH=$(ARCH) go build -o $(RELEASE_PATH) $(PATH_MAIN_PACKAGE)

lint: $(GOMETALINTER)
	@echo "---> Running lint"
	@gometalinter ../$(PROJECT)/... --vendor --disable=gocyclo --disable=gotype

$(GOMETALINTER):
	@echo "---> Installing gometalinter"
	go get -u -v github.com/alecthomas/gometalinter
	gometalinter -i

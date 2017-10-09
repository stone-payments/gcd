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

build: $(RELEASE_PATH)

build-image:
	@echo "---> Building the project using Dockerfile"
	@docker build -t guiferpa/$(PROJECT):$(VERSION) .
	
$(RELEASE_PATH): clear
	@echo "---> Building the project"
	@mkdir -p $(RELEASE_DIR)
	@CGO_ENABLED=0 GOOS=$(PLATFORM) GOARCH=$(ARCH) go build -o $(RELEASE_PATH) $(PATH_MAIN_PACKAGE)

clear:	test
	@echo "---> Cleaning up directory"
	@rm -rf $(RELEASE_DIR)

test:	lint
	@echo "---> Testing"
	@rm -rf *.out
	@echo "mode: count" > coverage-all.out
	@for pkg in $(PKGS); do \
		@go test -v -cover -coverprofile=coverage.out -covermode=count $$pkg -timeout 10ms && \
		tail -n +2 coverage.out >> ./coverage-all.out; \
	done

lint: $(GOMETALINTER)
	@echo "---> Running lint"
	@gometalinter ../$(PROJECT)/... --vendor --disable=gocyclo --disable=gotype

$(GOMETALINTER):
	@echo "---> Installing gometalinter"
	go get -u -v github.com/alecthomas/gometalinter
	gometalinter -i

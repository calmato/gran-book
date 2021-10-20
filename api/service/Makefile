.PHONY: install update gazelle mockgen protoc test

GOLANGCI_VERSION := $(abspath $(CURDIR)/bin)

LINT_PACKAGES := $(shell go list $(CURDIR)/... | grep -v -e "mock" -v -e "proto")
TEST_PACKAGES := $(shell go list $(CURDIR)/internal/... $(CURDIR)/pkg/... | grep -v -e "pkg/test")
EXTERNAL_PACKAGE=./third_party

setup: install
	wget -O - -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v${GOLANGCI_VERSION}

install: install-third-party
	go install github.com/bazelbuild/bazelisk@v1.10.1
	go install github.com/golang/mock/mockgen@v1.6.0
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@1.11.0

install-third-party:
	rm -rf ${EXTERNAL_PACKAGE}/**/*
	mkdir -p ${EXTERNAL_PACKAGE}/github.com/envoyproxy
	git clone https://github.com/envoyproxy/protoc-gen-validate.git ${EXTERNAL_PACKAGE}/github.com/envoyproxy/protoc-gen-validate

# update:
# 	go mod tidy
# 	$(MAKE) gazelle

# gazelle: WORKSPACE
# 	bazelisk run //:gazelle

# WORKSPACE: go.mod
# 	bazelisk run //:gazelle -- update-repos -prune -from_file ./go.mod

# checksum:
# 	mkdir -p ./tmp; rm -f ./tmp/checksum-file
# 	wget -O ./tmp/checksum-file ${DOWNLOAD_URL}
# 	shasum -a 256 ./tmp/checksum-file
# 	rm -f ./tmp/checksum-file

mockgen:
	rm -rf $(CURDIR)/mock
	go generate ./...

protoc:
	find $(CURDIR)/proto/**/*.go -type f | xargs rm -rf
	protoc \
		-I ./proto \
		-I ./third_party \
		--go_out ./proto \
		--go_opt paths=source_relative \
		--go-grpc_out ./proto \
		--go-grpc_opt paths=source_relative \
		--validate_out lang=go:./proto \
		--validate_opt paths=source_relative \
		./proto/**/*.proto

fmt:
	! gofmt -d -s . | grep '^'

vet:
	go vet $(LINT_PACKAGES)

lint:
	$(BIN)/golangci-lint run -c .golangci.yaml $(LINT_PACKAGES)

test:
	go test -v -cover -coverprofile=coverage.txt -covermode=atomic $(TEST_PACKAGES)

build:
	go build -o $(CURDIR)/app $(CURDIR)/cmd/${SERVICE}/main.go

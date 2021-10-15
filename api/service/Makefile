.PHONY: install gazelle gazelle-update mockgen test

GOLANGCI_VERSION := $(abspath ./bin)

LINT_PACKAGES := $(shell go list ./... | grep -v -e "mock" -v -e "proto")
TEST_PACKAGES := $(shell go list ./internal/... ./pkg/... | grep -v -e "pkg/test")

setup: install
	wget -O - -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v${GOLANGCI_VERSION}

install:
	go install github.com/bazelbuild/bazelisk@v1.10.1
	go install github.com/golang/mock/mockgen@v1.6.0

gazelle:
	bazelisk run //:gazelle

gazelle-update:
	bazelisk run //:gazelle -- update-repos -from_file ./go.mod

mockgen:
	rm -rf ./mock
	go generate ./...

fmt:
	! gofmt -d -s . | grep '^'

vet:
	go vet $(LINT_PACKAGES)

lint:
	$(BIN)/golangci-lint run -c .golangci.yaml $(LINT_PACKAGES)

test:
	go test -v -cover -coverprofile=coverage.txt -covermode=atomic $(TEST_PACKAGES)

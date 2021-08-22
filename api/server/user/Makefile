GOLANGCI_VERSION := 1.33.0

BIN := $(abspath ./bin)
PACKAGES := $(shell go list ./... | grep -v -e "mock" -v -e "proto")
TEST_PACKAGES := $(shell go list ./internal/... ./pkg/... | grep -v -e "domain" -v -e "infrastructure/storage" -v -e "pkg/firebase" -v -e "pkg/metadata" -v -e "pkg/test")

.PHONY: setup run build start fmt vet lint test

setup:
	wget -O - -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v${GOLANGCI_VERSION}
	go install github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen

run:
	go run ./cmd/main.go

build:
	go build -o app ./cmd/main.go

start:
	./app

fmt:
	! gofmt -d -s . | grep '^'

vet:
	go vet ./... $(PACKAGES)

lint:
	$(BIN)/golangci-lint run -c .golangci.yaml ./...

test:
	go test -v -cover -coverprofile=./tmp/coverage.txt -covermode=atomic $(TEST_PACKAGES)

mockgen:
	$(BIN)/build-mocks.sh

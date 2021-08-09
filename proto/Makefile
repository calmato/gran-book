PROTOC_GEN_GRPC_PATH=$(shell which grpc_tools_node_protoc_plugin)
PROTOC_GEN_VALIDATE_VERSION=0.6.1

PACKAGE_PATH=./pkg
ARCHIVE_PATH=./archive

ADMIN_GATEWAY_PATH=./../api/gateway/admin
NATIVE_GATEWAY_PATH=./../api/gateway/native
USER_SERVER_PATH=./../api/server/user
BOOK_SERVER_PATH=./../api/server/book
INFORMATION_SERVER_PATH=./../api/server/information

##################################################
# Main
##################################################
.PHONY: build

build: install clean-all generate-all

##################################################
# For Configuration
##################################################
.PHONY: install

install:
	go get \
		github.com/envoyproxy/protoc-gen-validate@v${PROTOC_GEN_VALIDATE_VERSION}
	go install \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc
	mkdir -p ${PACKAGE_PATH}/github.com/envoyproxy
	if [ -L ${PACKAGE_PATH}/github.com/envoyproxy/protoc-gen-validate ]; then \
		unlink ${PACKAGE_PATH}/github.com/envoyproxy/protoc-gen-validate; \
	fi
	ln -s \
		${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v${PROTOC_GEN_VALIDATE_VERSION} \
		${PACKAGE_PATH}/github.com/envoyproxy/protoc-gen-validate

##################################################
# For Protocol Buffer
##################################################
.PHONY: format clean generate

format:
	prototool format -d

format-fix:
	prototool format -w

clean-all:
	rm -rf ${ARCHIVE_PATH}/proto
	$(MAKE) clean SERVICE_PATH=${NATIVE_GATEWAY_PATH}
	$(MAKE) clean SERVICE_PATH=${USER_SERVER_PATH}
	$(MAKE) clean SERVICE_PATH=${BOOK_SERVER_PATH}
	$(MAKE) clean SERVICE_PATH=${INFORMATION_SERVER_PATH}

clean:
	rm -rf ${SERVICE_PATH}/proto/*.pb.go
	rm -rf ${SERVICE_PATH}/proto/*.pb.validate.go

generate-all:
	$(MAKE) protoc ARCHIVE_PATH=${ARCHIVE_PATH}
	$(MAKE) generate SERVICE_PATH=${NATIVE_GATEWAY_PATH}
	$(MAKE) generate SERVICE_PATH=${USER_SERVER_PATH}
	$(MAKE) generate SERVICE_PATH=${BOOK_SERVER_PATH}
	$(MAKE) generate SERVICE_PATH=${INFORMATION_SERVER_PATH}

generate:
	if [ ! -d ${ARCHIVE_PATH}/proto ]; then \
		$(MAKE) protoc ARCHIVE_PATH=${ARCHIVE_PATH}; \
	fi
	cp -r ${ARCHIVE_PATH}/proto ${SERVICE_PATH}

protoc:
	protoc \
		-I . \
		-I ./pkg \
		--go_out ${ARCHIVE_PATH} \
		--go_opt paths=source_relative \
		--go-grpc_out ${ARCHIVE_PATH} \
		--go-grpc_opt paths=source_relative \
		--validate_out lang=go:${ARCHIVE_PATH} \
		--validate_opt paths=source_relative \
		./proto/*.proto

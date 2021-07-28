PROTOC_GEN_GRPC_PATH=$(shell which grpc_tools_node_protoc_plugin)
PROTOC_GEN_TS_PATH=$(shell which protoc-gen-ts)

PROTOC_GEN_VALIDATE_VERSION=0.6.1

ADMIN_GATEWAY_PATH=./../api/gateway/admin/src
NATIVE_GATEWAY_PATH=./../api/gateway/native/src
USER_SERVER_PATH=./../api/server/user
BOOK_SERVER_PATH=./../api/server/book
INFORMATION_SERVER_PATH=./../api/server/information

##################################################
# Main
##################################################
protoc: install format-fix clean generate

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
	mkdir -p ./pkg/github.com/envoyproxy
	if [ -L ./pkg/github.com/envoyproxy/protoc-gen-validate ]; then \
		unlink ./pkg/github.com/envoyproxy/protoc-gen-validate; \
	fi
	ln -s \
		${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v${PROTOC_GEN_VALIDATE_VERSION} \
		./pkg/github.com/envoyproxy/protoc-gen-validate

##################################################
# For Protocol Buffer
##################################################
.PHONY: format clean generate

format:
	prototool format -d

format-fix:
	prototool format -w

clean:
	$(MAKE) clean-gateway GATEWAY_PATH=${ADMIN_GATEWAY_PATH}
	$(MAKE) clean-gateway GATEWAY_PATH=${NATIVE_GATEWAY_PATH}
	$(MAKE) clean-server SERVER_PATH=${USER_SERVER_PATH}
	$(MAKE) clean-server SERVER_PATH=${BOOK_SERVER_PATH}
	$(MAKE) clean-server SERVER_PATH=${INFORMATION_SERVER_PATH}

clean-gateway:
	rm -rf ${GATEWAY_PATH}/proto/*_pb.js
	rm -rf ${GATEWAY_PATH}/proto/*_pb.d.ts

clean-server:
	rm -rf ${SERVER_PATH}/proto/*.pb.go
	rm -rf ${SERVER_PATH}/proto/*.pb.validate.go

generate:
	$(MAKE) generate-gateway GATEWAY_PATH=${ADMIN_GATEWAY_PATH}
	$(MAKE) generate-gateway GATEWAY_PATH=${NATIVE_GATEWAY_PATH}
	$(MAKE) generate-server SERVER_PATH=${USER_SERVER_PATH}
	$(MAKE) generate-server SERVER_PATH=${BOOK_SERVER_PATH}
	$(MAKE) generate-server SERVER_PATH=${INFORMATION_SERVER_PATH}

generate-gateway:
	protoc \
		-I . \
		-I ./pkg \
		--plugin=protoc-gen-ts=${PROTOC_GEN_TS_PATH} \
		--plugin=protoc-gen-grpc=${PROTOC_GEN_GRPC_PATH} \
		--js_out=import_style=commonjs,binary:${GATEWAY_PATH} \
		--grpc_out=grpc_js:${GATEWAY_PATH} \
		--ts_out=grpc_js:${GATEWAY_PATH} \
		./proto/*.proto

generate-server:
	protoc \
		-I . \
		-I ./pkg \
		--go_out ${SERVER_PATH} \
		--go_opt paths=source_relative \
		--go-grpc_out ${SERVER_PATH} \
		--go-grpc_opt paths=source_relative \
		--validate_out lang=go:${SERVER_PATH} \
		--validate_opt paths=source_relative \
		./proto/common.proto \
		./proto/*_api*.proto \
		./proto/*_service.proto


PROTOC_GEN_GRPC_PATH=$(shell which grpc_tools_node_protoc_plugin)
PROTOC_GEN_TS_PATH=$(shell which protoc-gen-ts)

ADMIN_GATEWAY_PATH='./../api/gateway/admin/src'
NATIVE_GATEWAY_PATH='./../api/gateway/native/src'
USER_API_PATH='./../api/server/user'
BOOK_API_PATH='./../api/server/book'
INFORMATION_API_PATH='./../api/server/information'

##################################################
# For Configuration
##################################################
.PHONY: install

install:
	go install \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

##################################################
# For Protocol Buffer
##################################################
.PHONY: generate format

format:
	prototool format -d

format-fix:
	prototool format -w

generate:
	$(MAKE) install
	$(MAKE) protoc-gateway GATEWAY_PATH=${ADMIN_GATEWAY_PATH}
	$(MAKE) protoc-gateway GATEWAY_PATH=${NATIVE_GATEWAY_PATH}
	$(MAKE) protoc-service SERVICE_PATH=${USER_API_PATH}
	$(MAKE) protoc-service SERVICE_PATH=${BOOK_API_PATH}
	$(MAKE) protoc-service SERVICE_PATH=${INFORMATION_API_PATH}

protoc-gateway:
	protoc -I . \
		--plugin=protoc-gen-ts=${PROTOC_GEN_TS_PATH} \
		--plugin=protoc-gen-grpc=${PROTOC_GEN_GRPC_PATH} \
		--js_out=import_style=commonjs,binary:${GATEWAY_PATH} \
		--grpc_out=grpc_js:${GATEWAY_PATH} \
		--ts_out=grpc_js:${GATEWAY_PATH} \
		./proto/*.proto

protoc-service:
	protoc -I . \
		--go_out ${SERVICE_PATH} \
		--go_opt paths=source_relative \
		--go-grpc_out ${SERVICE_PATH} \
		--go-grpc_opt paths=source_relative \
		./proto/*api*.proto

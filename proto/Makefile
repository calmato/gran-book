PROTOC_GEN_GRPC_PATH=$(shell which grpc_tools_node_protoc_plugin)
PROTOC_GEN_TS_PATH=$(shell which protoc-gen-ts)

ADMIN_GATEWAY_PATH='./../api/gateway/admin/src'
USER_API_PATH='./../api/server/user'

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
.PHONY: generate

generate:
	$(MAKE) install
	$(MAKE) generate-admin-gateway
	$(MAKE) generate-user-api

generate-admin-gateway:
	protoc -I . \
		--plugin=protoc-gen-ts=${PROTOC_GEN_TS_PATH} \
		--plugin=protoc-gen-grpc=${PROTOC_GEN_GRPC_PATH} \
		--js_out=import_style=commonjs,binary:${ADMIN_GATEWAY_PATH} \
		--grpc_out=grpc_js:${ADMIN_GATEWAY_PATH} \
		--ts_out=grpc_js:${ADMIN_GATEWAY_PATH} \
		./proto/*.proto

generate-user-api:
	protoc -I . \
		--go_out ${USER_API_PATH} \
		--go_opt paths=source_relative \
		--go-grpc_out ${USER_API_PATH} \
		--go-grpc_opt paths=source_relative \
		./proto/user_apiv1.proto

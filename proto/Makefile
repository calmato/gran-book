PROTOC_GEN_GRPC_PATH=$(shell which grpc_tools_node_protoc_plugin)
PROTOC_GEN_TS_PATH=$(shell which protoc-gen-ts)

ADMIN_GATEWAY_PATH='./../api/gateway/admin/src'
NATIVE_GATEWAY_PATH='./../api/gateway/native/src'
USER_API_PATH='./../api/server/user'
BOOK_API_PATH='./../api/server/book'

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
	$(MAKE) generate-admin-gateway
	$(MAKE) generate-native-gateway
	$(MAKE) generate-user-api
	$(MAKE) generate-book-api

generate-admin-gateway:
	protoc -I . \
		--plugin=protoc-gen-ts=${PROTOC_GEN_TS_PATH} \
		--plugin=protoc-gen-grpc=${PROTOC_GEN_GRPC_PATH} \
		--js_out=import_style=commonjs,binary:${ADMIN_GATEWAY_PATH} \
		--grpc_out=grpc_js:${ADMIN_GATEWAY_PATH} \
		--ts_out=grpc_js:${ADMIN_GATEWAY_PATH} \
		./proto/*.proto

generate-native-gateway:
	protoc -I . \
		--plugin=protoc-gen-ts=${PROTOC_GEN_TS_PATH} \
		--plugin=protoc-gen-grpc=${PROTOC_GEN_GRPC_PATH} \
		--js_out=import_style=commonjs,binary:${NATIVE_GATEWAY_PATH} \
		--grpc_out=grpc_js:${NATIVE_GATEWAY_PATH} \
		--ts_out=grpc_js:${NATIVE_GATEWAY_PATH} \
		./proto/*.proto

generate-user-api:
	protoc -I . \
		--go_out ${USER_API_PATH} \
		--go_opt paths=source_relative \
		--go-grpc_out ${USER_API_PATH} \
		--go-grpc_opt paths=source_relative \
		./proto/user_apiv1.proto

generate-book-api:
	protoc -I . \
		--go_out ${BOOK_API_PATH} \
		--go_opt paths=source_relative \
		--go-grpc_out ${BOOK_API_PATH} \
		--go-grpc_opt paths=source_relative \
		./proto/book_apiv1.proto

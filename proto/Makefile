PROTOC_GEN_GRPC_PATH=$(shell which grpc_tools_node_protoc_plugin)
PROTOC_GEN_TS_PATH=$(shell which protoc-gen-ts)
PROTOC_GEN_VALIDATE_VERSION=0.6.1

PACKAGE_PATH=./pkg
ARCHIVE_PATH=./archive

ADMIN_CLIENT_PATH=./../web/admin
NATIVE_CLIENT_PATH=./../native

ADMIN_GATEWAY_PATH=./../api/gateway/admin
NATIVE_GATEWAY_PATH=./../api/gateway/native
USER_SERVER_PATH=./../api/server/user
BOOK_SERVER_PATH=./../api/server/book
INFORMATION_SERVER_PATH=./../api/server/information

##################################################
# Main
##################################################
.PHONY: build

build: install clean generate

# ------------------------------------------------
# For Protocol Buffer - generate
# ------------------------------------------------
generate: protoc
	$(MAKE) generate-client OUTPUT_PATH=${NATIVE_CLIENT_PATH} CLIENT=native
	$(MAKE) generate-client OUTPUT_PATH=${ADMIN_CLIENT_PATH} CLIENT=admin
	$(MAKE) generate-gateway OUTPUT_PATH=${NATIVE_GATEWAY_PATH} CLIENT=native
	$(MAKE) generate-gateway OUTPUT_PATH=${ADMIN_GATEWAY_PATH} CLIENT=admin
	$(MAKE) generate-service OUTPUT_PATH=${USER_SERVER_PATH}
	$(MAKE) generate-service OUTPUT_PATH=${BOOK_SERVER_PATH}
	$(MAKE) generate-service OUTPUT_PATH=${INFORMATION_SERVER_PATH}

generate-client:
	cp -r ${ARCHIVE_PATH}/common/*_pb.d.ts ${OUTPUT_PATH}/app/types/api
	cp -r ${ARCHIVE_PATH}/gateway/${CLIENT}/*_pb.d.ts ${OUTPUT_PATH}/app/types/api

generate-gateway:
	cp -r ${ARCHIVE_PATH}/common/*.pb.go ${OUTPUT_PATH}/proto
	cp -r ${ARCHIVE_PATH}/service/*.pb.go ${OUTPUT_PATH}/proto

generate-service:
	cp -r ${ARCHIVE_PATH}/common/*.pb.go ${OUTPUT_PATH}/proto
	cp -r ${ARCHIVE_PATH}/common/*.pb.validate.go ${OUTPUT_PATH}/proto
	cp -r ${ARCHIVE_PATH}/service/*.pb.go ${OUTPUT_PATH}/proto
	cp -r ${ARCHIVE_PATH}/service/*.pb.validate.go ${OUTPUT_PATH}/proto

# ------------------------------------------------
# For Protocol Buffer - clean
# ------------------------------------------------
clean: clean-archive
	$(MAKE) clean-client OUTPUT=${NATIVE_CLIENT_PATH}
	$(MAKE) clean-client OUTPUT=${ADMIN_CLIENT_PATH}
	$(MAKE) clean-gateway OUTPUT_PATH=${NATIVE_GATEWAY_PATH}
	$(MAKE) clean-gateway OUTPUT_PATH=${ADMIN_GATEWAY_PATH}
	$(MAKE) clean-service OUTPUT_PATH=${USER_SERVER_PATH}
	$(MAKE) clean-service OUTPUT_PATH=${BOOK_SERVER_PATH}
	$(MAKE) clean-service OUTPUT_PATH=${INFORMATION_SERVER_PATH}

clean-archive:
	rm -rf ${ARCHIVE_PATH}/**/*_pb.d.ts
	rm -rf ${ARCHIVE_PATH}/**/*.pb.go
	rm -rf ${ARCHIVE_PATH}/**/*.pb.validate.go

clean-client:
	rm -rf ${OUTPUT_PATH}/app/types/api/*_pb.d.ts

clean-gateway:
	rm -rf ${OUTPUT_PATH}/proto/*.pb.go

clean-service:
	rm -rf ${OUTPUT_PATH}/proto/*.pb.go
	rm -rf ${OUTPUT_PATH}/proto/*.pb.validate.go

##################################################
# For Configuration
##################################################
.PHONY: install

install:
	npm install -g grpc grpc-tools grpc_tools_node_protoc_ts
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
.PHONY: format clean generate protoc

format:
	prototool format -d

format-fix:
	prototool format -w

# ------------------------------------------------
# For Protocol Buffer - protoc
# ------------------------------------------------
protoc: protoc-gateway protoc-service

protoc-gateway:
	protoc \
		-I ./pkg \
		-I ./proto \
		--ts_out=grpc_js:${ARCHIVE_PATH} \
		./proto/common/*.proto \
		./proto/gateway/**/*.proto

protoc-service:
	protoc \
		-I ./pkg \
		-I ./proto \
		--go_out ${ARCHIVE_PATH} \
		--go_opt paths=source_relative \
		--go-grpc_out ${ARCHIVE_PATH} \
		--go-grpc_opt paths=source_relative \
		--validate_out lang=go:${ARCHIVE_PATH} \
		--validate_opt paths=source_relative \
		./proto/common/*.proto \
		./proto/service/*.proto

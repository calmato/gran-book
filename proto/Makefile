PROTOC_GEN_GRPC_PATH=$(shell which grpc_tools_node_protoc_plugin)
PROTOC_GEN_TS_PATH=$(shell which protoc-gen-ts)
PROTOC_GEN_VALIDATE_VERSION=0.6.1

PACKAGE_PATH=./pkg
ARCHIVE_PATH=./archive

ADMIN_CLIENT_PATH=./../web/admin
NATIVE_CLIENT_PATH=./../native

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

generate-client:
	mkdir -p ${OUTPUT_PATH}/app/types/api
	cp -r ${ARCHIVE_PATH}/common/*_pb.d.ts ${OUTPUT_PATH}/app/types/api
	cp -r ${ARCHIVE_PATH}/gateway/${CLIENT}/*_pb.d.ts ${OUTPUT_PATH}/app/types/api

# ------------------------------------------------
# For Protocol Buffer - clean
# ------------------------------------------------
clean: clean-archive
	$(MAKE) clean-client OUTPUT=${NATIVE_CLIENT_PATH}
	$(MAKE) clean-client OUTPUT=${ADMIN_CLIENT_PATH}

clean-archive:
	find ${ARCHIVE_PATH} -type f | grep -v -e '.keep' | xargs rm -rf

clean-client:
	find ${OUTPUT_PATH}/app/types/api -type f | grep -v -e '.keep' | xargs rm -rf

##################################################
# For Configuration
##################################################
.PHONY: install

install:
	npm install -g @grpc/grpc-js grpc-tools grpc_tools_node_protoc_ts

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
protoc: protoc-gateway

protoc-gateway:
	protoc \
		-I ./pkg \
		-I ./proto \
		--ts_out=grpc_js:${ARCHIVE_PATH} \
		./proto/common/*.proto \
		./proto/gateway/**/*.proto

##################################################
# For Configuration
##################################################
.PHONY: install

install:
	go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

##################################################
# For Protocol Buffer
##################################################
.PHONY: generate generate-gateway

generate:
	$(MAKE) generate-gateway

generate-gateway:
	protoc -I . \
		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out ./../infra/gateway/proto \
		--go-grpc_out ./../infra/gateway/proto \
		--grpc-gateway_out ./../infra/gateway/proto \
		./api/hello_world.proto

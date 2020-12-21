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
.PHONY: generate

generate:
	$(MAKE) generate-gateway
	$(MAKE) generate-swagger
	$(MAKE) generate-helloworld

generate-gateway:
	protoc -I . \
		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out ./../infra/gateway \
		--grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
		./proto/hello_world.proto

generate-swagger:
	protoc -I . \
		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--openapiv2_out ./../docs/12_backend/11_swagger \
		--openapiv2_opt logtostderr=true \
		./proto/hello_world.proto

generate-helloworld:
	protoc -I . \
		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out ./../api/hello \
		--go_opt paths=source_relative \
		--go-grpc_out ./../api/hello \
		--go-grpc_opt paths=source_relative \
		./proto/hello_world.proto

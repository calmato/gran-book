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
	$(MAKE) generate-user-api

generate-gateway:
	protoc -I . \
		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out ./../infra/gateway \
		--go_opt paths=source_relative \
		--go-grpc_out ./../infra/gateway \
		--go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./../infra/gateway \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
		./proto/*.proto

generate-swagger:
	protoc -I . \
		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--openapiv2_out ./../docs/12_backend/11_swagger \
		--openapiv2_opt allow_merge=true \
		--openapiv2_opt merge_file_name=apiv1 \
		--openapiv2_opt logtostderr=true \
		./proto/*.proto

generate-user-api:
	protoc -I . \
		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out ./../api/user \
		--go_opt paths=source_relative \
		--go-grpc_out ./../api/user \
		--go-grpc_opt paths=source_relative \
		./proto/user_apiv1.proto

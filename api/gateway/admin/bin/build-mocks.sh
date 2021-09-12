#!/bin/sh

MOCKGEN_PATH=$(which mockgen)
if [ ! -f "${MOCKGEN_PATH}" ]; then
  echo "Not found mockgen. Please install this tool."
  echo "-> $ go install github.com/golang/mock/mockgen@latest"
  exit 1
fi

#############################
# Function
#############################
build_mock() {
  path=${1##$(PWD)/proto/service/}

  dirname=${path%%/*}
  filename=${path##*/}

  mockgen -source ./proto/service/${path} -destination mock/${dirname}/${filename}.go
}

#############################
# Main
#############################
rm -rf ./mock

paths=$(find $(PWD)/proto/service/**/*_service_grpc.pb.go -type f)
for path in "${paths}"; do
  build_mock ${path}
done

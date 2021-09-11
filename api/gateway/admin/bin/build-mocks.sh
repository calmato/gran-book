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
  path=${1##.\/proto\/}

  dirname=${path%%/*}
  filename=${path##*/}

  mockgen -package mock -source ./proto/${path} -destination mock/${dirname}/${filename}.go
}

#############################
# Main
#############################
rm -rf mock/*.go

paths=$(find ./proto/**/*_service.pb.go)
for path in ${paths}; do
  build_mock ${path}
done

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
  filename=$(echo ${1} | awk -F '_' '{ print $1 }')

  mockgen -package mock -source ${1} -destination mock/${filename}.go
}

#############################
# Main
#############################
rm -rf mock/*.go

paths=$(find ./proto/**/*_service.pb.go)
for file in ${paths}; do
  build_mock ${file}
done

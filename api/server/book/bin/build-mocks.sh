#!/bin/sh

MOCKGEN_PATH=$(which mockgen)
if [ ! -f "${MOCKGEN_PATH}" ]; then
  echo "Not found mockgen. Please install this tool."
  echo "-> $ go install github.com/golang/mock/mockgen@latest"
  exit 1
fi

#############################
# Variables
#############################
ignore_lists='test|exception|validator.go'

#############################
# Function
#############################
build_mock() {
  target=$1

  mockgen -source internal/${target} -destination mock/${target}
}

build_file() {
  dir_name=$1
  file_name=$2

  paths=$(find internal/${dir_name} -name ${file_name} \
    | grep -vE ${ignore_lists} \
    | awk -v var=${dir_name} '{ print substr($0, index($0, var)) }' \
    | sed -e "s/\/${file_name}$//")

  for path in ${paths}; do
    mkdir -p mock/${path}
    build_mock "${path}/${file_name}"
  done
}

build_package() {
  dir_name="$1/$2"
  mkdir -p mock/${dir_name}

  paths=$(find internal/${dir_name} -name '*.go' \
    | grep -vE ${ignore_lists} \
    | awk -v var=${dir_name} '{ print substr($0, index($0, var)) }')

  for path in ${paths}; do
    build_mock ${path}
  done
}

#############################
# Target
#############################
# --- Domain ---
build_file 'domain' 'repository.go'
build_file 'domain' 'service.go'
build_file 'domain' 'uploader.go'
build_file 'domain' 'messaging.go'
build_file 'domain' 'validation.go'

# --- Application ---
build_package 'application' 'validation'

# --- Infrastructure ---
# build_package 'infrastructure' 'api'

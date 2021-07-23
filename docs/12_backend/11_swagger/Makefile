.PHONY: build lint generate validate

NATIVE_INPUT_DIR := './native'
NATIVE_OUTPUT_DIR := './../../../tmp/data/swagger/native'
ADMIN_INPUT_DIR := './admin'
ADMIN_OUTPUT_DIR := './../../../tmp/data/swagger/admin'

build:
	$(MAKE) generate INPUT_DIR=${NATIVE_INPUT_DIR} OUTPUT_DIR=${NATIVE_OUTPUT_DIR}
	$(MAKE) generate INPUT_DIR=${ADMIN_INPUT_DIR} OUTPUT_DIR=${ADMIN_OUTPUT_DIR}

generate:
	openapi-generator-cli generate \
		-g html2 \
		-g openapi-yaml \
		-i ${INPUT_DIR}/openapi.yaml \
		-o ${OUTPUT_DIR}

lint:
	$(MAKE) validate INPUT_DIR=${NATIVE_INPUT_DIR}
	$(MAKE) validate INPUT_DIR=${ADMIN_INPUT_DIR}

validate:
	openapi-generator-cli validate \
		-i ${INPUT_DIR}/openapi.yaml \
		--recommend

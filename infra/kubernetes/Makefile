CURRENT := $(shell date "+%Y%m%d-%H%M%S")

#############################
# Common
#############################
.PHONY: get-external-ip

get-external-ip:
	kubectl get service ${APP_NAME} -o jsonpath='{.status.loadBalancer.ingress[0].ip}'

#############################
# Setup
#############################
.PHONY: generate-parivate-key generate-manifest generate-manifest-all

generate-private-key:
	mkdir -p $(PWD)/tmp/${EXTERNAL_IP}
	openssl req \
		-x509 -nodes -newkey rsa:2048 -days 365 \
		-keyout $(PWD)/tmp/${EXTERNAL_IP}/privkey.pem \
		-out $(PWD)/tmp/${EXTERNAL_IP}/cert.pem \
		-subj "/CN=${EXTERNAL_IP}"

generate-manifest:
	@if [[ ! -e "${MANIFEST}" ]]; then \
		sed 's/<PROJECT-ID>/$(PROJECT_ID)/' $(MANIFEST).temp > $(MANIFEST); \
	fi

generate-manifest-all:
	@if [[ -z "${PROJECT_ID}" ]] || [[ -z "${ENV}" ]]; then \
		echo "PROJECT_ID or ENV is null."; \
		exit 1; \
	fi
	@for i in $$(find ./${ENV} | grep '.temp'); do \
		$(MAKE) generate-manifest PROJECT_ID=${PROJECT_ID} MANIFEST=$${i%%.temp}; \
	done

#############################
# Update
#############################
.PHONY: update-manifest update-manifest-all

update-manifest:
	kubectl patch ${RESOURCE} ${APP_NAME} -p "{\"spec\": {\"template\": {\"metadata\": {\"labels\": {\"date\": \"${CURRENT}\" }}}}}"

update-manifest-all:
	$(MAKE) update-manifest RESOURCE=deploy APP_NAME=user-api

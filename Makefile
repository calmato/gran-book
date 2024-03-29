##################################################
# Container Commands - Run All
##################################################
.PHONY: setup build install start stop down remove logs

setup:
	if [ ! -f $(PWD)/.env ]; then \
		cp $(PWD)/.env.temp $(PWD)/.env; \
	fi
	$(MAKE) build
	$(MAKE) install
	$(MAKE) proto
	$(MAKE) swagger

build:
	docker-compose build --parallel

install:
	docker-compose run --rm admin yarn
	docker-compose run --rm native yarn
	docker-compose run --rm swagger_generator yarn

start:
	$(PWD)/bin/get-local-ip-addr.sh
	docker-compose up --remove-orphans

stop:
	docker-compose stop

down:
	docker-compose down

remove:
	docker-compose down --rmi all --volumes --remove-orphans

logs:
	docker-compose logs

##################################################
# Container Commands - Run Container Group
##################################################
.PHONY: start-native start-admin start-api start-swagger

start-native:
	$(PWD)/bin/get-local-ip-addr.sh
	docker-compose up native

start-admin:
	docker-compose up admin

start-api:
	$(MAKE) proto
	docker-compose up native_gateway admin_gateway user_api book_api mysql mysql_test

start-swagger:
	docker-compose up swagger_native swagger_admin swagger_generator

##################################################
# Container Commands - Single
##################################################
.PHONY: proto swagger migrate

proto:
	docker-compose run --rm proto make build

swagger:
	docker-compose run --rm swagger_generator yarn generate

migrate:
	docker-compose start mysql mysql_test
	docker-compose exec mysql bash -c "mysql -u root -p${MYSQL_ROOT_PASSWORD} < /docker-entrypoint-initdb.d/*.sql"
	docker-compose exec mysql_test bash -c "mysql -u root -p${MYSQL_ROOT_PASSWORD} < /docker-entrypoint-initdb.d/*.sql"
	docker-compose stop mysql mysql_test

##################################################
# Container Commands - Terraform
##################################################
.PHONY: terraform-setup terraform-lint terraform-plan terraform-apply terraform-destroy

terraform-init:
	docker-compose run --rm terraform make init ENV=${ENV}

terraform-fmt:
	docker-compose run --rm terraform make fmt ENV=${ENV}

terraform-plan:
	docker-compose run --rm terraform make plan ENV=${ENV}

terraform-apply:
	docker-compose run --rm terraform make apply ENV=${ENV}

terraform-destroy:
	docker-compose run --rm terraform make destroy ENV=${ENV}

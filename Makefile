##################################################
# Container Commands - Run All
##################################################
.PHONY: setup build install start stop down remove logs

setup:
	cp $(PWD)/.env.temp $(PWD)/.env
	$(MAKE) build
	$(MAKE) install
	$(MAKE) migrate

build:
	docker-compose build --parallel

install:
	docker-compose run --rm admin yarn
	docker-compose run --rm native yarn

start:
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
.PHONY: start-native start-admin start-api

start-native:
	$(PWD)/bin/get-local-ip-addr.sh
	docker-compose up native

start-admin:
	docker-compose up admin

start-api:
	docker-compose up gateway user_api mysql swagger_editor

start-swagger:
	docker-compose up swagger swagger_editor

##################################################
# Container Commands - Single
##################################################
.PHONY: proto migrate

proto:
	docker-compose run --rm proto bash -c "make install && make generate"

migrate:
	docker-compose start mysql
	docker-compose exec mysql bash -c "mysql -u root -p${MYSQL_ROOT_PASSWORD} < /docker-entrypoint-initdb.d/*.sql"
	docker-compose stop mysql

##################################################
# Container Commands - Terraform
##################################################
.PHONY: terraform-setup terraform-lint terraform-plan terraform-apply terraform-destroy

terraform-setup:
	docker-compose run --rm terraform make init ENV=${ENV}

terraform-lint:
	docker-compose run --rm terraform make fmt ENV=${ENV}

terraform-plan:
	docker-compose run --rm terraform make plan ENV=${ENV}

terraform-apply:
	docker-compose run --rm terraform make apply ENV=${ENV}

terraform-destroy:
	docker-compose run --rm terraform make destroy ENV=${ENV}

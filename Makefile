##################################################
# Container Commands - Run containers
##################################################
.PHONY: setup build install start start-native start-admin start-api stop down remove logs proto migrate

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

start-native:
	$(PWD)/bin/get-local-ip-addr.sh
	docker-compose up native

start-admin:
	docker-compose up admin

start-api:
	docker-compose up gateway user_api mysql swagger_editor

start-swagger:
	docker-compose up swagger swagger_editor

stop:
	docker-compose stop

down:
	docker-compose down

remove:
	docker-compose down --rmi all --volumes --remove-orphans

logs:
	docker-compose logs

proto:
	docker-compose run --rm proto bash -c "make install && make generate"

migrate:
	docker-compose start mysql
	docker-compose exec mysql bash -c "mysql -u root -p${MYSQL_ROOT_PASSWORD} < /docker-entrypoint-initdb.d/*.sql"
	docker-compose stop mysql

##################################################
# Container Commands - Run all containers
##################################################
.PHONY: setup build install start stop remove log proto

setup:
	$(MAKE) build
	$(MAKE) install
	cp $(PWD)/.env.temp $(PWD)/.env
	yarn global add expo-cli

build:
	docker-compose build --parallel

install:
	docker-compose run --rm admin yarn
	docker-compose run --rm native yarn

start:
	docker-compose up

start-native:
	cd $(PWD)/native && yarn start

start-api:
	echo "WIP"

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
	docker-compose run --rm mysql bash -c "mysql -u root -p < /docker-entrypoint-initdb.d/01-create.sql"

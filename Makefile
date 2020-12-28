##################################################
# Container Commands - Run all containers
##################################################
.PHONY: setup build install start stop remove log proto

setup:
	cp $(PWD)/.env.temp $(PWD)/.env
	$(MAKE) build
	$(MAKE) install

build:
	docker-compose build

install:
	docker-compose run --rm admin yarn
	docker-compose run --rm native yarn

start:
	docker-compose up

stop:
	docker-compose stop

remove:
	docker-compose down --rmi all --volumes --remove-orphans

logs:
	docker-compose logs

proto:
	docker-compose run --rm proto bash -c "make install && make generate"

include .env
export

ifeq ($(DOCKER_BUILD_ENV), test)
	include ./deploy/test/.env.test
	export
	DOCKER_COMPOSE=./deploy/test/docker-compose.test.yml
endif

ifeq ($(DOCKER_BUILD_ENV), prod)
	# ..
endif

.PHONY: build
build:
	docker-compose -f $(DOCKER_COMPOSE) rm -sf
	docker-compose -f $(DOCKER_COMPOSE) down --remove-orphans
	docker-compose -f $(DOCKER_COMPOSE) build

.PHONY: up
up:
	docker-compose -f $(DOCKER_COMPOSE) up -d

.PHONY: stop
stop:
	docker-compose -f $(DOCKER_COMPOSE) stop

.PHONY: run-local-app
run-local-app:
	go run cmd/metadata-server/main.go

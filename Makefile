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

.PHONY: build-image
build-image:
	docker build --platform linux/amd64 -t eggsbit-metaserver -f build/docker/metaserver-app/Dockerfile .
	docker build --platform linux/amd64 -t eggsbit-nginx-metaserver -f build/docker/nginx-metaserver/Dockerfile .
	docker build --platform linux/amd64 -t eggsbit-nginx-static -f build/docker/nginx-static/Dockerfile .

.PHONY: tag-image
tag-image:
	docker tag eggsbit-metaserver registry.gitlab.com/eggsbit/eggsbit-metaserver/metaserver:$(APP_IMAGE_TAG)
	docker tag eggsbit-nginx-metaserver registry.gitlab.com/eggsbit/eggsbit-metaserver/nginx-metaserver:$(NGINX_METASERVER_IMAGE_TAG)
	docker tag eggsbit-nginx-static registry.gitlab.com/eggsbit/eggsbit-metaserver/nginx-static:$(NGINX_STATIC_IMAGE_TAG)

.PHONY: push-image
push-image:
	docker push registry.gitlab.com/eggsbit/eggsbit-metaserver/metaserver:$(APP_IMAGE_TAG)
	docker push registry.gitlab.com/eggsbit/eggsbit-metaserver/nginx-metaserver:$(NGINX_METASERVER_IMAGE_TAG)
	docker push registry.gitlab.com/eggsbit/eggsbit-metaserver/nginx-static:$(NGINX_STATIC_IMAGE_TAG)

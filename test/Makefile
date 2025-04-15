DOCKER_COMPOSE := docker compose -f docker-compose.yml

DOCKER_RELAY_PARENT := relay-parent
DOCKER_RELAY_CLIENT := relay-client
DOCKER_PROMETHEUS := prometheus
DOCKER_GRAFANA := grafana


RUN_DOCKER_RELAY_PARENT := $(DOCKER_COMPOSE) up -d $(DOCKER_RELAY_PARENT)
RUN_DOCKER_RELAY_CLIENT := $(DOCKER_COMPOSE) up -d $(DOCKER_RELAY_CLIENT)
RUN_DOCKER_PROMETHEUS := $(DOCKER_COMPOSE) up -d $(DOCKER_PROMETHEUS)
RUN_DOCKER_GRAFANA := $(DOCKER_COMPOSE) up -d $(DOCKER_GRAFANA)

STOP := $(DOCKER_COMPOSE) down --volumes

.PHONY: run
run:
	$(RUN_DOCKER_RELAY_PARENT)
	$(RUN_DOCKER_RELAY_CLIENT)

.PHONY: dbg
dbg: run
	$(RUN_DOCKER_PROMETHEUS)
	$(RUN_DOCKER_GRAFANA)
.PHONY: stop
stop:
	$(STOP)

.PHONY: build-docker
build-docker: ## Builds a docker image with datastream relay binary
	cd .. && docker build -t datastream-relay:test -f ./Dockerfile . && cd test




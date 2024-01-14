CONFIG = config/config-docker.yml
CONFIG_LOCAL = config/config-local.yml
COMPOSE_ALL = ./docker-compose.yml
COMPOSE_LOCAL = ./docker-compose.local.yml

CONF_EXISTS=$(shell [ -e $(CONFIG_F) ] && echo 1 || echo 0 )


all:
			docker compose -f $(COMPOSE_ALL) up --build

local:
			@(docker compose -f $(COMPOSE_LOCAL) up -d --build)

config:
			@(echo "Creating configs for launch. Don't forget change sample credentials.")
			@(cp ./config/config-sample.yml $(CONFIG))
			@(cp ./config/config-sample.yml $(CONFIG_LOCAL))

swagger:
			@swag fmt -g=./internal/app/app.go
			@swag init -g=./internal/app/app.go --output=./docs/swagger


.PHONY: all local config swagger

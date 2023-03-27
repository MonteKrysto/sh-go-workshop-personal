#!make

NETWORKS="$(shell docker network ls)"
POSTGRES_USER="pguser"
DONE=[ done "\xE2\x9C\x94" ]
CONTAINERS=workshop-api workshop-db
COMPOSER=docker-compose

# Initial setup of building the containers, the databases and seed data
setup: build up

# Build all the containers
build:
	@echo [ building and starting containers... ]
	$(COMPOSER) build --no-cache $(CONTAINERS)
	@echo $(DONE)

# Bring all the containers up
up:
	@echo [ starting containers... ]
	$(COMPOSER) up $(CONTAINERS)

# Bring all the containers to a stop
stop:
	@echo [ stopping all containers... ]
	$(COMPOSER) stop
	@echo $(DONE)

# Tear all containers down
down:
	@echo [ downing all containers... ]
	$(COMPOSER) down
	@echo $(DONE)

build-app:
	go build -o bin/main cmd/main.go

help:
	@grep '^.PHONY: .* #' Makefile | sed 's/\.PHONY: \(.*\) # \(.*\)/\1 \2/'


.PHONY: help                    # Generate list of targets with descriptions
.PHONY: setup                   # Initial setup of building the containers, the databases and seed data
.PHONY: build                   # Build all the containers
.PHONY: up                      # Start all the containers
.PHONY: stop                    # Stop all the containers
.PHONY: down                    # Down all the containers
.PHONY: build-app               # Build the app

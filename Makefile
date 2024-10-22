#!make
-include .env
export

DUMB_DIR= ./migrations/dump
DUMB_FILE := $(wildcard ${DUMB_DIR}/*.sql)

CURRENT_DIR = $(shell pwd)

PG_CONTAINER_NAME = test-lamoda-postgres

build:
	CGO_ENABLED=0 go build \
	-o output ./cmd/main.go
.PHONY: build

run:	
	go run ./cmd/main.go
.PHONY: run

@run:	
	docker compose up -d --force-recreate --build
.PHONY: @run

@up:
	docker-compose --env-file .env up
.PHONY: @up

@create_dump:
	docker-compose exec postgres pg_dump -U ${POSTGRES_USER} ${POSTGRES_DB} > ${DUMB_DIR}/pgdump-$$(date +%Y_%m_%d_%H_%M).sql
.PHONY: @create_dump

@restore_dump:
	docker exec -i $(PG_CONTAINER_NAME) /bin/bash -c "PGPASSWORD=${POSTGRES_PASSWORD} psql --username ${POSTGRES_USER} ${POSTGRES_DB}" < ${DUMB_FILE}
.PHONY: @restore_dump

@lint:
	docker run --rm -v $(CURRENT_DIR):/app -w /app golangci/golangci-lint:v1.57.1 golangci-lint run -v
.PHONY: @lint


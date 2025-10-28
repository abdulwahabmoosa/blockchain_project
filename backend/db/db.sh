#!/usr/bin/env bash

set -e

DB_USER="db"
DB_PASSWORD="mypassword"
DB_NAME="db"
DB_PORT="5432"
DB_CONTAINER_NAME="blockchain-db"
POSTGRES_VERSION="18"
MIGRATE_VERSION="v4.19.0"

DB_PATH="./db"
MIGRATIONS_PATH="${DB_PATH}/migrations"

DB_URL="postgres://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable"

COMMAND=$1

case "$COMMAND" in
  "start")
    docker run --name ${DB_CONTAINER_NAME} \
        -p ${DB_PORT}:${DB_PORT} \
        -e POSTGRES_PASSWORD=${DB_PASSWORD} \
        -e POSTGRES_USER=${DB_USER} \
        -e POSTGRES_DB=${DB_NAME} \
        -d postgres:${POSTGRES_VERSION}
    ;;

  "migrate-up")
    docker run --rm -v ${MIGRATIONS_PATH}:/migrations \
        --network host migrate/migrate:${MIGRATE_VERSION} \
        -path=/migrations/ \
        -database "${DB_URL}" up
    ;;

  "migrate-down")
    docker run --rm -v ${MIGRATIONS_PATH}:/migrations \
        --network host migrate/migrate:${MIGRATE_VERSION} \
        -path=/migrations/ \
        -database "${DB_URL}" down --all
    ;;

  "create-migration")
    MIGRATION_NAME=$2
    if [ -z "$MIGRATION_NAME" ]; then
      echo "Error: Migration name not provided."
      exit 1
    fi
    
    mkdir -p ${MIGRATIONS_PATH}

    docker run -it --rm --network host \
        --volume "${DB_PATH}:/db" \
        --user "$(id -u):$(id -g)" \
        migrate/migrate:${MIGRATE_VERSION} \
        create -ext sql -dir /db/migrations -seq "$MIGRATION_NAME"
    ;;
esac

exit 0


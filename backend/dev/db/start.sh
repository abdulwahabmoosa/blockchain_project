#!/usr/bin/env bash

set -e

DB_USER="blockchain-db"
DB_PASSWORD="password"
DB_NAME="db"
DB_PORT="5432"
DB_CONTAINER_NAME="blockchain-db"
POSTGRES_VERSION="18"
# MIGRATE_VERSION="v4.19.0"

# DB_PATH="./db"
# MIGRATIONS_PATH="${DB_PATH}/migrations"

# DB_URL="postgres://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable"

docker run --name ${DB_CONTAINER_NAME} \
  -p ${DB_PORT}:${DB_PORT} \
  -e POSTGRES_PASSWORD=${DB_PASSWORD} \
  -e POSTGRES_USER=${DB_USER} \
  -e POSTGRES_DB=${DB_NAME} \
  -d postgres:${POSTGRES_VERSION}

#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE USER restus ENCRYPTED PASSWORD 'restus' LOGIN;
	CREATE DATABASE restus OWNER restus;
EOSQL

psql -v ON_ERROR_STOP=1 --username "restus" --dbname "restus" -f /app/sql/init-db.sql

FROM postgres:14.1-alpine
COPY ./create_database.sql /docker-entrypoint-initdb.d/create_database.sql

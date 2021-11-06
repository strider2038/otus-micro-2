FROM golang:1-alpine AS build-env

COPY . /project

RUN set -xe \
    && cd /project/cmd/db-waiter \
    && go mod download \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o db-waiter

FROM scratch

ENV DATABASE_URL="" \
    MIGRATION_VERSION=""

COPY --from=build-env /project/cmd/db-waiter/db-waiter /app/

ENTRYPOINT ["/app/db-waiter"]

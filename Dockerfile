FROM golang:1.21-alpine3.17

LABEL maintainer="pluque01@correo.ugr.es" \
  version="1.1"

RUN adduser -D -u 1001 test

USER test

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

WORKDIR /app/test

ENTRYPOINT ["go", "run", "./build/", "-v", "test"]

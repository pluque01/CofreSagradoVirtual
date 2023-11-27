FROM golang:1.21-alpine3.17

LABEL maintainer="pluque01@correo.ugr.es" \
  version="1.0"

RUN adduser -D -u 1001 test

USER test

WORKDIR /app/test

COPY go.mod go.sum ./

ENTRYPOINT ["go", "run", "./build/", "-v", "test"]

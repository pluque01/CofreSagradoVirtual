FROM golang:latest AS build

FROM debian:stable-slim AS final

LABEL maintainer="pluque01@correo.ugr.es" \
  version="1.1"

# Multi-stage build to copy golang toolchain
COPY --from=golang:latest /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"
# Copy certificates
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

RUN adduser --disabled-password -u 1001 golanguser

USER golanguser

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

WORKDIR /app/test

ENTRYPOINT ["go", "run", "./build/", "-v", "test"]

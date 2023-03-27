# BASE IMAGE
FROM golang:1.19-alpine AS base

WORKDIR /app/cmd

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

RUN apk update \
    && apk add --no-cache \
    build-base \
    ca-certificates \
    curl \
    tzdata \
    git \
    vips \
    vips-dev \
    && update-ca-certificates

#RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go mod download && go mod verify


# DEVELOPMENT IMAGE
FROM base AS dev

WORKDIR /app/cmd

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

#RUN go install github.com/cosmtrek/air@latest

EXPOSE 6000
EXPOSE 2345

# CMD ["air", "-c", "./air.toml"]
CMD air

# BUILDER IMAGE
FROM base AS builder

WORKDIR /app/cmd

RUN go build -o workshop


# PRODUCTION IMAGE
FROM alpine:latest as prod

COPY --from=builder /app/workshop /usr/local/bin/workshop

EXPOSE 5000

ENTRYPOINT ["/usr/local/bin/sh-go-workshop/cmd"]

# Stage 1 - Binary Build
# BUILD_X args should be passed at build time as docker build args
FROM golang:1.9.1-alpine AS builder
ARG BUILD_TIME
ARG BUILD_VERSION
ARG BUILD_COMMIT
RUN apk update && apk add build-base libressl-dev
WORKDIR /go/src/queuemanager
COPY ./ /go/src/queuemanager
RUN COMPRESS_BINARY=1 GOBUILD_VERBOSE=1 BIN_NAME=bin make build

# Stage 2 - Final Image
# The application should be statically linked
FROM alpine:3.6
RUN apk update && apk add --no-cache ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /go/src/queuemanager/bin /usr/bin/queuemanager
VOLUME ["/etc/queuemanager"]
ENTRYPOINT ["queuemanager"]

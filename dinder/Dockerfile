# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:1.13-alpine AS builder

ENV CGO_ENABLED 0

ENV REDIS_URL redis:6379 

# Add Maintainer Info
LABEL maintainer="Ali Souidan <alsouidan@gmail.com>"

RUN apk --no-cache -U add git

RUN go get -u github.com/kardianos/govendor

# Set the Current Working Directory inside the container
WORKDIR /go/src/dinder
# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Install dependencies
RUN govendor sync

# Build the app
RUN govendor build 

# Smallest container image
FROM alpine

# RUN apk update && apk add --no-cache openssl

RUN apk update && apk add  ca-certificates

# Copy built executable from base image to here
COPY --from=builder go/src/dinder /

EXPOSE 8081

# Command to run the executable
CMD ["./dinder"]
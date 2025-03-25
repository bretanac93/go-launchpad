# syntax=docker/dockerfile:1
FROM golang:1.24 AS builder
ARG TARGETOS
ARG TARGETARCH
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make build APP_NAME=api GOOS=${TARGETOS} GOARCH=${TARGETARCH}

FROM alpine:latest

ARG TARGETOS
ARG TARGETARCH

WORKDIR /app

COPY --from=builder /app/bin/api-${TARGETOS}-${TARGETARCH} .

EXPOSE 8080
CMD ["./api-${TARGETOS}-${TARGETARCH}"]

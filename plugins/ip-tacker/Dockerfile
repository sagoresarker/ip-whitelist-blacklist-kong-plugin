FROM golang:1.21 AS builder

WORKDIR /build
COPY . .
RUN go mod download
# Build a static Linux binary for the plugin server
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags='-s -w' -o ip-tacker main.go

FROM kong:latest

USER root

RUN mkdir -p /usr/local/kong/plugins/ip-tacker

COPY --from=builder /build/ip-tacker /usr/local/kong/plugins/ip-tacker/

USER kong
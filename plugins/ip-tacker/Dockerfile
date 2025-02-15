FROM golang:1.23 AS builder

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -buildmode=plugin -o ip-tacker.so main.go

FROM kong:latest

USER root

RUN mkdir -p /usr/local/kong/plugins/ip-tacker

COPY --from=builder /build/ip-tacker.so /usr/local/kong/plugins/ip-tacker/
COPY --from=builder /build/ip-tacker.so  /kong

USER kong
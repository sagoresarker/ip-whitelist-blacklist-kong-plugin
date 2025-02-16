FROM golang:1.21 AS builder

WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -buildmode=plugin -o ip-tacker.so main.go
RUN CGO_ENABLED=1 GOOS=linux go build -o ip-tacker-server main.go

FROM kong:latest

USER root

RUN mkdir -p /usr/local/kong/plugins/ip-tacker

COPY --from=builder /build/ip-tacker.so /usr/local/kong/plugins/ip-tacker/
COPY --from=builder /build/ip-tacker-server /usr/local/kong/plugins/ip-tacker/


USER kong
# Build all services

FROM golang:1.23.6 AS builder

COPY . /app
WORKDIR /app

RUN GOPROXY=https://goproxy.cn GO111MODULE=on go mod download \
        && apt-get update && apt-get install -y make git \
        && make build

FROM alpine:3.17 AS admin-service
COPY --from=builder /app/config.toml /
COPY --from=builder /app/bin/admin-service /
CMD ["./admin-service"]

FROM alpine:3.17 AS gateway-service
COPY --from=builder /app/config.toml /
COPY --from=builder /app/bin/gateway-service /
CMD ["./gateway-service"]

FROM alpine:3.17 AS order-service
COPY --from=builder /app/config.toml /
COPY --from=builder /app/bin/order-service /
CMD ["./order-service"]

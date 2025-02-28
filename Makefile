VERSION= 0.0.1

ADMIN_NAME := admin-service
ADMIN_PATH := ./app/admin

GATEWAY_NAME := gateway-service
GATEWAY_PATH := ./app/gateway

ORDER_NAME := order-service
ORDER_PATH := ./app/order

# 构建目录
BIN_DIR := ./bin

# 默认目标
.DEFAULT_GOAL := build

build:
	@echo "Building..."
	@make build-admin
	@make build-gateway
	@make build-order

build-admin:
	@echo "Building admin..."
	@GOARCH=amd64 GOOS=linux go build -o $(BIN_DIR)/$(ADMIN_NAME) -ldflags "-X main.Version=$(VERSION)" $(ADMIN_PATH)/cmd

build-gateway:
	@echo "Building gateway..."
	@GOARCH=amd64 GOOS=linux go build -o $(BIN_DIR)/$(GATEWAY_NAME) -ldflags "-X main.Version=$(VERSION)" $(GATEWAY_PATH)/cmd

build-order:
	@echo "Building order..."
	@GOARCH=amd64 GOOS=linux go build -o $(BIN_DIR)/$(ORDER_NAME) -ldflags "-X main.Version=$(VERSION)" $(ORDER_PATH)/cmd

docker:
	@echo "Building docker image..."
	@docker build -t $(ADMIN_NAME):$(VERSION) --target admin-service .
	@docker build -t $(GATEWAY_NAME):$(VERSION) --target gateway-service .
	@docker build -t $(ORDER_NAME):$(VERSION) --target order-service .

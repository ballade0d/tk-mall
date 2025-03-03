VERSION= 0.0.1

ADMIN_NAME := admin-service
ADMIN_PATH := ./app/admin

CALLBACK_NAME := callback-service
CALLBACK_PATH := ./app/callback

GATEWAY_NAME := gateway-service
GATEWAY_PATH := ./app/gateway

ORDER_NAME := order-service
ORDER_PATH := ./app/order

PAYMENT_NAME := payment-service
PAYMENT_PATH := ./app/payment

USER_NAME := user-service
USER_PATH := ./app/user

# 构建目录
BIN_DIR := ./bin

# 默认目标
.DEFAULT_GOAL := build

build:
	@echo "Building..."
	@make build-admin
	@make build-callback
	@make build-gateway
	@make build-order
	@make build-payment
	@make build-user

build-admin:
	@echo "Building admin..."
	@GOARCH=amd64 GOOS=linux go build -o $(BIN_DIR)/$(ADMIN_NAME) -ldflags "-X main.Version=$(VERSION)" $(ADMIN_PATH)/cmd

build-callback:
	@echo "Building callback..."
	@GOARCH=amd64 GOOS=linux go build -o $(BIN_DIR)/$(CALLBACK_NAME) -ldflags "-X main.Version=$(VERSION)" $(CALLBACK_PATH)/cmd

build-gateway:
	@echo "Building gateway..."
	@GOARCH=amd64 GOOS=linux go build -o $(BIN_DIR)/$(GATEWAY_NAME) -ldflags "-X main.Version=$(VERSION)" $(GATEWAY_PATH)/cmd

build-order:
	@echo "Building order..."
	@GOARCH=amd64 GOOS=linux go build -o $(BIN_DIR)/$(ORDER_NAME) -ldflags "-X main.Version=$(VERSION)" $(ORDER_PATH)/cmd

build-payment:
	@echo "Building payment..."
	@GOARCH=amd64 GOOS=linux go build -o $(BIN_DIR)/$(PAYMENT_NAME) -ldflags "-X main.Version=$(VERSION)" $(PAYMENT_PATH)/cmd

build-user:
	@echo "Building user..."
	@GOARCH=amd64 GOOS=linux go build -o $(BIN_DIR)/$(USER_NAME) -ldflags "-X main.Version=$(VERSION)" $(USER_PATH)/cmd

docker:
	@echo "Building docker image..."
	@DOCKER_BUILDKIT=1 docker build -t $(ADMIN_NAME):$(VERSION) --target admin-service .
	@DOCKER_BUILDKIT=1 docker build -t $(CALLBACK_NAME):$(VERSION) --target callback-service .
	@DOCKER_BUILDKIT=1 docker build -t $(GATEWAY_NAME):$(VERSION) --target gateway-service .
	@DOCKER_BUILDKIT=1 docker build -t $(ORDER_NAME):$(VERSION) --target order-service .
	@DOCKER_BUILDKIT=1 docker build -t $(PAYMENT_NAME):$(VERSION) --target payment-service .
	@DOCKER_BUILDKIT=1 docker build -t $(USER_NAME):$(VERSION) --target user-service .

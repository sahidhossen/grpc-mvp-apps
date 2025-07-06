
OK_COLOR=\033[32;01m
NO_COLOR=\033[0m

PROTO=./proto/task_service.proto
PB_OUT=./proto
PROJECT=github.com/sahidhossen/todo/proto

BUILD_DIR ?= build

.PHONY: proto build run-api run-storage lint fmt run

proto: 
	protoc --go_out=$(PB_OUT) --go_opt=module=${PROJECT} --go-grpc_out=$(PB_OUT) --go-grpc_opt=module=${PROJECT} ${PROTO}

run-api:
	@echo "$(OK_COLOR)==> Starting API Gateway...$(NO_COLOR)"
	cd api-gateway && go run ./cmd/server || true

run-storage:
	@echo "$(OK_COLOR)==> Starting Storage Service...$(NO_COLOR)"
	cd storage-service && go run ./cmd/server || true

build:
	go build -o ${BUILD_DIR}/api-gateway ./api-gateway/cmd/server
	go build -o ${BUILD_DIR}/storage-service ./storage-service/cmd/server

test-api:
	@printf "$(OK_COLOR)==> Running Test $(NO_COLOR)\n"
	mkdir -p "${BUILD_DIR}/api-gateway"
	mkdir -p reports/api-gateway
	go test -v -race -coverprofile=reports/api-gateway/cov.out ./api-gateway/...

test-storage:
	@printf "$(OK_COLOR)==> Running Test $(NO_COLOR)\n"
	mkdir -p "${BUILD_DIR}/storage-service"
	mkdir -p reports/storage-service
	go test -v -race -coverprofile=reports/storage-service/cov.out ./storage-service/...

lint:
	@printf "$(OK_COLOR)==> Running Linter following https://github.com/golang/go/wiki/CodeReviewComments $(NO_COLOR)\n"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run --timeout=5m ./api-gateway/... ./storage-service/...

fmt:
	go install mvdan.cc/gofumpt@latest
	gofumpt -w ./api-gateway/... ./storage-service/...


# 发布时根据情况修改
APP = ginny-template
CONF = dev.yaml
#-------------------------------------	
.PHONY: run
run: tidy proto wire
	go run ./cmd/ -f configs/$(CONF)  & \
#-------------------------------------	
.PHONY: tidy
tidy:
	go mod tidy
#-------------------------------------	
.PHONY: tidy wire
wire:
	wire ./...
#-------------------------------------	
.PHONY: test
test: tidy mock
	go test -v ./internal/app/$(APP)/... -f `pwd`/configs/$(CONF) -covermode=count -coverprofile=dist/cover-$(APP).out
#-------------------------------------	
.PHONY: build
build: tidy
	GOOS=linux GOARCH="amd64" go build -o dist/$(APP)-linux-amd64 ./cmd/; \
	GOOS=darwin GOARCH="amd64" go build -o dist/$(APP)-darwin-amd64 ./cmd/
#-------------------------------------	
.PHONY: cover
cover: test
	go tool cover -html=dist/cover-$(APP).out
#-------------------------------------	
.PHONY: mock
mock:
	mockery --all
#-------------------------------------	
.PHONY: lint
lint:
	golint ./...
#-------------------------------------	
.PHONY: proto
proto:
    # 遇到错误： "--go_out: protoc-gen-go: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC" 解决：go install github.com/golang/protobuf/protoc-gen-go@latest
	protoc -I api/proto ./api/proto/* --go_out=plugins=grpc:api/proto
#-------------------------------------	
.PHONY: docker
docker-compose: build dash rules
	docker-compose -f deploy/docker-compose.yml up --build -d
all: lint cover docker

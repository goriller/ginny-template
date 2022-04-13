# 发布时根据情况修改
APP = APP_NAME
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
.PHONY: wire
wire: tidy
	wire ./...
#-------------------------------------	
.PHONY: test
test: tidy mock
	go test -v ./internal/... -f `pwd`/configs/$(CONF) -covermode=count -coverprofile=dist/cover-$(APP).out
#-------------------------------------	
.PHONY: build
build: tidy
	GOARCH="amd64" go build -o dist/$(APP)-amd64 ./cmd/; \

.PHONY: build_linux
build_linux: tidy
	GOOS=linux GOARCH="amd64" go build -o dist/${BINARY}-linux-amd64  ./cmd/

.PHONY: build_darwin
build_darwin: tidy
	GOOS=darwin GOARCH="amd64" go build -o dist/${BINARY}-darwin-amd64 ./cmd/

.PHONY: build_win
build_win: tidy
	GOOS=windows GOARCH="amd64" go build -o dist/${BINARY}-windows-amd64.exe ./cmd/
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
# fetches this repo into $GOPATH
# go install github.com/envoyproxy/protoc-gen-validate@latest
# go install github.com/golang/protobuf/protoc-gen-go@latest
.PHONY: proto
proto:
	protoc \
  -I api/proto \
  --go_out="plugins=grpc:api/proto" \
  --validate_out="lang=go:api/proto" \
  ./api/proto/*.proto
#-------------------------------------	
.PHONY: docker
docker-compose: build dash rules
	docker-compose -f deploy/docker-compose.yml up --build -d
all: lint cover docker

# How to run "make" command in gitbash in windows?
# https://gist.github.com/evanwill/0207876c3243bbb6863e65ec5dc3f058
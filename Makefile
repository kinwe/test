# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=main
BINARY_UNIX=$(BINARY_NAME)_unix
BUILD_PATH=./build

all: deps start

build:
	$(GOBUILD) -o $(BUILD_PATH)/$(BINARY_NAME) -v ./...
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BUILD_PATH)/$(BINARY_NAME)
	rm -f $(BUILD_PATH)/$(BINARY_UNIX)
start:clean
	$(GOBUILD) -o $(BUILD_PATH)/$(BINARY_NAME)
	$(BUILD_PATH)/$(BINARY_NAME)
deps:
	export GO111MODULE=on
	export GOPROXY=https://mirrors.aliyun.com/goproxy/



# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
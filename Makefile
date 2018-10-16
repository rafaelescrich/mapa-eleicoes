APP_NAME=mapa-eleicao
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=$(APP_NAME)
BINARY_LINUX=$(BINARY_NAME)_linux
BUILD_ENV= CGO_ENABLED=0 GOTRACEBACK=none

all: clean deps test build

build: 
	$(BUILD_ENV) $(GOBUILD) -o $(BINARY_NAME) -v

download:
	xargs -n 1 curl -O < urls.txt

unzip:
	unzip 'files/*.zip' -d files

test: 
	$(GOTEST) -v ./...

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run: build
	env $$(grep -v '^#' $(ENV_FILE) | xargs) ./$(BINARY_NAME)

deps:
	$(GOGET) ./...

# Cross compilation
build-linux:
	$(BUILD_ENV) GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_LINUX) -v

GOPATH=$(shell pwd)/vendor:$(shell pwd)
GOBIN=$(shell pwd)/dist
GOFILES=$(wildcard *.go)
GONAME=gosrv

all: lint test build

build:
	@echo "Building $(GOFILES) to ./dist"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -mod=vendor -o dist/$(GONAME) $(GOFILES)

get:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get .

install:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go install $(GOFILES)

lint:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go vet ./...

test:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test ./... -coverprofile coverage.out -covermode=atomic -v

test_html:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go tool cover -html=coverage.out

start:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run $(GOFILES)

watch:
	gin -p 4000 -a 4001 -i run main.go

.PHONY: build get install start watch
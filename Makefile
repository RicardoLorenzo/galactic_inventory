.MAIN: all

GOCMD=go
GOBINARY=server
GITREP=github.com/RicardoLorenzo/server-test

deps:
	go mod vendor
	go mod tidy

init:
	go mod init $(GITREP)
	go mod tidy

test: deps
	$(GOCMD) test ./...

run:
	./bin/$(GOBINARY)-darwin

build: deps
	GOOS=linux GOARCH=amd64 $(GOCMD) build -o bin/$(GOBINARY)-linux go/*.go
	GOOS=darwin GOARCH=amd64 $(GOCMD) build -o bin/$(GOBINARY)-darwin go/*.go

docker:
	docker build .

all: test build
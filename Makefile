.PHONY: build
start:
	go build -v ./cmd/apiServer
	.\apiServer.exe

build:
	swag init -g .\cmd\apiServer\main.go
	go build -v ./cmd/apiServer

fullBuild:
	swag init -g .\cmd\apiServer\main.go
	go build -v ./cmd/apiServer

.PHONY: test
test:
	go test -v -race -timeout 15s ./...

.DEFAULT_GOAL := build

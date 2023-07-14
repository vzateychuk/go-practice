.DEFAULT_GOAL := build

fmt:
 go fmt ./...
.PHONY:fmt

lint: fmt
 golangci-lint ./...
.PHONY:lint

vet: fmt
 go vet ./...
.PHONY:vet

build: vet
 go build hello.go
.PHONY:buil

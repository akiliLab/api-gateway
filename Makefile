.PHONY: install test build serve clean pack deploy ship

TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

install:
	go get ./...

test: install
	go test ./...

build:
	 docker build -t api-gateway:v1 .

proto-gen:
	protoc -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I. --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. ./proto/balance/*.proto
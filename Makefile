.PHONY: install test build serve clean pack deploy ship

TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

install:
	go get ./...

test: install
	go test ./...

proto-gen:
	 protoc --go_out=plugins=grpc:. ./balance/pb/*.proto
	 protoc --go_out=plugins=grpc:. ./transaction/pb/*.proto

build-balance:
	 docker build -t balance:v1 \
	 --build-arg ssh_prv_key="$(cat ~/.ssh/id_rsa)" \
	 --build-arg ssh_pub_key="$(cat ~/.ssh/id_rsa.pub)" \
	 --squash ./balance

build-transaction:
	 cd transaction && docker build -t transaction:v1 \
	 --build-arg ssh_prv_key="$(cat ~/.ssh/id_rsa)" \
	 --build-arg ssh_pub_key="$(cat ~/.ssh/id_rsa.pub)" \
	 --squash ./transaction

build:
	 docker build -t api-gateway:v1 \
	 --build-arg ssh_prv_key="$(cat ~/.ssh/id_rsa)" \
	 --build-arg ssh_pub_key="$(cat ~/.ssh/id_rsa.pub)" \
	 --squash .
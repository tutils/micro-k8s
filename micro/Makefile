
GOPATH:=$(shell go env GOPATH)

.PHONY: build
build:
	go build -o micro *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: build
	docker build . -t 10.9.19.140/test/micro:latest

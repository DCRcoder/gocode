SHELL="/bin/bash"
GOPATH=$(shell go env GOPATH)
GOBIN=$(GOPATH)/bin/

build:
	go fmt ./...
	go build -v ./cmd/gocode.go
	cp ./gocode $(GOBIN)

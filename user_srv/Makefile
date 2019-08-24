
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path="proto/user"  --micro_out="proto/user" --go_out="proto/user" proto/user/*.proto

.PHONY: build
build: proto

	go build -o user-srv main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t user-srv:v1

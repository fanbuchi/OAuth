
.PHONY: proto
proto:
	protoc --proto_path="proto/oauth"  --micro_out="proto/oauth" --go_out="proto/oauth" proto/oauth/*.proto

.PHONY: build
build: proto

	go build -o oauth-srv main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t oauth-srv:v1

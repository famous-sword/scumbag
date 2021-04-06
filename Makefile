.PHONY : clean

buildTime = `date +%Y-%m-%dT%T%z`
target = cmd/web.go
ldflags = -ldflags="-s -w -X main.buildTime=${buildTime}"
gcflags = -gcflags="-trimpath=${PWD}"
output = -o=doc-cloud

build:
	CGO_ENABLED=0 go build ${ldflags} ${gcflags} ${output} ${target}

tidy:
	go mod tidy

clean:
	rm -rf ./var/storage/*

transcoding:
	protoc --proto_path=./transcoding --go_out=./transcoding --go-grpc_out=./transcoding transcoding.proto

build-docker:
	docker build -f build/transcoder/Dockerfile . -t "africanwave/transcoder:1.0"
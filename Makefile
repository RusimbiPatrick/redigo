BINARY=redigo
VERSION=0.1.0

build:
	go build -o ${BINARY} cmd/redigo-server/main.go

run: build
	./${BINARY}

test:
	go test -v ./...

clean:
	go clean
	rm -f ${BINARY}

release:
	GOOS=linux GOARCH=amd64 go build -o ${BINARY}-linux-amd64
	GOOS=darwin GOARCH=amd64 go build -o ${BINARY}-darwin-amd64

.PHONY: build run test clean release
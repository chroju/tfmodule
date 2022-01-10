BINARY_NAME=tfmodule

.PHONY: lint test install build run clean

lint:
	gofmt -s -l .
	golint ./...
	go vet ./...

test: lint
	go test -v ./... -cover

install: test
	go install

build: test
	rm -rf bin/
	mkdir bin
	go build -o bin/${BINARY_NAME}

run:
	go run main.go

clean:
	go fmt
	go mod tidy
	go clean
	rm -f bin/

BINARY_NAME=process_tracker
HOSTOS=$(shell uname -s | awk '{print tolower($0)}')

build/${BINARY_NAME}:
	GOARCH=arm64 GOOS=darwin go build -o ./build/${BINARY_NAME}-darwin ./cmd/process_tracker/main.go
	GOARCH=amd64 GOOS=linux go build -o ./build/${BINARY_NAME}-linux ./cmd/process_tracker/main.go
	GOARCH=amd64 GOOS=windows go build -o ./build/${BINARY_NAME}-windows ./cmd/process_tracker/main.go

run: build/${BINARY_NAME}
	./build/${BINARY_NAME}-${HOSTOS}

clean:
	go clean
	rm ./build/${BINARY_NAME}-darwin
	rm ./build/${BINARY_NAME}-linux
	rm ./build/${BINARY_NAME}-windows

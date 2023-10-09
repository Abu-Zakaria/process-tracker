PROCESS_TRACKER=process_tracker
REPORT_GENERATOR=report_generator
HOSTOS=$(shell uname -s | awk '{print tolower($0)}')

build_main:
	GOARCH=arm64 GOOS=darwin go build -o ./build/${PROCESS_TRACKER}-darwin ./cmd/process_tracker/main.go
	GOARCH=amd64 GOOS=linux go build -o ./build/${PROCESS_TRACKER}-linux ./cmd/process_tracker/main.go
	GOARCH=amd64 GOOS=windows go build -o ./build/${PROCESS_TRACKER}-windows ./cmd/process_tracker/main.go

build_report_generator:
	GOARCH=arm64 GOOS=darwin go build -o ./build/${REPORT_GENERATOR}-darwin ./cmd/report_generator/main.go
	GOARCH=amd64 GOOS=linux go build -o ./build/${REPORT_GENERATOR}-linux ./cmd/report_generator/main.go
	GOARCH=amd64 GOOS=windows go build -o ./build/${REPORT_GENERATOR}-windows ./cmd/report_generator/main.go

run: build_main
	./build/${BINARY_NAME}-${HOSTOS}

run_report_generator: build_report_generator
	./build/${REPORT_GENERATOR}-${HOSTOS}

clean:
	go clean
	rm ./build/${BINARY_NAME}-darwin
	rm ./build/${BINARY_NAME}-linux
	rm ./build/${BINARY_NAME}-windows
	rm ./build/${REPORT_GENERATOR}-darwin
	rm ./build/${REPORT_GENERATOR}-linux
	rm ./build/${REPORT_GENERATOR}-windows

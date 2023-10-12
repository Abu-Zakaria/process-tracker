PROCESS_TRACKER=process_tracker
TOP_TEN_REPORT=top_ten_report
HOSTOS=$(shell uname -s | awk '{print tolower($0)}')

build_main:
	GOARCH=arm64 GOOS=darwin go build -o ./build/${PROCESS_TRACKER}-darwin ./cmd/process_tracker/main.go
	GOARCH=amd64 GOOS=linux go build -o ./build/${PROCESS_TRACKER}-linux ./cmd/process_tracker/main.go
	GOARCH=amd64 GOOS=windows go build -o ./build/${PROCESS_TRACKER}-windows ./cmd/process_tracker/main.go

build_top_ten_report:
	GOARCH=arm64 GOOS=darwin go build -o ./build/${TOP_TEN_REPORT}-darwin ./cmd/top_ten_report/main.go
	GOARCH=amd64 GOOS=linux go build -o ./build/${TOP_TEN_REPORT}-linux ./cmd/top_ten_report/main.go
	GOARCH=amd64 GOOS=windows go build -o ./build/${TOP_TEN_REPORT}-windows ./cmd/top_ten_report/main.go

run: build_main
	GOOS=${HOSTOS} ./build/${PROCESS_TRACKER}-${HOSTOS}

run_top_ten_report: build_top_ten_report
	GOOS=${HOSTOS} ./build/${TOP_TEN_REPORT}-${HOSTOS}

clean:
	go clean
	rm ./build/${PROCESS_TRACKER}-darwin
	rm ./build/${PROCESS_TRACKER}-linux
	rm ./build/${PROCESS_TRACKER}-windows
	rm ./build/${TOP_TEN_REPORT}-darwin
	rm ./build/${TOP_TEN_REPORT}-linux
	rm ./build/${TOP_TEN_REPORT}-windows

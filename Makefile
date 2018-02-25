build:
	@go build -x -v -o shell_logger client/client.go

build_linux64:
	@env GOARCH=amd64 GOOS=linux go build -x -v -o shell_logger_linux64 client/client.go

build_darwin64:
	@env GOARCH=amd64 GOOS=darwin go build -x -v -o shell_logger_darwin64 client/client.go

build_windows64:
	@env GOARCH=amd64 GOOS=windows go build -x -v -o shell_logger_windows64.exe client/client.go

build_all: build_linux64 build_darwin64 build_windows64

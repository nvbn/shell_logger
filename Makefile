all: linux darwin

linux:
	mkdir -p build; \
	dep ensure; \
	GOOS=linux GOARCH=amd64 go build -o build/shell_logger_linux shell_logger.go

darwin:
	mkdir -p build; \
	dep ensure; \
	GOOS=darwin GOARCH=amd64 go build -o build/shell_logger_darwin shell_logger.go

test:
	dep ensure; \
	test -z $(go fmt ./...) && \
	go test -v ./...

fmt:
	go fmt ./...

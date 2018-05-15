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
	go test -race -v ./...

functional_test:
	dep ensure; \
	pip3 install --user -r functional_tests/requirements.txt; \
	GOOS=linux GOARCH=amd64 go build -o functional_tests/shell_logger shell_logger.go; \
	py.test functional_tests -vvvv --capture=sys

fmt:
	go fmt ./...

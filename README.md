# shell_logger [WIP]

[![Build Status](https://travis-ci.org/nvbn/shell_logger.svg?branch=master)](https://travis-ci.org/nvbn/shell_logger)

Dead simple logger of shell commands with output and exit codes.

## Installation

Only 64 bit Linux and osx with zsh are supported at the moment:

```bash
sh -c "$(curl -sL https://raw.githubusercontent.com/nvbn/shell_logger/master/install.sh)"
```

## Development

Install dependencies with [dep](https://github.com/golang/dep):

```bash
dep ensure
```

Run tests:

```bash
make test
```

Run functional tests (requires docker and python 3.4+):

```bash
make functional_test
```

Build:

```bash
make linux
make darwin
make all
```

## License MIT

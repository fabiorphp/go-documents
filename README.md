# Brazanation Documents - Golang

[![Build Status](https://img.shields.io/travis/brazanation/go-documents/master.svg?style=flat-square)](https://travis-ci.org/brazanation/go-documents)
[![Codecov branch](https://img.shields.io/codecov/c/github/brazanation/go-documents/master.svg?style=flat-square)](https://codecov.io/gh/brazanation/go-documents)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/brazanation/go-documents)
[![Go Report Card](https://goreportcard.com/badge/github.com/brazanation/go-documents?style=flat-square)](https://goreportcard.com/report/github.com/brazanation/go-documents)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](https://github.com/brazanation/go-documents/blob/master/LICENSE)

A Golang library to provide Brazilian Documents safer, easier and fun!

## Installation

go-documents requires Go 1.11 or later.

```
go get github.com/brazanation/go-documents
```

If you want to get an specific version, please use the example below:

```
go get gopkg.in/brazanation/go-documents.v0
```

## Development

### Requirements

- Install [Go](https://golang.org)

### Makefile
```sh
// Clean up
$ make clean

// Run tests and generates html coverage file
make cover

// Download project dependencies
make depend

// Format all go files
make fmt

// Run linters
make lint

// Run tests
make test
```

## Documentation

Read the full documentation at [https://godoc.org/github.com/brazanation/go-documents](https://godoc.org/github.com/brazanation/go-documents).


## License

This project is released under the MIT licence. See [LICENSE](https://github.com/brazanation/go-documents/blob/master/LICENSE) for more details.

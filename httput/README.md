# Httput

[![Source](https://img.shields.io/badge/git-source-orange.svg?style=flat-square)](https://github.com/ymohl-cl/gopkg/tree/httput-release/httput)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/ymohl-cl/gopkg/httput)
[![Build Status](https://travis-ci.org/ymohl-cl/gopkg.svg?branch=httput-release&style=flat-square)](https://travis-ci.org/ymohl-cl/gopkg)
[![codecov](https://codecov.io/gh/ymohl-cl/gopkg/branch/httput-release/graph/badge.svg?style=flat-square)](https://codecov.io/gh/ymohl-cl/gopkg)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/ymohl-cl/gopkg/httput-release/LICENSE)

httput package provide a tool suite to make an unitaries tests arround http and echo framework.

## Requirements

Golang 1.12.4 or higher

download the package

``` bash
go get -u github.com/ymohl-cl/gopkg/httput
```

## Usage

``` Golang
import "github.com/ymohl-cl/gopkg/httput"

func main() {
    // build a request
    request := httptest.NewRequest(...)

    // build a fake echo context
    c := httput.NewContext(request)

    // the call your handlers to test them
    _ = MyHandler(c)

    // finally print the response content
    fmt.Printf("status code: %d", c.Rec.Code)
    fmt.Printf("content body: %s", c.Rec.Body.String())
}
```

Output:

``` bash
> "status code: 200"
> "content body: {\"status\":\"200\"}"
```

## Changelog

### v0.0.1

Initial commit

- fake echo context to tests handlers
- tests and documentation
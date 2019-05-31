# Errorx

[![Source](https://img.shields.io/badge/git-source-orange.svg?style=flat-square)](https://github.com/ymohl-cl/gopkg/tree/errorx-release/errorx)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/ymohl-cl/gopkg/errorx)
[![Build Status](https://travis-ci.org/ymohl-cl/gopkg.svg?branch=errorx-release&style=flat-square)](https://travis-ci.org/ymohl-cl/gopkg)
[![codecov](https://codecov.io/gh/ymohl-cl/gopkg/branch/errorx-release/graph/badge.svg?style=flat-square)](https://codecov.io/gh/ymohl-cl/gopkg)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/ymohl-cl/gopkg/errorx-release/LICENSE)

The errorx package implement a manager error which implement the standar error interface.

Provide a dictionnary and wrapper errors.

## Requirements

Golang 1.12.4 or higher

download the package

``` bash
go get -u github.com/ymohl-cl/gopkg/errorx
```

## Usage

``` Golang
import "github.com/ymohl-cl/gopkg/errorx"

func main() {
    // build a dictionnary error
    d := NewDico()
    d.SetEntry(uint64(0), "error example 1")
    d.SetEntry(uint64(1), "error example 2")

    // then catch an error
    err := func() errorx.Errorx {
        return d.Error(uint64(0))
    }()
    err.Wrap("context error example")

    // finally print error to see the result
    fmt.Println(err.Error())
}
```

Output:

``` bash
> "context error example -> error example 1"
```

## Changelog

### v1.0.0

Initial commit

- dictionnary management
- implement standar error interface
- tests and documentation

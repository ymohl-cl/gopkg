# Example

[![Source](https://img.shields.io/badge/git-source-orange.svg?style=flat-square)](https://github.com/ymohl-cl/gopkg/tree/example-release/example)
[![Build Status](https://travis-ci.org/ymohl-cl/gopkg.svg?branch=example-release&style=flat-square)](https://travis-ci.org/ymohl-cl/gopkg)
[![codecov](https://codecov.io/gh/ymohl-cl/gopkg/branch/example-release/graph/badge.svg?style=flat-square)](https://codecov.io/gh/ymohl-cl/gopkg)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/ymohl-cl/gopkg/example-release/LICENSE)

the example package shows an example of the life cycle of a go package on this repository

## Requirements

Golang 1.12.4 or higher

download the package

``` bash
go get -u github.com/ymohl-cl/gopkg/example
```

## Usage

``` Golang
package main

import (
    "fmt"

    "github.com/ymohl-cl/gopkg/example"
)

func main() {
    fmt.Println(example.Hello("toto"))
}
```

Output:

``` bash
> "Hello toto !"
```

## Changelog

### v0.0.1

- Tests example
- Documentation example
# Go pkg

Go pkg is a tool box to develop with the Golang language.

To maintain a clear history to each packages, this main branch describe the global documentation and the continue integration process.

## Table of contents

- [Go pkg](#go-pkg)
  - [Table of contents](#table-of-contents)
  - [Packages](#packages)
    - [example](#example)
    - [errorx](#errorx)
    - [httput](#httput)
    - [server](#server)
    - [gosource](#gosource)
  - [Requirements](#requirements)
  - [Build usage](#build-usage)
    - [Rules list](#rules-list)
  - [Repository usage](#repository-usage)
    - [Versionning](#versionning)
    - [CI rules](#ci-rules)
      - [Triggers list and behaviors](#triggers-list-and-behaviors)
    - [New package](#new-package)
    - [Update package](#update-package)
    - [Update versionning package](#update-versionning-package)
    - [How to contribute](#how-to-contribute)
      - [Process to contribute](#process-to-contribute)
  - [Contact](#contact)
  - [Acknowledgements](#acknowledgements)

## Packages

### example

[![Source](https://img.shields.io/badge/git-source-orange.svg?style=flat-square)](https://github.com/ymohl-cl/gopkg/tree/example-release/example)
[![Build Status](https://travis-ci.org/ymohl-cl/gopkg.svg?branch=example-release&style=flat-square)](https://travis-ci.org/ymohl-cl/gopkg)
[![codecov](https://codecov.io/gh/ymohl-cl/gopkg/branch/example-release/graph/badge.svg?style=flat-square)](https://codecov.io/gh/ymohl-cl/gopkg)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/ymohl-cl/gopkg/example-release/LICENSE)

example package provide an example usage to package life in this library.

### errorx

[![Source](https://img.shields.io/badge/git-source-orange.svg?style=flat-square)](https://github.com/ymohl-cl/gopkg/tree/errorx-release/errorx)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/ymohl-cl/gopkg/errorx)
[![Build Status](https://travis-ci.org/ymohl-cl/gopkg.svg?branch=errorx-release&style=flat-square)](https://travis-ci.org/ymohl-cl/gopkg)
[![codecov](https://codecov.io/gh/ymohl-cl/gopkg/branch/errorx-release/graph/badge.svg?style=flat-square)](https://codecov.io/gh/ymohl-cl/gopkg)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/ymohl-cl/gopkg/errorx-release/LICENSE)

errorx package provide an error management which implement the standar error interface.

### httput

[![Source](https://img.shields.io/badge/git-source-orange.svg?style=flat-square)](https://github.com/ymohl-cl/gopkg/tree/httput-release/httput)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/ymohl-cl/gopkg/httput)
[![Build Status](https://travis-ci.org/ymohl-cl/gopkg.svg?branch=httput-release&style=flat-square)](https://travis-ci.org/ymohl-cl/gopkg)
[![codecov](https://codecov.io/gh/ymohl-cl/gopkg/branch/httput-release/graph/badge.svg?style=flat-square)](https://codecov.io/gh/ymohl-cl/gopkg)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/ymohl-cl/gopkg/httput-release/LICENSE)

httput package provide a tool suite to make an unitaries tests arround http.

### server

[![Source](https://img.shields.io/badge/git-source-orange.svg?style=flat-square)](https://github.com/ymohl-cl/gopkg/tree/server-release/server)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/ymohl-cl/gopkg/server)
[![Build Status](https://travis-ci.org/ymohl-cl/gopkg.svg?branch=server-release&style=flat-square)](https://travis-ci.org/ymohl-cl/gopkg)
[![codecov](https://codecov.io/gh/ymohl-cl/gopkg/branch/server-release/graph/badge.svg?style=flat-square)](https://codecov.io/gh/ymohl-cl/gopkg)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/ymohl-cl/gopkg/server-release/LICENSE)

server package provide a common server http.

### gosource

[![Source](https://img.shields.io/badge/git-source-orange.svg?style=flat-square)](https://github.com/ymohl-cl/gopkg/tree/gosource-release/gosource)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/ymohl-cl/gopkg/gosource)
[![Build Status](https://travis-ci.org/ymohl-cl/gopkg.svg?branch=gosource-release&style=flat-square)](https://travis-ci.org/ymohl-cl/gopkg)
[![codecov](https://codecov.io/gh/ymohl-cl/gopkg/branch/gosource-release/graph/badge.svg?style=flat-square)](https://codecov.io/gh/ymohl-cl/gopkg)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/ymohl-cl/gopkg/gosource-release/LICENSE)

gosource package allows write a go source file dynamically.

## Requirements

- GNU Make >= 4.1
- Compilation tested only on debian and on ubuntu Xenial

## Build usage

The Makefile build one makefile to each packages.

process build:

- print skin gopkg
- install dependancies
- linter with golint
- test and generate coverage
- build packages

``` bash
make
```

### Rules list

To the make's rules work on the packages, the package need initialized with go module and have a go.mod file.

Make's rules are executed to each go packages and invoke the same rule to their Makefile (the generated makefile behavior are describe by the '*' character).

``` bash
# install: create Makefile to each go package funded and invoke their rule 'install'
# * download the dependencies describes by the go.mod file
make install
```

``` bash
# lint: invoke the 'lint' rule to each generated makefile
# * run golint tool
make lint
```

``` bash
# test: invoke the 'test' rule to each generated makefile
# * run tests and generate the rate coverage
make test
```

``` bash
# build: invoke the 'build' rule to each generated makefile
# * run the build
make build
```

``` bash
# clean: invoke the 'clean' rule to each generated makefile
# * remove the coverages files and binaries files
make clean
```

``` bash
# fclean: invoke the 'clean' rule and remove all generated makefile
make fclean
```

``` bash
# skin: draw the gopkg's skin
make skin
```

## Repository usage

The main branch describe only continue integration and documentation.

To see the specific documentation to one package, click on the github source badge associated in the [Packages section](#packages)

### Versionning

Each packages are versioned with [semver rules 2.0.0](https://semver.org/).

``` markdown
    - MAJOR version when you make incompatible API changes
    - MINOR version when you add functionality in a backwards-compatible manner
    - PATCH version when you make backwards-compatible bug fixes.
```

### CI rules

The ci is executed by [travis ci](https://travis-ci.org/ymohl-cl/gopkg).

All packages start from main branch to inherit of ci functionnalities.

Next, equivalent of master to the new package, should be branchname-release

*-release branchs are protected to be up to date. So force push will not be possible.

#### Triggers list and behaviors

If branch name match with the next patterns, each commits will be executed by travis ci.

- .*-dev-.*:
  - linter
  - test
  - build
- .*-release$
  - linter
  - test
  - build
  - release creation

### New package

Example process to create a new package 'example'

``` bash
git clone git@github.com:ymohl-cl/gopkg.git && cd gopkg
git checkout main
git checkout -b example-release
# make your dev ... and add your first version in the package's README file
vim example/README.md
# clean history (important: don't modify the commits which comes from the main branch)
git rebase -i HEAD~?
git push origin example-release
```

### Update package

Example process to update an existing package 'example'

``` bash
git clone git@github.com:ymohl-cl/gopkg.git && cd gopkg
git checkout example-release
git checkout -b example-dev-myfeaturename
# make your dev ... and update your version in the package's README file
vim example/README.md
# clean history (important: don't modify the commits which comes from the main branch)
git rebase -i example-release
# import ci update if exist
git merge main
git push origin example-dev-myfeaturename
# when the pull request will be accepted, the new release will be automatically created
```

### Update versionning package

You need add the release number in the README.md file for the ci detect the creation when your commit has been merged in the release branch.

The package's README.md file must describe the releases like this example:

``` markdown
## Changelog

### v0.1.2

describe the release like you want

### v0.1.1

describe the release like you want

### v0.1.0

describe the release like you want
```

- only the title '## Changelog' and the next title '###' will be catched
- Order your releases from most recent to the oldest

### How to contribute

Hello guys, thanks for you interest.

Before contribute, you can open an issue to explain your problem(s) and your solution(s).
If you prefer, you can [contact me](#contact).

#### Process to contribute

Fork the repositry and submit your feature(s) on the pull request.

Please, respect the process to [update package](#update-package) as musch as possible.

## Contact

mohl.clauzade@gmail.com

## Acknowledgements

Thanks for reading so far ;)
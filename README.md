# Go pkg

Go pkg is a tool box to dev projects

## errorx

errorx define a specific error management with dictionnary and wrapper
Errorx type implement the standard error interface

### Usage

``` Golang
import "github.com/ymohl-cl/gopkg/errorx"
```

To build a dictionnary:

``` Golang
    d := NewDico()
    d.SetEntry(uint64(0), "message_1")
    d.SetEntry(uint64(1), "message_2")
```

To get an error from dictionnary:

``` Golang
    err := d.Error(uint64(0))
```

To Wrap an error on other error:

``` Golang
    err.Wrap("message_3")
    fmt.Println(err.Error())
    // should print: message_3 -> message_1
```

### Realeases

* error/v1.0.0 implement errorx which provide an error manager
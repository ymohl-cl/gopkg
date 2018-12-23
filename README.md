# Go pkg

Go pkg is a tool box to dev projects

## errorx

errorx define a specific error management with dictionnary and wrapper
Errorx type implement the standard error interface

### errorx usage

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

### errorx realeases

* error/v1.0.0 implement errorx which provide an error manager

## httput

httpput is a tool box to http unitary test

### httput usage

``` Golang
import "github.com/ymohl-cl/gopkg/httput"
```

To get a new context for your handler and run your test:

``` Golang
    c := NewContext(request)
    MyHandler(c.Input)
```

You could find the content response on Rec field

``` Golang
    fmt.Printf("status code: %d", c.Rec.Code)
    fmt.Printf("content body: %s", c.Rec.Body.String())
```

### httput realeases

* httput/v1.0.0 implement httput which provide an echo context to test your handlers

## server

server provide a TLS http server with a logger request
server provide a ping route

### server usage

Requirement:

``` bash
export MYAPP_SSL_CERTIFICATE="path_to_certificats/cert.pem"
export MYAPP_SSL_KEY="path_to_key/key.pem"
export MYAPP_PORT="4242"
```

``` Golang
import "github.com/ymohl-cl/gopkg/server"
```

To create a new server:

``` Golang
    s, err := New("appName")
    if err != nil {
        panic(err)
    }
    if err = s.Start(); err != nil {
        panic(err)
    }
```

### server realeases

* server/v1.0.0 implement server which provide a tls http server

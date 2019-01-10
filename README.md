# Go pkg

Go pkg is a tool box to dev projects

## httput

httput is a tool box to http unitary test

### Usage

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

### Realeases

* httput/v1.0.0 implement httput which provide an echo context to test your handlers
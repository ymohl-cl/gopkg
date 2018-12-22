# Go pkg

Go pkg is a tool box to dev projects

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
# Favicon sysadmServer's middleware

sysadmServer middleware to support favicon.

## Usage

### Start using it

Download and install it:

```sh
$ go get github.com/wangyysde/sysadmServer/favicon
```

Import it in your code:

```go
import "github.com/wangyysde/sysadmServer/favicon"
```

### Canonical example:

```go
package main
            
import (
    "github.com/wangyysde/sysadmServer"
    "github.com/wangyysde/sysadmServer/favicon"
)
            
func main() {
    r := sysadmServer.Default()
    r.Use(favicon.New("./favicon.ico")) // set favicon middleware 

    r.GET("/ping", func(c *sysadmServer.Context) {
        c.String(200, "Hello favicon.")
    })

    r.Run(":8080")
}
```

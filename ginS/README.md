# sysadmServer Default Server

This is API experiment for sysadmServer.

```go
package main

import (
	"github.com/wangyysde/sysadmServer"
	"github.com/wangyysde/sysadmServer/ginS"
)

func main() {
	ginS.GET("/", func(c *sysadmserver.Context) { c.String(200, "Hello World") })
	ginS.Run()
}
```

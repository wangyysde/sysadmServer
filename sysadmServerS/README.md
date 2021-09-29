# sysadmServer Default Server

This is API experiment for sysadmServer.

```go
package main

import (
	"github.com/wangyysde/sysadmServer"
	"github.com/wangyysde/sysadmServer/sysadmServerS"
)

func main() {
	sysadmServerS.GET("/", func(c *sysadmserver.Context) { c.String(200, "Hello World") })
	sysadmServerS.Run()
}
```

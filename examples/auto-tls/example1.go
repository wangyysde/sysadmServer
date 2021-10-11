package main

import (
	"fmt"

	"github.com/wangyysde/sysadmServer"
	"github.com/wangyysde/sysadmServer/autotls"
)

func main() {
	r := sysadmServer.Default()

	// Ping handler
	r.GET("/ping", func(c *sysadmServer.Context) {
		c.String(200, "pong")
	})

	sysadmServer.Log(fmt.Sprintf("%s",autotls.Run(r, "harbor.sysadm.cn")),"info")
}

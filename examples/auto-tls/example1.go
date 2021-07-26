package main

import (
	"log"

	"github.com/wangyysde/sysadmServer/autotls"
	"github.com/wangyysde/sysadmServer"
)

func main() {
	r := sysadmServer.Default()

	// Ping handler
	r.GET("/ping", func(c *sysadmServer.Context) {
		c.String(200, "pong")
	})

	log.Fatal(autotls.Run(r, "harbor.sysadm.cn"))
}

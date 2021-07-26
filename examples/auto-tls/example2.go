package main

import (
	"log"

	"github.com/wangyysde/sysadmServer/autotls"
	"github.com/wangyysde/sysadmServer"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	r := sysadmServer.Default()

	// Ping handler
	r.GET("/ping", func(c *sysadmServer.Context) {
		c.String(200, "pong")
	})

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("harbor.sysadm.cn"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}

	log.Fatal(autotls.RunWithManager(r, &m))
}

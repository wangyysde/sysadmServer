package main

import (
	"fmt"
	"log"

	"github.com/wangyysde/sysadmServer"
)

const (
	Addr = "127.0.0.1:2003"
)

func main() {
	r := sysadmServer.Default()
	r.GET("/:path", func(c *sysadmServer.Context) {
		// in this handler, we just simply send some basic info back to proxy response.
		req := c.Request
		urlPath := fmt.Sprintf("http://%s%s", Addr, req.URL.Path)
		realIP := fmt.Sprintf("RemoteAddr=%s,X-Forwarded-For=%v,X-Real-Ip=%v", req.RemoteAddr, req.Header.Get("X-Forwarded-For"), req.Header.Get("X-Real-Ip"))
		c.JSON(200, sysadmServer.H{
			"path": urlPath,
			"ip":   realIP,
		})
	})

	if err := r.Run(Addr); err != nil {
		log.Printf("Error: %v", err)
	}
}

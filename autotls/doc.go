// Package autotls support Let's Encrypt for a Go server application.
//
// 	package main
//
// 	import (
//
// 		"github.com/wangyysde/sysadmServer/autotls"
// 		"github.com/wangyysde/sysadmServer"
//		"github.com/wangyysde/sysadmLog"
// 	)
//
// 	func main() {
// 		r := sysadmServer.Default()
//
// 		// Ping handler
// 		r.GET("/ping", func(c *sysadmServer.Context) {
// 			c.String(200, "pong")
// 		})
//
//	    sysadmLog.Log(autotls.Run(r, "example1.com", "example2.com"),"fatal")
// 	}
//
package autotls

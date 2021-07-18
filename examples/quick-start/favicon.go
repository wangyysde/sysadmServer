package main

import (
	"net/http"
				
	"github.com/wangyysde/sysadmServer"
	"github.com/wangyysde/sysadmServer/favicon"
)

func main() {
	app := sysadmServer.Default()
	app.Use(favicon.New("./favicon.ico"))
	app.GET("/ping", func(c *sysadmServer.Context) {
					c.String(http.StatusOK, "Hello favicon.")
						})
	app.Run(":8080")
}

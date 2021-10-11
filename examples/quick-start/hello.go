package main

import (
	"net/http"
	"os"
	"fmt"

	"github.com/wangyysde/sysadmServer"
	"github.com/wangyysde/sysadmLog"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		sysadmServer.Log(fmt.Sprintf("Defaulting to port %s", port),"info")
	}

	// Starts a new sysadmServer instance with no middle-ware
	r := sysadmServer.New()

	// Define handlers
	r.GET("/", func(c *sysadmServer.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.GET("/ping", func(c *sysadmServer.Context) {
		c.String(http.StatusOK, "echo ping message")
	})

	// Listen and serve on defined port
	sysadmServer.Log(fmt.Sprintf("Listening on port %s", port),"info")
	r.Run(":" + port)
}

package routes

import (
	"net/http"

	"github.com/wangyysde/sysadmServer"
)

func addPingRoutes(rg *sysadmServer.RouterGroup) {
	ping := rg.Group("/ping")

	ping.GET("/", func(c *sysadmServer.Context) {
		c.JSON(http.StatusOK, "pong")
	})
}

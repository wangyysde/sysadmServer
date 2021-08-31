package routes

import (
	"net/http"

	"github.com/wangyysde/sysadmServer"
)

func addUserRoutes(rg *sysadmServer.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/", func(c *sysadmServer.Context) {
		c.JSON(http.StatusOK, "users")
	})
	users.GET("/comments", func(c *sysadmServer.Context) {
		c.JSON(http.StatusOK, "users comments")
	})
	users.GET("/pictures", func(c *sysadmServer.Context) {
		c.JSON(http.StatusOK, "users pictures")
	})
}

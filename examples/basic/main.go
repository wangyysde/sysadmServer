package main

import (
	"net/http"

	"github.com/wangyysde/sysadmServer"
)

var db = make(map[string]string)

func setupRouter() *sysadmServer.Engine {
	// Disable Console Color
	// sysadmServer.DisableConsoleColor()
	r := sysadmServer.Default()

	// Ping test
	r.GET("/ping", func(c *sysadmServer.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *sysadmServer.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, sysadmServer.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, sysadmServer.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses sysadmServer.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(sysadmServer.BasicAuth(sysadmServer.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", sysadmServer.BasicAuth(sysadmServer.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *sysadmServer.Context) {
		user := c.MustGet(sysadmServer.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, sysadmServer.H{"status": "ok", "authuserkey": sysadmServer.AuthUserKey, "user": user})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

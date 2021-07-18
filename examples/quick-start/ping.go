package main 

import "github.com/wangyysde/sysadmServer"

func main() {
    r := sysadmServer.Default()
    r.GET("/ping", func(c *sysadmServer.Context) {
        c.JSON(200, sysadmServer.H{
            "message": "ping",
        })   
    })   
    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

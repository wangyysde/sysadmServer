package main

import (
	"flag"
	"time"
	"fmt"

	"github.com/fatih/color"
	"github.com/wangyysde/sysadmServer"
	"go.uber.org/ratelimit"
	"github.com/wangyysde/sysadmLog"
)

var (
	limit ratelimit.Limiter
	rps   = flag.Int("rps", 5, "request per second")
)


func leakBucket() sysadmServer.HandlerFunc {
	prev := time.Now()
	return func(ctx *sysadmServer.Context) {
		now := limit.Take()
		sysadmLog.Log(fmt.Sprintf("%v",color.CyanString("%v", now.Sub(prev)),"info")
		prev = now
	}
}

func sysadmServerRun(rps int) {
	limit = ratelimit.New(rps)

	app := sysadmServer.Default()
	app.Use(leakBucket())

	app.GET("/rate", func(ctx *sysadmServer.Context) {
		ctx.JSON(200, "rate limiting test")
	})

	sysadmLog.Log(color.CyanString("Current Rate Limit: %v requests/s", rps))
	app.Run(":8080")
}

func main() {
	flag.Parse()
	sysadmServerRun(*rps)
}

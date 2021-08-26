package main

import (
	"flag"
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/wangyysde/sysadmServer"
	"go.uber.org/ratelimit"
)

var (
	limit ratelimit.Limiter
	rps   = flag.Int("rps", 5, "request per second")
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("[sysadmServer ")
	log.SetOutput(sysadmServer.DefaultWriter)
}

func leakBucket() sysadmServer.HandlerFunc {
	prev := time.Now()
	return func(ctx *sysadmServer.Context) {
		now := limit.Take()
		log.Print(color.CyanString("%v", now.Sub(prev)))
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

	log.Printf(color.CyanString("Current Rate Limit: %v requests/s", rps))
	app.Run(":8080")
}

func main() {
	flag.Parse()
	sysadmServerRun(*rps)
}

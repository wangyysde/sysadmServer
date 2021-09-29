package main

import (
	"net/http"
	"os"
	"fmt"

	"github.com/wangyysde/sysadmServer"
	"github.com/newrelic/go-agent"
	"github.com/wangyysde/sysadmLog"
)

const (
	// NewRelicTxnKey is the key used to retrieve the NewRelic Transaction from the context
	NewRelicTxnKey = "NewRelicTxnKey"
)

// NewRelicMonitoring is a middleware that starts a newrelic transaction, stores it in the context, then calls the next handler
func NewRelicMonitoring(app newrelic.Application) sysadmServer.HandlerFunc {
	return func(ctx *sysadmServer.Context) {
		txn := app.StartTransaction(ctx.Request.URL.Path, ctx.Writer, ctx.Request)
		defer txn.End()
		ctx.Set(NewRelicTxnKey, txn)
		ctx.Next()
	}
}

func main() {
	router := sysadmServer.Default()

	cfg := newrelic.NewConfig(os.Getenv("APP_NAME"), os.Getenv("NEW_RELIC_API_KEY"))
	app, err := newrelic.NewApplication(cfg)
	if err != nil {
		sysadmLog.Log(fmt.Sprintf("failed to make new_relic app: %v", err),"info")
	} else {
		router.Use(NewRelicMonitoring(app))
	}

	router.GET("/", func(c *sysadmServer.Context) {
		c.String(http.StatusOK, "Hello World!\n")
	})
	router.Run()
}

// +build go1.8

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/wangyysde/sysadmLog"
	"github.com/wangyysde/sysadmServer"
)

func main() {
	router := sysadmServer.Default()
	router.GET("/", func(c *sysadmServer.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome sysadmServer")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			sysadmLog.Log(fmt.Sprintf("listen: %s\n", err),"fatal")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	sysadmLog.Log("Shutting down server...","info")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		sysadmLog.Log(fmt.Sprintf("Server forced to shutdown:%s",err),"fatal")
	}
	
	sysadmLog.Log("Server exiting")
}

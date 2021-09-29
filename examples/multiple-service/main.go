package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/wangyysde/sysadmLog"
	"github.com/wangyysde/sysadmServer"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func router01() http.Handler {
	e := sysadmServer.New()
	e.Use(sysadmServer.Recovery())
	e.GET("/", func(c *sysadmServer.Context) {
		c.JSON(
			http.StatusOK,
			sysadmServer.H{
				"code":  http.StatusOK,
				"error": "Welcome server 01",
			},
		)
	})

	return e
}

func router02() http.Handler {
	e := sysadmServer.New()
	e.Use(sysadmServer.Recovery())
	e.GET("/", func(c *sysadmServer.Context) {
		c.JSON(
			http.StatusOK,
			sysadmServer.H{
				"code":  http.StatusOK,
				"error": "Welcome server 02",
			},
		)
	})

	return e
}

func main() {
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		sysadmLog.Log(fmt.Sprintf("%v", err),"fatal")
	}
}

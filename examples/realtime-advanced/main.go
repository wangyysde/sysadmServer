package main

import (
	"fmt"
	"runtime"

	"github.com/wangyysde/sysadmServer"
)

func main() {
	ConfigRuntime()
	StartWorkers()
	StartsysadmServer()
}

// ConfigRuntime sets the number of operating system threads.
func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

// StartWorkers start starsWorker by goroutine.
func StartWorkers() {
	go statsWorker()
}

// StartGin starts sysadmServer web server with setting router.
func StartsysadmServer() {
	sysadmServer.SetMode(sysadmServer.ReleaseMode)

	router := sysadmServer.New()
	router.Use(rateLimit, sysadmServer.Recovery())
	router.LoadHTMLGlob("resources/*.templ.html")
	router.Static("/static", "resources/static")
	router.GET("/", index)
	router.GET("/room/:roomid", roomGET)
	router.POST("/room-post/:roomid", roomPOST)
	router.GET("/stream/:roomid", streamRoom)

	router.Run(":80")
}

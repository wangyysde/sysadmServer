package main

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/wangyysde/sysadmServer"
)

//go:embed assets/* templates/*
var f embed.FS

func main() {
	
	router := sysadmServer.Default()
	templ := template.Must(template.New("").ParseFS(f, "templates/*.tmpl", "templates/foo/*.tmpl"))
	router.SetHTMLTemplate(templ)

	// example: /public/assets/images/example.png
	router.StaticFS("/public", http.FS(f))

	router.GET("/", func(c *sysadmServer.Context) {
		c.HTML(http.StatusOK, "index.tmpl", sysadmServer.H{
			"title": "Main website",
		})
	})

	router.GET("/foo", func(c *sysadmServer.Context) {
		c.HTML(http.StatusOK, "bar.tmpl", sysadmServer.H{
			"title": "Foo website",
		})
	})

	router.GET("favicon.ico", func(c *sysadmServer.Context) {
		file, _ := f.ReadFile("assets/favicon.ico")
		c.Data(
			http.StatusOK,
			"image/x-icon",
			file,
		)
	})

	router.Run(":8080")
}

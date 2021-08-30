package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/wangyysde/sysadmServer"
)

func main() {
	r := sysadmServer.New()
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)
	r.GET("/", func(c *sysadmServer.Context) {
		c.HTML(http.StatusOK, "/html/index.tmpl", sysadmServer.H{
			"Foo": "World",
		})
	})
	r.GET("/bar", func(c *sysadmServer.Context) {
		c.HTML(http.StatusOK, "/html/bar.tmpl", sysadmServer.H{
			"Bar": "World",
		})
	})
	r.Run(":8080")
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

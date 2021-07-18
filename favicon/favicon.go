package favicon

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/wangyysde/sysadmServer"
)

func New(path string) sysadmServer.HandlerFunc {
	path = filepath.FromSlash(path)
	if len(path) > 0 && !os.IsPathSeparator(path[0]) {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		path = filepath.Join(wd, path)
	}

	info, err := os.Stat(path)
	if err != nil || info == nil || info.IsDir() {
		panic("Invalid favicon path: " + path)
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(file)

	return func(c *sysadmServer.Context) {
		if c.Request.RequestURI != "/favicon.ico" {
			return
		}
		if c.Request.Method != "GET" && c.Request.Method != "HEAD" {
			status := http.StatusOK
			if c.Request.Method != "OPTIONS" {
				status = http.StatusMethodNotAllowed
			}
			c.Header("Allow", "GET,HEAD,OPTIONS")
			c.AbortWithStatus(status)
			return
		}
		c.Header("Content-Type", "image/x-icon")
		http.ServeContent(c.Writer, c.Request, "favicon.ico", info.ModTime(), reader)
		return
	}
}

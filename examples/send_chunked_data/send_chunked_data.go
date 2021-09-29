package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/wangyysde/sysadmServer"
)

func main(){
	r := sysadmServer.Default()
	r.GET("/test_stream", func(c *sysadmServer.Context){
		w := c.Writer
		header := w.Header()
		header.Set("Transfer-Encoding", "chunked")
		header.Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
			<html>
					<body>
		`))
		w.(http.Flusher).Flush()
		for i:=0 ;i<10; i++{
			w.Write([]byte(fmt.Sprintf(`
				<h1>%d</h1>
			`,i)))
			w.(http.Flusher).Flush()
			time.Sleep(time.Duration(1) * time.Second)
		}
		w.Write([]byte(`
			
					</body>
			</html>
		`))
		w.(http.Flusher).Flush()
	})

	r.Run(":8080")
}

/*
browser test url:
http://127.0.0.1:8080/test_stream
*/

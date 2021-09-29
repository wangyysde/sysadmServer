package main

import (
	"fmt"
	"html/template"

	"github.com/wangyysde/sysadmLog"
	"github.com/wangyysde/sysadmServer"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, sysadmServer!</h1>
</body>
</html>
`))

func main() {
	r := sysadmServer.Default()
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html)

	r.GET("/", func(c *sysadmServer.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// use pusher.Push() to do server push
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				sysadmLog.Log(fmt.Sprintf("Failed to push: %v",err),"info")
			}
		}
		c.HTML(200, "https", sysadmServer.H{
			"status": "success",
		})
	})

	// Listen and Server in https://127.0.0.1:8080
	r.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}

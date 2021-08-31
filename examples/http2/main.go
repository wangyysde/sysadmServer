package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/wangyysde/sysadmServer"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
</head>
<body>
  <h1 style="color:red;">Welcome!</h1>
</body>
</html>
`))

func main() {
	logger := log.New(os.Stderr, "", 0)
	logger.Println("[WARNING] DON'T USE THE EMBED CERTS FROM THIS EXAMPLE IN PRODUCTION ENVIRONMENT, GENERATE YOUR OWN!")

	r := sysadmServer.Default()
	r.SetHTMLTemplate(html)

	r.GET("/welcome", func(c *sysadmServer.Context) {
		c.HTML(http.StatusOK, "https", sysadmServer.H{
			"status": "success",
		})
	})

	// Listen and Server in https://127.0.0.1:8080
	r.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}

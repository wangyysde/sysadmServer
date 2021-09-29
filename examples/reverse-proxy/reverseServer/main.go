package main

import (
	"bufio"
	"net/http"
	"net/url"
	"fmt"

	"github.com/wangyysde/sysadmServer"
	"github.com/wangyysde/sysadmLog"
)

const (
	// this is our reverse server ip address
	ReverseServerAddr = "127.0.0.1:2002"
)

var (
	// maybe we can have many real server addresses and do some load balanced strategy.
	RealAddr = []string{
		"http://127.0.0.1:2003",
	}
)

// a fake function that we can do strategy here.
func getLoadBalanceAddr() string {
	return RealAddr[0]
}

func main() {
	r := sysadmServer.Default()
	r.GET("/:path", func(c *sysadmServer.Context) {
		// step 1: resolve proxy address, change scheme and host in requets
		req := c.Request
		proxy, err := url.Parse(getLoadBalanceAddr())
		if err != nil {
			sysadmLog.Log(fmt.Sprintf("error in parse addr: %v", err),"info")
			c.String(500, "error")
			return
		}
		req.URL.Scheme = proxy.Scheme
		req.URL.Host = proxy.Host

		// step 2: use http.Transport to do request to real server.
		transport := http.DefaultTransport
		resp, err := transport.RoundTrip(req)
		if err != nil {
			sysadmLog.Log(fmt.Sprintf("error in roundtrip: %v", resp),"info")
			c.String(500, "error")
			return
		}

		// step 3: return real server response to upstream.
		for k, vv := range resp.Header {
			for _, v := range vv {
				c.Header(k, v)
			}
		}
		defer resp.Body.Close()
		bufio.NewReader(resp.Body).WriteTo(c.Writer)
		return
	})

	if err := r.Run(ReverseServerAddr); err != nil {
		sysadmLog.Log(fmt.Sprintf("Error: %v", err),"error")
	}
}

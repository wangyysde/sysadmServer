package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wangyysde/sysadmServer"
	pb "github.com/wangyysde/sysadmServer/examples/grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	// Set up a http server.
	r := sysadmServer.Default()
	r.GET("/rest/n/:name", func(c *sysadmServer.Context) {
		name := c.Param("name")

		// Contact the server and print out its response.
		req := &pb.HelloRequest{Name: name}
		res, err := client.SayHello(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, sysadmServer.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, sysadmServer.H{
			"result": fmt.Sprint(res.Message),
		})
	})

	// Run http server
	if err := r.Run(":8052"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

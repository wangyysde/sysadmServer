package main

import (
	"fmt"
	"net"

	"github.com/wangyysde/sysadmLog"
	pb "github.com/wangyysde/sysadmServer/examples/grpc/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		sysadmLog.Log(fmt.Sprintf("failed to listen: %v", err),"fatal")
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		sysadmLog.Log(fmt.Sprintf("Failed to serve: %s", err),"fatal")
	}
}

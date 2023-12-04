package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/akarshippili/networking/gRpc/helloworld"
	"google.golang.org/grpc"
)

type Server struct {
	helloworld.UnimplementedGreeterServer
}

func (server *Server) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received %v", request.GetName())

	return &helloworld.HelloReply{
		Message: "Yo yo yo. 148-3 to the 3 to the 6 to the 9, representing the ABQ, what up, biatch?!",
	}, nil
}

func main() {
	port := flag.Int("port", 5000, "Server Port")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &Server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

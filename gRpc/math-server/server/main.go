package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"

	math_server "github.com/akarshippili/networking/gRpc/math-server"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 50051, "The server port")

type Server struct {
	math_server.UnimplementedMathServer
}

func (server *Server) Add(ctx context.Context, req *math_server.Request) (*math_server.Response, error) {
	return &math_server.Response{Result: req.Var1 + req.Var2}, nil
}

func (server *Server) Sub(ctx context.Context, req *math_server.Request) (*math_server.Response, error) {
	return &math_server.Response{Result: req.Var1 - req.Var2}, nil
}

func (server *Server) Mul(ctx context.Context, req *math_server.Request) (*math_server.Response, error) {
	return &math_server.Response{Result: req.Var1 * req.Var2}, nil
}

func (server *Server) Div(ctx context.Context, req *math_server.Request) (*math_server.Response, error) {
	if req.Var2 == 0 {
		return nil, errors.New("arithmetic exception: divide by zero")
	}
	return &math_server.Response{Result: req.Var1 / req.Var2}, nil
}

func main() {
	flag.Parse()

	// opening a tcp listener
	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// registering server
	s := grpc.NewServer()
	math_server.RegisterMathServer(s, &Server{})

	err = s.Serve(l)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

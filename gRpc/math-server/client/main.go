package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	math_server "github.com/akarshippili/networking/gRpc/math-server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var port = flag.Int("port", 50051, "The server port")
var var1 = flag.Int64("var1", 1, "value for var1")
var var2 = flag.Int64("var2", 1, "value for var2")

func main() {
	flag.Parse()

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", *port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := math_server.NewMathClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Add(ctx, &math_server.Request{Var1: *var1, Var2: *var2})
	if err != nil {
		log.Fatalf("error while making a reqest, cause: %v", err)
	}

	log.Printf("response: %d", res.GetResult())
}

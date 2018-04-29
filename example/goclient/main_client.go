package main

import (
	"context"
	"log"
	"time"

	"github.com/goboilerplates/micro-rpc/example/goclient/config"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50052"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	config.SetRequesters(ctx, conn)
}

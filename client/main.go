package main

import (
	"context"
	"grpc-load-balance-test/pb"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	host := "echo.test:80"
	if len(os.Args) > 1 {
		host = os.Args[1]
	}

	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewPingerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	// Ping every 10 ms
	for {
		_, err := c.Ping(ctx, &pb.Empty{})
		if err != nil {
			log.Fatalf("did not ping: %v", err)
		}

		time.Sleep(10 * time.Millisecond)
	}
}

package main

import (
	"context"
	"log"
	"net"

	"grpc-load-balance-test/pb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPingerServer
	counter int64
}

func (s *server) Ping(_ context.Context, _ *pb.Empty) (*pb.Empty, error) {
	s.counter++

	log.Printf("count of pings: %d", s.counter)

	return &pb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPingerServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

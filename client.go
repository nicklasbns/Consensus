package main

import (
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	s := &server{
		participants: make(map[string]int64),
		messages:     []*pb.BroadcastMessage{},
		clients:      make(map[string]chan *pb.BroadcastMessage), // Initialize the map for client channels
	}
	pb.RegisterChittyChatServer(grpcServer, s)
}

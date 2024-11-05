package main

import (
	pb "Consensus/consensus"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedConsensusServer
}

func main() {
	conn, err := grpc.NewClient("nicklas-laptop.taild2c7d.ts.net:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewConsensusClient(conn)

	lis, err := net.Listen("tcp", "127.0.0.2:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	s := &server{}
	pb.RegisterConsensusServer(grpcServer, s)

	fmt.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

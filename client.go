package main

import (
	pb "Consensus/consensus"
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedConsensusServer
	address int
	client  pb.ConsensusClient
}

func main() {
	var lis net.Listener
	var err error
	address := 2
	var addrString string

	// server stuff
	for {
		addrString = "127.0.0." + strconv.Itoa(address) + ":50051"
		lis, err = net.Listen("tcp", addrString)
		if err != nil {
			address++
			continue
		}
		break
	}

	grpcServer := grpc.NewServer()
	s := &server{
		address: address,
	}
	pb.RegisterConsensusServer(grpcServer, s)

	fmt.Println("Server is running on address", addrString)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func (s *server) StartFunction(ctx context.Context, empty *pb.Empty) (*pb.Empty, error) {
	log.Println("Server " + strconv.Itoa(s.address) + " started")

	connectAddress := "127.0.0." + strconv.Itoa(s.address+1) + ":50051"
	conn, err := grpc.NewClient(connectAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	s.client = pb.NewConsensusClient(conn)
	s.client.StartFunction(context.Background(), &pb.Empty{})
	log.Println("Connect Success to address" + strconv.Itoa(s.address+1))
	return &pb.Empty{
		message: "Server " + strconv.Itoa(s.address) + " started",
	}, nil
}

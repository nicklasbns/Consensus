package main

import (
	pb "Consensus/consensus"
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedConsensusServer
}

func main() {
	var lis net.Listener
	var err error
	address := 2
	var addrString string
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
	s := &server{}
	pb.RegisterConsensusServer(grpcServer, s)

	fmt.Println("Server is running on address", addrString)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func (s *server) SendMessage(ctx context.Context, empty *pb.Empty) (*pb.Empty, error) {
	log.Println("New Message")
	return &pb.Empty{}, nil
}

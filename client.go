package main

import (
	pb "Consensus/consensus"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedConsensusServer
	address       int
	targetAddress int
	started       bool
	client        pb.ConsensusClient
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

func (s *server) PassToken(ctx context.Context, empty *pb.Token) (*pb.Empty, error) {
	log.Println("Token recieved")
	writeToFile()

	s.client.PassToken(context.Background(), &pb.Token{})

	return &pb.Empty{}, nil
}

func (s *server) StartFunction(ctx context.Context, empty *pb.Empty) (*pb.SuccessStart, error) {
	if s.started {
		log.Println("Server already started")
		return &pb.SuccessStart{
			Message: "Server " + strconv.Itoa(s.address) + " started",
		}, nil
	}
	s.started = true
	log.Println("Server " + strconv.Itoa(s.address) + " started")

	s.targetAddress = s.address + 1
	message, client := connectToServer(s.targetAddress)
	if message == nil {
		s.targetAddress = 2
		message, client = connectToServer(s.targetAddress)

	}
	s.client = client

	log.Println("Connect Success to address" + strconv.Itoa(s.targetAddress))
	return &pb.SuccessStart{
		Message: "Server " + strconv.Itoa(s.address) + " started",
	}, nil
}

func connectToServer(address int) (*pb.SuccessStart, pb.ConsensusClient) {
	connectAddress := "127.0.0." + strconv.Itoa(address) + ":50051"
	conn, err := grpc.NewClient(connectAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewConsensusClient(conn)
	message, err := client.StartFunction(context.Background(), &pb.Empty{})
	return message, client

}

func writeToFile() {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	data := []byte(currentTime + " HAHA")
	err := os.WriteFile("CitricalSection.txt", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Second)
}

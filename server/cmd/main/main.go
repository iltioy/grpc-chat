package main

import (
	"log"
	"net"

	pb "github.com/iltioy/server/proto"
	"github.com/iltioy/server/services"
	"google.golang.org/grpc"
)

const (
	port = ":8000"
)


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

	grpcServer := grpc.NewServer()

	service := &services.ChatServer{}
	pb.RegisterChatServiceServer(grpcServer, service)
	log.Printf("Server started on port %v", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
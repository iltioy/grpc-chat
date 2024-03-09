package services

import (
	"context"
	"fmt"

	pb "github.com/iltioy/server/proto"
)

type ChatServer struct {
	pb.ChatServiceServer
}


func (srv *ChatServer) ChatInitiate(ctx context.Context, req *pb.InitiateRequest) (*pb.InitiateResponse, error) {
	fmt.Print("fsdfsdf")
	return nil, nil
} 
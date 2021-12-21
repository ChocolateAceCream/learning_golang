package main

import (
	"context"
	"log"
	"net"
	"tutorail/pb"

	"google.golang.org/grpc"
)

const (
	port = ":50001"
)

type server struct {
	pb.UnimplementedMessageSenderServer
}

func (s *server) MessageSender(ctx context.Context, in *pb.User) (*pb.Reply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.Reply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMessageSenderServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

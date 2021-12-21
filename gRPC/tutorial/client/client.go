package main

import (
	"context"
	"log"
	"os"
	"time"
	"tutorail/pb"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50001"
	defaultname = "nuodi"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMessageSenderClient(conn)

	name := defaultname
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Send(ctx, &pb.User{Name: name})
	if err != nil {
		log.Fatalf("failed to send request: %v", err)
	}
	log.Printf("res: %v", r.GetMessage())
}

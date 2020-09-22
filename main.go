package main

import (
	"context"
	"fmt"
	"log"
	pb "playlistclient/proto"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {
	fmt.Println("try to connect")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("no; %v", err)
	}

	defer conn.Close()
	c := pb.NewPlaylistManagerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.NewPlaylist(ctx, &pb.NewPlaylistRequest{Name: "my playlist"})
	if err != nil {
		log.Fatalf("uhoh: %v", err)
	}

	log.Printf("Greeting: %s", r.GetMessage())
}

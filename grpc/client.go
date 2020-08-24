package main

import (
	"context"
	"go-edge/grpc/chat"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	response, err := c.BroadcastMessage(context.Background(), &chat.Message{Body: "Hello From CLient! BroadcastMessage"})
	if err != nil {
		log.Fatalf("Error when calling BroadcastMessage: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
	response, err = c.SayHello(context.Background(), &chat.Message{Body: "Hello From CLient!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}

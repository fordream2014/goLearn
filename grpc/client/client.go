package main

import (
	"google.golang.org/grpc"
	"log"
	"goLearn/grpc/proto"
	"os"
	"context"
)

const (
	address = "127.0.0.1:10086"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("failed to listen: %v", err)
		return
	}
	defer conn.Close()

	c := proto.NewHelloServiceClient(conn)
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	resp, err := c.SayHello(context.Background(), &proto.HelloRequest{Greeting: name})
	if err != nil {
		log.Fatal("could not greet: %v", err)
		return
	}

	log.Printf("Greeting: %s", resp.Reply)
}

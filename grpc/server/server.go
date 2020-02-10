package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"goLearn/grpc/proto"
	"context"
)

const (
	port = ":10086"
)

type Server struct {}

func (s *Server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {
	log.Printf("request message: %v \n", in.Greeting)
	return &proto.HelloResponse{
		Reply: "hello, " + in.Greeting,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	proto.RegisterHelloServiceServer(s, &Server{})
	s.Serve(lis)
}

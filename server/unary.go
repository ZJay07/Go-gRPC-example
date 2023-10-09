package main

import (
	"context"
	pb "github.com/ZJay07/Go-gRPC-example/proto"
)

func(s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error){
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}

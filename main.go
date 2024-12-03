package main

import (
	"github.com/zhang121923/rpc-proto-registry/app"
	proto "github.com/zhang121923/rpc-proto-registry/protos/go"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterCreateUserServer(s, &app.UserServer{})
	proto.RegisterStudentServiceServer(s, &app.StudentServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

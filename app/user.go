package app

import (
	"context"
	pb "github.com/zhang121923/rpc-proto-registry/protos/go"
	"log"
)

type UserServer struct {
	pb.UnimplementedCreateUserServer
}

func (u *UserServer) GetUserName(ctx context.Context, usr *pb.User) (*pb.UserNameResp, error) {
	log.Print("call success.")
	return &pb.UserNameResp{Name: usr.Name}, nil
}

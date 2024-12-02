package app

import (
	"context"
	pb "github.com/zhang121923/rpc-proto-registry/protos/go"
)

type UserServer struct {
	pb.UnimplementedCreateUserServer
}

func (u *UserServer) GetUserName(context.Context, *pb.User) (*pb.UserNameResp, error) {
	return &pb.UserNameResp{
		Name: "zhangsan",
	}, nil
}

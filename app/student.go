package app

import (
	"context"
	pb "github.com/zhang121923/rpc-proto-registry/protos/go"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type StudentServer struct {
	pb.UnimplementedStudentServiceServer
}

func (s *StudentServer) JoinSchool(ctx context.Context, stu *pb.Student) (*emptypb.Empty, error) {
	stu.SchoolName = "hahah school happy"
	log.Print("success join school")
	return &emptypb.Empty{}, nil
}

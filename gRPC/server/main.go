package main

import (
	"context"
	"log"
	"net"

	"github.com/gregerssr/DISYS_MandatoryExercise1/gRPC/api"
	pf "github.com/gregerssr/DISYS_MandatoryExercise1/gRPC/protofiles"
	"google.golang.org/grpc"
)

// Must implement the CourseServer interface
type Server struct {
	pf.UnimplementedCourseServer

	Repo api.Repo
}

const (
	port = ":50051"
)

func (s *Server) FindCourse(ctx context.Context, in *pf.CourseRequest) (*pf.CourseReply, error) {
	log.Printf("Recieved %v", in.GetId())
	return &pf.CourseReply{Name: s.Repo.FindCourse(in.GetId()).Title}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pf.RegisterCourseServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

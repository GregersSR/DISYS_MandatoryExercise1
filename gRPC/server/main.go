package gRPC

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Must implement the CourseServer interface
type Server struct {
	UnimplementedCourseServer

	Repo Repo
}

const (
	port = ":50051"
)

func (s *Server) FindCourse(ctx context.Context, in *CourseRequest) (*CourseReply, error) {
	log.Printf("Recieved %v", in.GetId())
	return &CourseReply{Name: s.Repo.FindCourse(in.GetId()).Title}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterCourseServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

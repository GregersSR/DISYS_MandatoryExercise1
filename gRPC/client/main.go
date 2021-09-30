package main

import (
	"context"
	"log"
	"time"

	pf "github.com/gregerssr/DISYS_MandatoryExercise1/gRPC/protofiles"
	"google.golang.org/grpc"
)

const (
	address       = "localhost:50051"
	defaultCourse = 1
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pf.NewCourseClient(conn)

	course := defaultCourse

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.FindCourse(ctx, &pf.CourseRequest{Id: int64(course)})
	if err != nil {
		log.Fatalf("Could not find the course: %v", err)
	}

	log.Printf("Course title: %s", r.GetName())
}

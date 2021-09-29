package main

import (
	"context"
	"fmt"
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
	c := pf.NewCourseServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Add a course to the server
	addCoursePayload := pf.Course{Id: 1, Title: "DISYS", Teachers: []int64{1, 2}, Ects: 75, SatisfactionScore: 42}
	addCourseResponse, e1 := c.AddCourse(ctx, &addCoursePayload)
	if e1 != nil {
		log.Fatalf(e1.Error())
	}
	log.Printf("%v", addCourseResponse.IsSuccess)

	// Get a course by ID on the server
	getCoursePayload := pf.CourseID{Id: 1}
	getCourseResponse, e2 := c.GetCourse(ctx, &getCoursePayload)
	if e2 != nil {
		log.Fatalf(e1.Error())
	}
	log.Printf("%s", toString(getCourseResponse))

	// Update a course on the server
	updateCoursePayload := pf.Course{Id: 1, Title: "BDSA", Teachers: []int64{3, 4}, Ects: 150, SatisfactionScore: 69}
	updateCourseRepsponse, e3 := c.UpdateCourse(ctx, &updateCoursePayload)
	if e3 != nil {
		log.Fatalf(e1.Error())
	}
	log.Printf("%v", updateCourseRepsponse.IsSuccess)

	// Get teachers of a course
	getCourseTeachersPayload := pf.CourseID{Id: 1}
	getCourseTeachersResponse, e4 := c.GetCourseTeachers(ctx, &getCourseTeachersPayload)
	if e4 != nil {
		log.Fatalf(e1.Error())
	}
	log.Printf("%v", getCourseTeachersResponse.Teachers)

	// Get satisfaction rating of a course
	getSatisfactionRatingPayload := pf.CourseID{Id: 1}
	getSatisfactionRatingResponse, e5 := c.GetCourseSatisfactionRating(ctx, &getSatisfactionRatingPayload)
	if e5 != nil {
		log.Fatalf(e1.Error())
	}
	log.Printf("%v", getSatisfactionRatingResponse.SatisfactionRating)

	// Delete the course on the server
	deleteCoursePayload := pf.CourseID{Id: 1}
	deleteCourseResponse, e6 := c.DeleteCourse(ctx, &deleteCoursePayload)
	if e6 != nil {
		log.Fatalf(e1.Error())
	}
	log.Printf("%v", deleteCourseResponse.IsSuccess)
}

func toString(course *pf.Course) string {
	return fmt.Sprintf("%v: %s. Taught by: %v, Satisfaction Rating: %v", course.Id, course.Title, course.Teachers, course.SatisfactionScore)
}

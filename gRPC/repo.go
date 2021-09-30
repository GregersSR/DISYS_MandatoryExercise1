package gRPC

type Course struct {
	ID                int64
	Title             string
	Teachers          []int64
	ECTS              int
	SatisfactionStore int
}

type Repo struct {
	Courses map[int64]Course
}

func (r *Repo) Init() {
	r.Courses = make(map[int64]Course)
}

func (r *Repo) FindCourse(id int64) Course {
	return r.Courses[id]
}

package api

type Course struct {
	ID                int64
	Title             string
	Teachers          []int64
	ECTS              int
	SatisfactionScore int
}

type Repo struct {
	Courses map[int64]Course
}

func (r *Repo) Init() {
	r.Courses = make(map[int64]Course)

	course := Course{ID: 1, Title: "DISYS", Teachers: []int64{1, 2}, ECTS: 10, SatisfactionScore: 10}
	r.Courses[course.ID] = course
}

func (r *Repo) FindCourse(id int64) Course {
	return r.Courses[id]
}

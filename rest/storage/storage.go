package storage

import (
	"errors"

	"github.com/gregerssr/DISYS_MandatoryExercise1/rest/model"
)

var courses map[int64]model.Course

func Init() {
	courses = make(map[int64]model.Course)
}
func InsertCourse(c model.Course) bool {
	if _, inMap := courses[c.Id]; !inMap {
		courses[c.Id] = c
		return true
	}
	return false
}

func GetCourse(id int64) (model.Course, error) {
	course, inMap := courses[id]
	if inMap {
		return course, nil
	} else {
		return course, errors.New("Course doesn't exist")
	}
}

func UpdateCourse(c model.Course) bool {
	if _, inMap := courses[c.Id]; inMap {
		courses[c.Id] = c
		return true
	}
	return false
}
func DeleteCourse(id int64) bool {
	if _, inMap := courses[id]; inMap {
		delete(courses, id)
		return true
	}
	return false
}

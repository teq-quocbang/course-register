package presenter

import "github.com/teq-quocbang/course-register/model"

type CourseResponseWrapper struct {
	Course model.Course `json:"course"`
}

type ListCourseResponseWrapper struct {
	Course []model.Course `json:"course"`
	Meta   interface{}    `json:"meta"`
}

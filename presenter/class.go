package presenter

import "github.com/teq-quocbang/course-register/model"

type ClassResponseWrapper struct {
	Class model.Class `json:"class"`
}

type ListClassResponseWrapper struct {
	Class []model.Class `json:"class"`
	Meta  interface{}   `json:"meta"`
}

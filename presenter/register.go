package presenter

import "github.com/teq-quocbang/course-register/model"

type RegisterResponseWrapper struct {
	Register model.Register `json:"register"`
}

type ListRegisterResponseWrapper struct {
	Register []model.Register `json:"register"`
	Meta     interface{}      `json:"meta"`
}

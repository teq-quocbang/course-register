package presenter

import "github.com/teq-quocbang/course-register/model"

type AccountResponseWrapper struct {
	Account *model.Account `json:"account"`
}

type ListAccountResponseWrapper struct {
	Account []model.Account `json:"accounts"`
	Meta    interface{}     `json:"meta"`
}

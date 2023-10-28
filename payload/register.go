package payload

import "github.com/go-playground/validator/v10"

type CreateRegisterRequest struct {
	SemesterID string `json:"semester_id" validate:"required"`
	ClassID    string `json:"class_id" validate:"required"`
	CourseID   string `json:"course_id" validate:"required"`
}

func (r *CreateRegisterRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

type ListSemesterInformationRequest struct {
	SemesterID string `json:"semester_id" validate:"required"`
}

func (s *ListSemesterInformationRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}

type ListRegisteredHistories struct {
	SemesterID string `json:"semester_id" validate:"required"`
}

func (r *ListRegisteredHistories) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

type UnRegisterRequest struct {
	SemesterID string `json:"semester_id" validate:"required"`
	ClassID    string `json:"class_id" validate:"required"`
	CourseID   string `json:"course_id" validate:"required"`
}

func (u *UnRegisterRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

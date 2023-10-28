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

type ListRegisterHistories struct {
	SemesterID string `json:"semester_id" validate:"required"`
}

func (r *ListRegisterHistories) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

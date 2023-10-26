package payload

import (
	"github.com/go-playground/validator/v10"
)

type CreateSemesterRequest struct {
	ID                string `json:"ID" validate:"required"`
	MinCredits        int    `json:"min_credits" validate:"required,min=1"`
	StartTime         string `json:"start_time" validate:"required"`
	EndTime           string `json:"end_time" validate:"required"`
	RegisterStartAt   string `json:"register_start_at" validate:"required"`
	RegisterExpiresAt string `json:"register_expires_at" validate:"required"`
}

func (s *CreateSemesterRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}

type GetListSemesterRequest struct {
	Year string `json:"year" validate:"required"`
}

func (s *GetListSemesterRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}

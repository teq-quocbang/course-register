package model

import "time"

type Course struct {
	ID         string    `json:"id"`
	SemesterID string    `json:"semester_id"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  *uint     `json:"created_by"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdatedBy  *uint     `json:"updated_by"`
}

func (Course) TableName() string {
	return "course"
}

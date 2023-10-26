package model

import "time"

type Class struct {
	ID          string    `json:"id"`
	CourseID    string    `json:"course_id"`
	SemesterID  string    `json:"semester_id"`
	Credits     uint      `json:"creates"`
	MaxSlot     uint      `json:"max_slot"`
	CurrentSlot uint      `json:"current_slot"`
	CanCancel   bool      `json:"can_cancel"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   *uint     `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   *uint     `json:"updated_by"`
}

func (Class) TableName() string {
	return "class"
}

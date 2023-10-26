package model

import "time"

type Semester struct {
	ID                string    `json:"id"`
	MinCredits        int       `json:"min_credits"`
	StartTime         time.Time `json:"start_time"`
	EndTime           time.Time `json:"end_time"`
	RegisterStartAt   time.Time `json:"register_start_at"`
	RegisterExpiresAt time.Time `json:"register_expires_at"`
	CreatedAt         time.Time `json:"created_at"`
	CreatedBy         *uint     `json:"created_by"`
	UpdatedAt         time.Time `json:"updated_at"`
	UpdatedBy         *uint     `json:"updated_by"`
}

func (Semester) TableName() string {
	return "semester"
}

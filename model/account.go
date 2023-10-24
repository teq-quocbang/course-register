package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Account struct {
	StudentID    uint      `json:"student_id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	HashPassword []byte    `json:"hash_password"`
	IsVerified   bool      `json:"is_verified"`
	RegisteredAt time.Time `json:"registered_at"`
}

func (Account) TableName() string {
	return "account"
}

type AccountVerify struct {
	StudentID  uint      `json:"student_id"`
	SecretCode uuid.UUID `json:"secret_code"`
	IsUsed     bool      `json:"is_used"`
	CreatedAt  time.Time `json:"created_at"`
}

func (AccountVerify) TableName() string {
	return "account_verify"
}

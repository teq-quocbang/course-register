package payload

type CreateAccountRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type CreateAccountResponse struct {
	StudentID uint `json:"student_id"`
}

func (a *CreateAccountRequest) Validate() error {
	var validate 
}
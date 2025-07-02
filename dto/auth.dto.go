package dto;

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,eqfield=Password"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type ForgotPasswordRequest struct {
	Email    string `json:"email" binding:"required,email"`
	VerificationCode string `json:"code" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required,min=6"`
	ConfirmNewPassword string `json:"confirmNewPassword" binding:"required,eqfield=NewPassword"`
}

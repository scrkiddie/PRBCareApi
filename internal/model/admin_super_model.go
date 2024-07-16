package model

type AdminSuperResponse struct {
	Token string
}
type AdminSuperLoginRequest struct {
	Username string `json:"username" validate:"required,min=6,max=50,not_contain_space"`
	Password string `json:"password" validate:"required,min=6,max=255,is_password_format,not_contain_space"`
}
type AdminSuperPasswordUpdateRequest struct {
	ID              int    `validate:"required,numeric"`
	CurrentPassword string `json:"currentPassword" validate:"required,min=6,max=255,is_password_format,not_contain_space"`
	NewPassword     string `json:"newPassword" validate:"required,min=6,max=255,is_password_format,not_contain_space"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=6,max=255,eqfield=NewPassword"`
}
type VerifyAdminSuperRequest struct {
	Token string `validate:"required"`
}

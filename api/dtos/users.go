package dtos

type NewUserInput struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	DisplayName     string `json:"display_name"`
}

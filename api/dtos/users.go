package dtos

type NewUserInput struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	DisplayName     string `json:"display_name"`
}

type UserUpdate struct {
	Email           *string `json:"email,omitempty"`
	DisplayName     *string `json:"display_name,omitempty"`
	Password        *string `json:"password,omitempty"`
	ConfirmPassword *string `json:"confirm_password,omitempyt"`
}

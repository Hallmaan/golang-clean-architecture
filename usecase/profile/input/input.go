package profile_input

type CreateProfileInput struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	PhoneNumber  int `json:"phoneNumber" binding:"required"`
}
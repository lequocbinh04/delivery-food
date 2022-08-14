package usermodel

type UserCreate struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password"  binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

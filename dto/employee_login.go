package dto

type EmployeeLoginDTO struct {
	Username string `json:"username" binding:"required"` // Username is required for login
	Password string `json:"password" binding:"required"` // Password is required for login
}

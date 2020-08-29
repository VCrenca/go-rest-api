package dto

// LoginRequest -
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUserRequest -
type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUserResponse -
type CreateUserResponse struct {
	ID string `json:"id"`
}

// UserResponse -
type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

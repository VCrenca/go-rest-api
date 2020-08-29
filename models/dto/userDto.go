package dto

// CreateUserRequest -
type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUserResponse -
type CreateUserResponse struct {
	ID string `json:"id"`
}

// GetUserByIDResponse -
type GetUserByIDResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

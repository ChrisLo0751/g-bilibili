package dto

type UserResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
}

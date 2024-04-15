package dto

type UserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateUserRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

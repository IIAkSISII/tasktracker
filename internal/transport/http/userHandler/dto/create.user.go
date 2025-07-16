package dto

type CreateUserRequest struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Id      int64  `json:"id" example:"42"`
	Message string `json:"message" example:"user created"`
}

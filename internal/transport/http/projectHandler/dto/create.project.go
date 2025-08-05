package dto

import "time"

type CreateProjectRequest struct {
	Name      string    `json:"name"`
	IsPublic  bool      `json:"is_public"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateProjectResponse struct {
	Id      int    `json:"id" example:"42"`
	Message string `json:"message" example:"project created"`
}

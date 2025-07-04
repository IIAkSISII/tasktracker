package models

import "time"

type Board struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	ProjectId int       `json:"project_id"`
}

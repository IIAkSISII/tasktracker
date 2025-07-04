package models

type Member struct {
	UserId    int `json:"user_id"`
	ProjectId int `json:"project_id"`
	RoleId    int `json:"role_id"`
}

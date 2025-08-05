package dto

type CreateBoardRequest struct {
	Name      string `json:"name"`
	ProjectId int    `json:"project_id"`
}

type CreateBoardResponse struct {
	Id      int    `json:"id" example:"42"`
	Message string `json:"message" example:"board created"`
}

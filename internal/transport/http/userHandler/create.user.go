package userHandler

import (
	"context"
	"encoding/json"
	"github.com/IIAkSISII/tasktracker/internal/logger"
	"github.com/IIAkSISII/tasktracker/internal/transport/http/userHandler/dto"
	"net/http"
)

// @Summary		Создать нового пользователя
// @Description	Создаем нового пользователя с указанием имени, почты и пароля.
// @Description	Параметры передаются в теле запроса в качестве json-объекта.
// @Description	Если пользователя не удается создать, возвращаем ошибку.
// @Tags			Users
// @Accept			json
// @Produce		json
// @Param			input	body		dto.CreateUserRequest	true	"Login, email и Password"
// @Success		201		{object}	dto.CreateUserResponse	"Пользователь успешно создан"
// @Failure		400		{string}	string					"Invalid json payload или обязательные поля пустые"
// @Failure		500		{string}	string					"Cannot create user: <описание ошибки>"
// @Router			/user [post]
func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid json payload", http.StatusBadRequest)
		return
	}

	if req.Login == "" || req.Password == "" || req.Email == "" {
		http.Error(w, "Login, password, and email cannot be empty", http.StatusBadRequest)
		return
	}

	id, err := u.service.Create(context.Background(), req.Login, req.Email, req.Password)
	if err != nil {
		http.Error(w, "Cannot create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	u.logger.Info(context.Background(), "User has been created", logger.Field{"id", id})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	resp := dto.CreateUserResponse{
		Id:      id,
		Message: "User created",
	}
	_ = json.NewEncoder(w).Encode(resp)
}

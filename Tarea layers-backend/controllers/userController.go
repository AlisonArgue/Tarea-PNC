package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"layersapi/entities/dto"
	"layersapi/services"
	"net/http"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (u UserController) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	resData, err := u.userService.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	res, err := json.Marshal(resData)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (u UserController) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")

	updateData := dto.UpdateUser{
		Name:  name,
		Email: email,
	}

	err := u.userService.Update(id, updateData)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User actualizado :)"))
}

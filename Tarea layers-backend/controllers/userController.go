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
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User actualizado :)"))

}

func (u UserController) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := u.userService.Delete(id)
	if err != nil {
		if err.Error() == "user not found" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("User not found"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error"))
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Usuario borrado :)"))
}

func (u UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")

	newUser := dto.CreateUser{
		Name:  name,
		Email: email,
	}

	err := u.userService.Create(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

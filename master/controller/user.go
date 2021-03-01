package controller

import (
	"database/sql"
	"encoding/json"
	"mekar/master/config"
	"mekar/master/middleware"
	"mekar/master/model"
	"mekar/master/tools"
	"mekar/master/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserUseCase usecase.UserUseCase
}

func UserController(r *mux.Router, service usecase.UserUseCase) {
	UserHandler := UserHandler{service}
	User := r.PathPrefix("/user").Subrouter()
	User.HandleFunc("", UserHandler.CreateUser).Methods(http.MethodPost)
	User.HandleFunc("/{id}", UserHandler.UpdateUser).Methods(http.MethodPut)
	User.HandleFunc("", UserHandler.GetAllDataUser).Methods(http.MethodGet)
	User.HandleFunc("/{id}", UserHandler.UserById).Methods(http.MethodGet)
}

func AdminManagementController(r *mux.Router, service usecase.UserUseCase) {
	UserHandler := UserHandler{service}
	User := r.PathPrefix("/admin/user").Subrouter()
	isAuthOn := config.AuthSwitch()
	if isAuthOn {
		User.Use(middleware.TokenValidationMiddleware)
		detailAdminManagementController(User, UserHandler)
	} else {
		detailAdminManagementController(User, UserHandler)
	}
}
func detailAdminManagementController(AdminManagement *mux.Router, UserHandler UserHandler) {
	AdminManagement.HandleFunc("/{id}", UserHandler.UpdateUser).Methods(http.MethodPut)
	AdminManagement.HandleFunc("/{id}", UserHandler.DeleteUser).Methods(http.MethodDelete)
	AdminManagement.HandleFunc("", UserHandler.GetAllDataUser).Methods(http.MethodGet)
	AdminManagement.HandleFunc("/{id}", UserHandler.UserById).Methods(http.MethodGet)
}
func (s UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var User *model.UserModel
	var response tools.Response
	_ = json.NewDecoder(r.Body).Decode(&User)
	data, err := s.UserUseCase.CreateUser(User)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Message = "Failed"
		response.Data = nil
		byteData, err := json.Marshal(response)
		if err != nil {
			w.Write([]byte("Something Wrong on Marshalling Data"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(byteData)
	} else {
		response.Status = http.StatusCreated
		response.Message = "Success"
		response.Data = data
		byteData, err := json.Marshal(response)
		if err != nil {
			w.Write([]byte("Something Wrong on Marshalling Data"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(byteData)
	}
}
func (s UserHandler) GetAllDataUser(w http.ResponseWriter, r *http.Request) {
	var response tools.Response
	data, err := s.UserUseCase.GetAllDataUser()
	if err != nil {
		response.Status = http.StatusNotFound
		response.Message = "Not Found"
		response.Data = nil
		byteData, err := json.Marshal(response)
		if err != nil {
			w.Write([]byte("Something Wrong on Marshalling Data"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(byteData)
	} else {
		response.Status = http.StatusOK
		response.Message = "Success"
		response.Data = data
		byteData, err := json.Marshal(response)
		if err != nil {
			w.Write([]byte("Something Wrong on Marshalling Data"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(byteData)
	}

}
func (s UserHandler) UserById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idUser := param["id"]
	var response tools.Response
	data, err := s.UserUseCase.GetUserById(idUser)
	if err != nil {
		if err == sql.ErrNoRows {
			response.Status = http.StatusNotFound
			response.Message = "Not Found"
			response.Data = nil
			byteData, err := json.Marshal(response)
			if err != nil {
				w.Write([]byte("Something Wrong on Marshalling Data"))
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write(byteData)
		}
	} else {
		response.Status = http.StatusOK
		response.Message = "Success"
		response.Data = data
		byteData, err := json.Marshal(response)
		if err != nil {
			w.Write([]byte("Something Wrong on Marshalling Data"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(byteData)
	}
}

func (s UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	User := new(model.UserModel)
	param := mux.Vars(r)
	idUser := param["id"]
	var response tools.Response
	_ = json.NewDecoder(r.Body).Decode(&User)
	data, err := s.UserUseCase.UpdateUserById(idUser, User)
	if err != nil {
		if err == sql.ErrNoRows {
			response.Status = http.StatusNotFound
			response.Message = "Failed"
			response.Data = nil
			byteData, err := json.Marshal(response)
			if err != nil {
				w.Write([]byte("Something Wrong on Marshalling Data"))
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write(byteData)
		}
	} else {
		response.Status = http.StatusOK
		response.Message = "Success"
		response.Data = data
		byteData, err := json.Marshal(response)
		if err != nil {
			w.Write([]byte("Something Wrong on Marshalling Data"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(byteData)
	}
}
func (s UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idUser := param["id"]
	var response tools.Response
	id, err := s.UserUseCase.DeleteUserById(idUser)
	if err != nil {
		if err == sql.ErrNoRows {
			response.Status = http.StatusNotFound
			response.Message = "Failed"
			response.Data = nil
			byteData, err := json.Marshal(response)
			if err != nil {
				w.Write([]byte("Something Wrong on Marshalling Data"))
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write(byteData)
		}
	} else {
		response.Status = http.StatusOK
		response.Message = "User with ID " + id + " Success Deleted"
		response.Data = id
		byteData, err := json.Marshal(response)
		if err != nil {
			w.Write([]byte("Something Wrong on Marshalling Data"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(byteData)
	}
}

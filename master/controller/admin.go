package controller

import (
	"encoding/json"
	"fmt"
	"mekar/master/model"
	"mekar/master/tools"
	"mekar/master/tools/jwt"
	"mekar/master/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

type AdminHandler struct {
	AdminUseCase usecase.AdminUsecase
}

func AdminController(r *mux.Router, s usecase.AdminUsecase) {
	adminHandler := AdminHandler{AdminUseCase: s}
	admin := r.PathPrefix("/admin").Subrouter()
	admin.HandleFunc("/login", adminHandler.AdminLogin).Methods(http.MethodPost)
}
func (uh *AdminHandler) AdminLogin(w http.ResponseWriter, r *http.Request) {
	var data model.AdminModel
	err := json.NewDecoder(r.Body).Decode(&data)
	fmt.Print(err, &data)
	dataAdmin, isValid, _ := uh.AdminUseCase.AdminLogin(&data)
	w.Header().Set("Content-type", "application/json")
	if isValid {
		token, err := jwt.JwtEncoder(data.Username, "rahasiadong")
		if err != nil {
			http.Error(w, "Failed token generation", http.StatusUnauthorized)
		} else {
			var tools tools.Response
			tools.Status = http.StatusOK
			tools.Message = "Success"
			tools.Token = token
			tools.Data = dataAdmin
			byteData, err := json.Marshal(tools)
			if err != nil {
				w.Write([]byte("Something Wrong on Marshalling Data"))
			}
			w.WriteHeader(http.StatusOK)
			w.Write(byteData)
		}
	} else {
		var tools tools.Response
		tools.Status = http.StatusBadRequest
		tools.Message = "Login User Failed"
		tools.Data = nil
		byteData, err := json.Marshal(tools)
		if err != nil {
			w.Write([]byte("Something Wrong on Marshalling Data"))
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(byteData)
	}
}

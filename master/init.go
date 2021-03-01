package master

import (
	"database/sql"
	"mekar/master/controller"
	"mekar/master/middleware"
	"mekar/master/repositories"
	"mekar/master/usecase"

	"github.com/gorilla/mux"
)

func InitAll(R *mux.Router, DB *sql.DB, activityLog bool) {
	AdminRepo := repositories.InitAdminRepoImpl(DB)
	AdminUsecase := usecase.InitAdminAccUsecase(AdminRepo)
	controller.AdminController(R, AdminUsecase)

	UserRepo := repositories.InitUserRepoImpl(DB)
	UserUseCase := usecase.InitUserUseCase(UserRepo)
	controller.UserController(R, UserUseCase)
	controller.AdminManagementController(R, UserUseCase)

	if activityLog == true {
		R.Use(middleware.ActivityLogMiddleware)
	}
}

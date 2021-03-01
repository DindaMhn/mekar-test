package usecase

import "mekar/master/model"

type AdminUsecase interface {
	AdminLogin(*model.AdminModel) (*model.AdminModel, bool, error)
}

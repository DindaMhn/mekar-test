package usecase

import (
	"mekar/master/model"
	"mekar/master/repositories"
)

type AdminAccUsecaseImpl struct {
	adminRepo repositories.AdminAccount
}

func InitAdminAccUsecase(adminRepo repositories.AdminAccount) AdminUsecase {
	return &AdminAccUsecaseImpl{adminRepo: adminRepo}
}

func (uc *AdminAccUsecaseImpl) AdminLogin(admin *model.AdminModel) (*model.AdminModel, bool, error) {

	data, isValid, err := uc.adminRepo.AdminLogin(admin)

	if err != nil {
		return nil, false, err
	}

	return data, isValid, nil
}

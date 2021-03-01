package usecase

import "mekar/master/model"

type UserUseCase interface {
	CreateUser(DataUser *model.UserModel) (*model.UserModel, error)
	UpdateUserById(id string, DataUser *model.UserModel) (*model.UserModel, error)
	GetUserById(id string) (*model.UserModel, error)
	GetAllDataUser() ([]*model.UserModel, error)
	DeleteUserById(id string) (string, error)
}

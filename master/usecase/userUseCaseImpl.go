package usecase

import (
	"fmt"
	"mekar/master/model"
	"mekar/master/repositories"
)

type UserUsecaseImpl struct {
	UserRepo repositories.UserRepo
}

func (s UserUsecaseImpl) CreateUser(DataUser *model.UserModel) (*model.UserModel, error) {
	data, err := s.UserRepo.CreateUser(DataUser)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (s UserUsecaseImpl) UpdateUserById(id string, DataUser *model.UserModel) (*model.UserModel, error) {

	data, err := s.UserRepo.UpdateUserById(id, DataUser)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (au UserUsecaseImpl) GetAllDataUser() ([]*model.UserModel, error) {
	data, err := au.UserRepo.GetAllDataUser()
	if err != nil {
		return nil, err
	}
	fmt.Println("User", data)
	return data, nil
}
func (s UserUsecaseImpl) GetUserById(id string) (*model.UserModel, error) {
	data, err := s.UserRepo.GetUserById(id)

	return data, err
}
func (s UserUsecaseImpl) DeleteUserById(id string) (string, error) {

	id, err := s.UserRepo.DeleteUserById(id)
	if err != nil {
		return "", err
	}
	return id, nil
}
func InitUserUseCase(UserRepo repositories.UserRepo) UserUseCase {
	return &UserUsecaseImpl{UserRepo}
}

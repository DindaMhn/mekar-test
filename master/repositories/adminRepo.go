package repositories

import "mekar/master/model"

type AdminAccount interface {
	AdminLogin(*model.AdminModel) (*model.AdminModel, bool, error)
}

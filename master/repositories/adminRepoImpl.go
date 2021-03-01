package repositories

import (
	"database/sql"
	"mekar/master/model"
	"mekar/master/tools"
	"mekar/master/tools/pwd"
)

type AdminRepoImpl struct {
	db *sql.DB
}

func InitAdminRepoImpl(db *sql.DB) AdminAccount {
	return &AdminRepoImpl{db: db}
}
func (ur *AdminRepoImpl) AdminLogin(admin *model.AdminModel) (*model.AdminModel, bool, error) {
	row := ur.db.QueryRow(tools.SELECT_ADMIN_LOGIN, admin.Username)
	var adminAcc = model.AdminModel{}
	err := row.Scan(&adminAcc.Username, &adminAcc.Password)

	if err != nil {
		return nil, false, err
	}
	isPwdValid := pwd.CheckPasswordHash(admin.Password, adminAcc.Password)
	if (admin.Username == adminAcc.Username) && isPwdValid {
		admin.Password = adminAcc.Password
		return admin, true, nil
	} else {
		return nil, false, err
	}
}

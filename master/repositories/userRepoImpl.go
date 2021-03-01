package repositories

import (
	"database/sql"
	"log"
	"mekar/master/model"
	"mekar/master/tools"

	"github.com/google/uuid"
)

type UserRepoImpl struct {
	db *sql.DB
}

func (ar UserRepoImpl) CreateUser(DataUser *model.UserModel) (*model.UserModel, error) {
	tr, err := ar.db.Begin()

	_, _ = ar.db.Exec("SET FOREIGN_KEY_CHECKS=0;")
	DataUser.IdUser = uuid.New().String()
	query := tools.CREATE_USER
	row, err := ar.db.Query(query, &DataUser.IdUser, &DataUser.NamaUser,
		&DataUser.TanggalLahir, &DataUser.NoKTP, &DataUser.Pekerjaan, &DataUser.Pendidikan,
		&DataUser.Status)
	if err != nil {
		tr.Rollback()
		log.Fatal(err)
	} else {
		tr.Commit()
	}
	row.Close()
	_, _ = ar.db.Exec("SET FOREIGN_KEY_CHECKS=1;")
	return DataUser, nil
}
func (ar UserRepoImpl) GetAllDataUser() ([]*model.UserModel, error) {

	User := []*model.UserModel{}
	query := tools.SELECT_ALL_USER
	data, err := ar.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer data.Close()
	for data.Next() {
		DataUser := model.UserModel{}
		var err = data.Scan(&DataUser.IdUser, &DataUser.NamaUser,
			&DataUser.TanggalLahir, &DataUser.NoKTP, &DataUser.Pekerjaan, &DataUser.Pendidikan,
			&DataUser.Status)
		if err != nil {
			return nil, err
		}

		User = append(User, &DataUser)
	}

	return User, nil
}
func (ar UserRepoImpl) GetUserById(id string) (*model.UserModel, error) {
	DataUser := new(model.UserModel)
	query := tools.SELECT_USER_BY_ID
	err := ar.db.QueryRow(query, id).Scan(&DataUser.IdUser, &DataUser.NamaUser,
		&DataUser.TanggalLahir, &DataUser.NoKTP, &DataUser.Pekerjaan, &DataUser.Pendidikan,
		&DataUser.Status)

	return DataUser, err
}

func (ar UserRepoImpl) UpdateUserById(id string, DataUser *model.UserModel) (*model.UserModel, error) {
	tr, err := ar.db.Begin()
	_, _ = ar.db.Exec("SET FOREIGN_KEY_CHECKS=0;")
	_, err = ar.GetUserById(id)
	if err != nil {
		tr.Rollback()
		return nil, err
	}
	query := tools.UPDATE_DATA_USER
	row, err := ar.db.Query(query, &DataUser.NamaUser,
		&DataUser.TanggalLahir, &DataUser.NoKTP, &DataUser.Pekerjaan, &DataUser.Pendidikan, id)
	if err != nil {
		tr.Rollback()
		return nil, err
	} else {
		tr.Commit()
	}

	row.Close()
	_, _ = ar.db.Exec("SET FOREIGN_KEY_CHECKS=1;")
	DataUser.IdUser = id
	return DataUser, err
}
func (ar UserRepoImpl) DeleteUserById(id string) (string, error) {
	tr, err := ar.db.Begin()
	_, _ = ar.db.Exec("SET FOREIGN_KEY_CHECKS=0;")
	_, err = ar.GetUserById(id)
	if err != nil {
		tr.Rollback()
		return "", err
	}
	query := tools.DELETE_USER
	row, err := ar.db.Query(query, id)
	if err != nil {
		tr.Rollback()
		return "", err
	} else {
		tr.Commit()
	}

	row.Close()
	_, _ = ar.db.Exec("SET FOREIGN_KEY_CHECKS=1;")
	return id, err
}
func InitUserRepoImpl(db *sql.DB) UserRepo {
	return &UserRepoImpl{db}
}

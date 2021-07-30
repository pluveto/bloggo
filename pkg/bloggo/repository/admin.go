package repository

import (
	"errors"

	"github.com/pluveto/bloggo/pkg/bloggo/model"
	"gorm.io/gorm"
)

func (repo *Repository) AdminCreate(admin *model.Admin) (err error) {
	err = repo.db.Create(admin).Error
	if err != nil && errors.Is(err, gorm.ErrInvalidValue) {
		return
	}
	return
}

func (repo *Repository) AdminGet(id uint64) (model *model.Admin, err error) {
	return repo.adminGetByUniqueKeyValue("id", id)
}

func (repo *Repository) AdminGetByUsername(username string) (model *model.Admin, err error) {
	return repo.adminGetByUniqueKeyValue("username", username)
}

func (repo *Repository) AdminGetByEmail(email string) (model *model.Admin, err error) {
	return repo.adminGetByUniqueKeyValue("email", email)
}

func (repo *Repository) adminGetByUniqueKeyValue(key string, value interface{}) (ret *model.Admin, err error) {
	var m = model.Admin{}
	err = repo.db.Where(key+" = ?", value).First(&m).Error
	err = checkDbError(err)
	if err != nil {
		return nil, err
	}
	return &m, err
}

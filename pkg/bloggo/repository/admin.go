package repository

import (
	"errors"

	"github.com/pluveto/bloggo/pkg/bloggo/model"
	"gorm.io/gorm"
)

func (repo *Repository) AdminCreate(admin *model.Admin) (err error) {
	err = repo.db.Create(admin).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	return
}

func (repo *Repository) GetAdmin(id uint64) (model model.Admin, err error) {
	err = repo.db.Where("id = ?", id).First(&model).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	return
}

package repository

import (
	"errors"

	"github.com/pluveto/bloggo/pkg/bloggo/model"
	"gorm.io/gorm"
)

func (repo *Repository) PostGetByID(id int64) (model *model.Post, err error) {
	return repo.PostGetByUniqueKeyValue("id", id)
}
func (repo *Repository) PostGetBySlug(slug string) (model *model.Post, err error) {
	return repo.PostGetByUniqueKeyValue("slug", slug)
}

func (repo *Repository) PostUpdate(model *model.Post) (err error) {
	return repo.db.Save(model).Error
}

func (repo *Repository) PostGetByUniqueKeyValue(key string, value interface{}) (ret *model.Post, err error) {
	var m = model.Post{}
	err = repo.db.Where(key+" = ?", value).First(&m).Error
	err = checkDbError(err)
	if err != nil {
		return nil, err
	}
	return &m, err
}

func (repo *Repository) PostCreate(post *model.Post) (err error) {
	err = repo.db.Create(post).Error
	if err != nil && errors.Is(err, gorm.ErrInvalidValue) {
		return
	}
	return
}

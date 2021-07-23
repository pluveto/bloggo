package repository

import (
	"errors"

	"github.com/pluveto/bloggo/pkg/bloggo/model"
	"gorm.io/gorm"
)

func (repo *Repository) GetSetting(key string) (value string, err error) {
	var setting model.Setting
	err = repo.db.Where("key = ?", key).First(&setting).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	value = setting.Value
	return
}

func (repo *Repository) SetSetting(key string, value string) (err error) {
	err = repo.db.Model(&model.Setting{}).Where("key = ?", key).Update("value", value).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	return
}

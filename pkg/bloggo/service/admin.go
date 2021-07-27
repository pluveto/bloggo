package service

import (
	"github.com/pluveto/bloggo/pkg/bloggo/model"
	"github.com/pluveto/bloggo/pkg/util"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string, salt string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password+salt), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string, salt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+salt))
	return err == nil
}

// AdminCreate 创建一个后台用户
func (s *Service) AdminCreate(username string, password string) (ret *model.Admin, err error) {
	var salt = util.RandString(8)
	hashedPassword, err := HashPassword(password, salt)
	if err != nil {
		return
	}
	var model = model.Admin{
		Username: username,
		Password: hashedPassword,
	}
	err = s.repo.AdminCreate(&model)
	if err != nil {
		ret = &model
	}
	return
}

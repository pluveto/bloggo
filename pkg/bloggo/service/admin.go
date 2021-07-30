package service

import (
	"github.com/pluveto/bloggo/pkg/bloggo/model"
	"github.com/pluveto/bloggo/pkg/errcode"
	"github.com/pluveto/bloggo/pkg/util"
	"golang.org/x/crypto/bcrypt"
)

// ! 注意，bcrypt 同样的输入两次调用此函数输出是不同的，不要用它来检验密码
func HashPassword(password string, salt string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password+salt), 14)
	return string(bytes), err
}

func IsPasswordValid(hashedPassword string, password string, salt string) bool {
	return nil == bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+salt))
}

func CheckPasswordHash(password, hash string, salt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+salt))
	return err == nil
}

// AdminCreateByEmail 创建一个后台用户
func (s *Service) AdminCreateByEmail(email string, password string) (ret *model.Admin, err error) {
	var salt = util.RandString(8)
	hashedPassword, err := HashPassword(password, salt)
	if err != nil {
		return
	}
	var model = &model.Admin{
		Email:    email,
		Password: hashedPassword,
	}
	err = s.repo.AdminCreate(model)
	if err != nil {
		return
	}
	ret = model
	return
}

func (s *Service) AdminGetByEmail(email string) (ret *model.Admin, err error) {
	ret, err = s.repo.AdminGetByEmail(email)
	return
}

func (s *Service) AdminGetByUsername(username string) (ret *model.Admin, err error) {
	ret, err = s.repo.AdminGetByUsername(username)
	return
}

func (s *Service) AdminAuthByEmail(email string, password string) (ret *model.Admin, err error) {
	ret, err = s.repo.AdminGetByEmail(email)
	if err != nil {
		return
	}
	if IsPasswordValid(ret.Password, password, ret.Salt) {
		err = errcode.ApiError(errcode.ErrWrongPassword)
		ret = nil
		return
	}
	return
}

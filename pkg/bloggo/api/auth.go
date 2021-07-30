package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pluveto/bloggo/pkg/errcode"
)

type AuthLoginReq struct {
	Email    string `form:"email"    json:"email"    binding:"email,required,min=6,max=255"`
	Password string `form:"password" json:"password" binding:"printascii,required"`
	Captcha  string `form:"captcha"  json:"captcha"  binding:"len=6"`
}

type AuthLoginRet struct {
	Token string `json:"token"`
}

func (api *API) AuthLogin(c *gin.Context) {
	// form check
	var (
		req AuthLoginReq
		err = c.ShouldBindJSON(&req)
	)
	if err != nil {
		api.retBadParam(c, err)
		return
	}
	// TODO: 对多次登录失败的进行验证码检查
	admin, err := api.Service.AdminAuthByEmail(req.Email, req.Password)
	if err != nil {
		api.RetJSON(c, nil, err)
		return
	}
	var claims = api.Service.TokenCreateClaims(admin.ID)
	tokenStr, err := api.Service.TokenCreate(*claims)
	if err != nil {
		api.RetJSON(c, nil, err)
		return
	}
	api.RetJSON(c, tokenStr, err)
}

func (api *API) retBadParam(c *gin.Context, err error) {
	api.RetJSON(c, nil, errcode.ApiError(errcode.ErrBadParam))
	var errs, ok = err.(validator.ValidationErrors)
	if ok {
		fmt.Printf("first param err: %v\n", errs[0])
	}
}

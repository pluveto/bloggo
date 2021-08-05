package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pluveto/bloggo/pkg/bloggo/model"
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

type AuthGetUserInfoRet struct {
	Username    string `json:"username"`
	ScreenName  string `json:"screenName"`
	ID          int64  `json:"id"`
	Email       string `json:"email"`
	Description string `json:"description"`
	AvatarUrl   string `json:"avatarUrl"`
}

func (api *API) AuthGetUserInfo(c *gin.Context) {
	tmpid, _ := c.Get("uid")
	var uid = tmpid.(int64)
	admin, _ := api.Service.AdminGet(uid)
	api.RetJSON(c, &AuthGetUserInfoRet{
		Username:    admin.Username,
		ID:          admin.ID,
		ScreenName:  admin.ScreenName,
		Email:       admin.Email,
		Description: admin.Description,
		AvatarUrl:   admin.AvatarUrl,
	}, nil)
}

// todo: 参数校验
type AuthUpdateUserInfoReq struct {
	Username    string `json:"username"`
	ScreenName  string `json:"screenName"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Description string `json:"description"`
	AvatarUrl   string `json:"avatarUrl"`
}

func (api *API) AuthUpdateUserInfo(c *gin.Context) {
	var (
		req AuthUpdateUserInfoReq
		err error = c.ShouldBindJSON(&req)
	)

	if err != nil {
		api.retBadParam(c, err)
		return
	}
	tmpid, _ := c.Get("uid")
	var uid = tmpid.(int64)
	// todo: 冲突检查
	var m = &model.Admin{
		ID: uid,
		Username:    req.Username,
		ScreenName:  req.ScreenName,
		Password:    req.Password,
		Email:       req.Email,
		Description: req.Description,
		AvatarUrl:   req.AvatarUrl,
	}
	err = api.Service.AdminUpdate(m)
	api.RetJSON(c, nil, err)
}

func (api *API) retBadParam(c *gin.Context, err error) {
	var errs, ok = err.(validator.ValidationErrors)
	if ok {
		fmt.Printf("first param err: %v\n", errs[0])
	} else {
		// ! 可能是 JSON 无效引起的
		fmt.Printf("first param err 2: %v\n", errs.Error())
	}
	api.RetJSON(c, nil, errcode.ApiError(errcode.ErrBadParam))

}

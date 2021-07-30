package api

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pluveto/bloggo/pkg/bloggo"
	"github.com/pluveto/bloggo/pkg/errcode"
)

type AdminCreateReq struct {
	Email    string `form:"email"    json:"email"    binding:"email,required,min=6,max=255"`
	Password string `form:"password" json:"password" binding:"printascii,required,min=8,max=255"`
}

type AdminCreateRet struct {
	ID string `json:"id"`
}

func (api *API) AdminCreate(c *gin.Context) {
	// ! Debug only
	// TODO: 放到中间件进行判断
	if gin.Mode() != gin.DebugMode {
		return
	}
	var (
		req AdminCreateReq
		err error = c.ShouldBindJSON(&req)
	)

	if err != nil {
		api.retBadParam(c, err)
		return
	}
	// 存在性检查，如果已经存在会返回 err，并为 existingAdmin 赋值
	existingAdmin, _ := api.Service.AdminGetByEmail(req.Email)
	if existingAdmin != nil {
		api.RetJSON(c, nil, errcode.ApiError(bloggo.ErrConflictEmail))
		return
	}
	newAdmin, err := api.Service.AdminCreateByEmail(req.Email, req.Password)
	if err != nil {
		api.RetJSON(c, nil, err)
		return
	}
	var retModel = AdminCreateRet{ID: strconv.FormatInt(int64(newAdmin.ID), 10)}
	log.Printf("%v", retModel)
	api.RetJSON(c, retModel, err)
}

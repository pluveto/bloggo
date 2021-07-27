/**
 * api package: 供 Http 层的所有 api 请求的入口
 * 使用方法：
 * 	使用 New 创建实例
 *  将实例上的方法绑定到路由
 */
package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pluveto/bloggo/pkg/bloggo"
	"github.com/pluveto/bloggo/pkg/bloggo/model"
	"github.com/pluveto/bloggo/pkg/bloggo/service"
	"github.com/pluveto/bloggo/pkg/errcode"
)

type API struct {
	service *service.Service
}

func New(c *bloggo.Config) *API {
	var api = &API{}
	api.service = service.New(c)
	return api
}

// 统一返回接口
// ! 错误请用 errcode 包生成，未定义的请在 New 函数中动态添加，不要四处散落自定义的错误
func (api *API) RetJSON(c *gin.Context, data interface{}, err error) {
	var apiErr = errcode.Cause(err)
	c.JSON(apiErr.ToHttpStatusCode(), &gin.H{
		"code":    apiErr.Code(),
		"message": apiErr.Message(),
		"data":    data,
	})
}

/**
 * 采用 RPC 命名
 * 命名约定见：https://google-cloud.gitbook.io/api-design-guide/naming_conventions
 * 使用模块名作为函数前缀，如 Admin 的 CreateAdmin，应该命名为 AdminCreate
 *
 * 参数校验的方法：https://segmentfault.com/a/1190000022541905
 */

func (api *API) GetPing(c *gin.Context) {
	api.RetJSON(c, gin.H{
		"message": "pong",
	}, nil)
}

func (api *API) GetHomeView(c *gin.Context) {
	var siteName, _ = api.service.GetSettingOrDefault("siteName", "未命名博客")
	var siteDescription, _ = api.service.GetSettingOrDefault("siteDescription", "")
	c.HTML(http.StatusOK, "home.html", gin.H{
		"siteName":        siteName,
		"siteDescription": siteDescription,
	})
}

func (api *API) AdminCreate(c *gin.Context) {
	// ! Debug only
	if gin.Mode() != gin.DebugMode {
		return
	}
	var err error
	var req model.AdminLoginReq
	err = c.Bind(&req)
	if err != nil {
		fmt.Printf("What the fuck? %v", err.Error())
		return
	}
	var m *model.Admin
	m, err = api.service.AdminCreate(req.Email, req.Password)
	api.RetJSON(c, &model.AdminCreateRet{ID: strconv.FormatInt(int64(m.ID), 10)}, err)
}

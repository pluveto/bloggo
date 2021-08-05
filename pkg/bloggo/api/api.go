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

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pluveto/bloggo/pkg/bloggo"
	"github.com/pluveto/bloggo/pkg/bloggo/service"
	"github.com/pluveto/bloggo/pkg/errcode"
	globalModel "github.com/pluveto/bloggo/pkg/model"
)

type API struct {
	Service *service.Service
	valid   *validator.Validate
}

func New(c *bloggo.Config) *API {
	var api = &API{}
	api.Service = service.New(c)
	api.valid = validator.New()

	return api
}

/**
 * 统一返回接口
 * ! 错误请用 errcode 包生成，未定义的请在 New 函数中动态添加，不要四处散落自定义的错误
 * ! 不要向这里传递指针，否则引发 runtime error: invalid memory address or nil pointer dereference
 */
func (api *API) RetJSON(c *gin.Context, data interface{}, err error) {
	var apiErr = errcode.Cause(err)
	var resp = &globalModel.RetWrap{
		Code:    apiErr.Code(),
		Message: apiErr.Message(),
		Data:    data,
	}
	c.JSON(apiErr.ToHttpStatusCode(), resp)
	fmt.Printf("response: %v\n", resp)
}

/**
 * 错误返回接口。已编码的错误尽量用 RetJSON 返回。
 * ! 此接口仅供测试使用，因为这里返回的错误没有编码到统一错误码中
 */
func (api *API) RetError(c *gin.Context, http_code int, message string) {
	if gin.Mode() != gin.DebugMode {
		panic("RetError method is only available in debug mode")
	}
	var resp = &gin.H{
		"code":    http_code * 1000,
		"message": message,
		"data":    nil,
	}
	c.JSON(http_code, resp)
	fmt.Printf("response: %v\n", resp)
}

/**
 * 采用 RPC 命名
 * 命名约定见：https://google-cloud.gitbook.io/api-design-guide/naming_conventions
 * List Get Create Update Rename Delete
 * 使用模块名作为函数前缀，如 Admin 的 CreateAdmin，应该命名为 AdminCreate
 *
 * 参数校验的方法：https://segmentfault.com/a/1190000022541905
 */

type PingReq struct{}

func (api *API) GetPing(c *gin.Context) {
	api.RetJSON(c, gin.H{
		"message": "pong",
	}, nil)
}

func (api *API) GetHomeView(c *gin.Context) {
	var siteName, _ = api.Service.GetSettingOrDefault("siteName", "未命名博客")
	var siteDescription, _ = api.Service.GetSettingOrDefault("siteDescription", "")
	c.HTML(http.StatusOK, "home.html", gin.H{
		"siteName":        siteName,
		"siteDescription": siteDescription,
	})
}

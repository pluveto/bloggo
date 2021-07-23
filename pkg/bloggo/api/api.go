package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pluveto/bloggo/pkg/bloggo"
	"github.com/pluveto/bloggo/pkg/bloggo/service"
)

type API struct {
	service *service.Service
}

func New(c *bloggo.Config) *API {
	var api = &API{}
	api.service = service.New(c)
	return api
}

func (api *API) GetPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (api *API) GetHomeView(c *gin.Context) {
	var siteName = api.service.GetSetting("siteName")
	c.HTML(http.StatusOK, "home.html", gin.H{
		"siteName":        siteName,
		"siteDescription": "Written in Golang",
	})
}

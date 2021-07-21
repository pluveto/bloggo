package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func GetHomeView(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"siteName":        "Cool Blog",
		"siteDescription": "Written in Golang",
	})
}

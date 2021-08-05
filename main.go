package main

import (
	"html/template"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pluveto/bloggo/pkg/bloggo/api"
	"github.com/pluveto/bloggo/pkg/bloggo/middleware"
)

func main() {
	// TODO: Config
	var api = api.New(nil)

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.SetFuncMap(template.FuncMap{
		"formatAsDate": func(tsec int64) string {
			tm := time.Unix(tsec, 0)
			return tm.Format("2006-01-02 03:04:05 PM")
		},
	})
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"PUT", "PATCH", "OPTIONS"},
		AllowHeaders: []string{"Origin", "content-type", "content-length", "authorization"},
		// ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	r.Use(func(c *gin.Context) {
		/* 		var buf bytes.Buffer
		   		tee := io.TeeReader(c.Request.Body, &buf)
		   		body, _ := ioutil.ReadAll(tee)
		   		c.Request.Body = ioutil.NopCloser(&buf)
		   		fmt.Printf("%v\n", c.Request.Header)
		   		fmt.Printf("%v\n", string(body)) */
		c.Next()
	})
	// var jwt = middleware.JWTAuth(api.Service)
	// r.Use(middleware.RequestValidator())
	r.GET("/", api.GetHomeView)
	r.GET("/ping", api.GetPing)
	r.POST("/admin/create", api.AdminCreate)
	r.POST("/auth/login", api.AuthLogin)
	r.POST("/auth/getUserInfo", middleware.JWTAuth(api.Service), api.AuthGetUserInfo)
	r.POST("/auth/updateUserInfo", middleware.JWTAuth(api.Service), api.AuthUpdateUserInfo)
	
	r.Run()
}

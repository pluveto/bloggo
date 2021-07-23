package main

import (
	"html/template"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pluveto/bloggo/pkg/bloggo/api"
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

	r.GET("/", api.GetHomeView)
	r.GET("/ping", api.GetPing)
	r.Run()
}

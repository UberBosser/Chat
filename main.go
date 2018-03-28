package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/gzip"
)

func indexHandler(c *gin.Context) {
	c.HTML(200, "index.tmpl", nil)
}

func noHandler(c *gin.Context) {
	c.HTML(404, "404.tmpl", nil)
}

func main() {
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.LoadHTMLGlob("./templates/*/*.tmpl")

	router.Static("/static", "./static")
	router.GET("/", indexHandler)

	router.NoRoute(noHandler)

	router.Run()
}

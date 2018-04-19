package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	godb "gopkg.in/gorethink/gorethink.v4"
	"log"
)

func indexHandler(c *gin.Context) {
	c.HTML(200, "index.tmpl", nil)
}

func noHandler(c *gin.Context) {
	c.HTML(404, "404.tmpl", nil)
}

func main() {
	session, err := godb.Connect(godb.ConnectOpts{
		Address:  "localhost:28015",
		Database: "chat",
	})
	if err != nil {
		log.Println(err.Error())
		return
	}
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.LoadHTMLGlob("./templates/*/*.tmpl")

	router.Static("/static", "./static")
	router.GET("/", indexHandler)

	socketRouter := newSocketRouter(session)
	router.GET("/socket", socketRouter.websocketHandler)
	socketRouter.Handle("send message", sendMessage)

	router.NoRoute(noHandler)

	router.Run()
}

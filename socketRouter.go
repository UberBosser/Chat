package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	godb "gopkg.in/gorethink/gorethink.v4"
	"log"
	"net/http"
)

type Handler func(*Client, interface{})

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Router struct {
	rules   map[string]Handler
	session *godb.Session
}

func newSocketRouter(session *godb.Session) *Router {
	return &Router{
		rules:   make(map[string]Handler),
		session: session,
	}
}

func (r *Router) Handle(code string, handler Handler) {
	r.rules[code] = handler
}

func (r *Router) FindHandler(code string) (Handler, bool) {
	handler, found := r.rules[code]
	return handler, found
}

func (r *Router) websocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	client := newClient(conn, r.FindHandler, r.session)
	go client.Read()
	go client.Write()
	go client.Subscribe()
}

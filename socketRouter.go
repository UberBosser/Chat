package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
	chans   []chan Message
	message chan Message
	close   chan chan Message
}

func newSocketRouter() *Router {
	return &Router{
		rules:   make(map[string]Handler),
		message: make(chan Message),
		close:   make(chan chan Message),
	}
}

func (r *Router) Handle(code string, handler Handler) {
	r.rules[code] = handler
}

func (r *Router) FindHandler(code string) (Handler, bool) {
	handler, found := r.rules[code]
	return handler, found
}

func (r *Router) SendAll() {
	for {
		select {
		case m := <-r.message:
			for _, client := range r.chans {
				client <- m
			}
		case c := <-r.close:
			for i, channel := range r.chans {
				if channel == c {
					r.chans = append(r.chans[:i], r.chans[i+1:]...)
				}
			}
		}
	}
}

func (r *Router) WebsocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	client := newClient(conn, r.FindHandler, r.message, r.close)
	go client.Read()
	go client.Write()
	r.chans = append(r.chans, client.message)
}

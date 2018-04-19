package main

import (
	"github.com/gorilla/websocket"
	godb "gopkg.in/gorethink/gorethink.v4"
)

type FindHandler func(string) (Handler, bool)

type Message struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
}

type Client struct {
	conn        *websocket.Conn
	message     chan Message
	session     *godb.Session
	stop        chan bool
	findHandler FindHandler
}

func (c *Client) Read() {
	var message Message
	for {
		if err := c.conn.ReadJSON(&message); err != nil {
			break
		}
		if handler, found := c.findHandler(message.Code); found {
			handler(c, message.Data)
		}
	}
	c.stop <- true
	c.conn.Close()
}

func (c *Client) Write() {
	for m := range c.message {
		if err := c.conn.WriteJSON(m); err != nil {
			break
		}
	}
	c.stop <- true
	c.conn.Close()
}

func newClient(conn *websocket.Conn, findHandler FindHandler, session *godb.Session) *Client {
	return &Client{
		conn:        conn,
		message:     make(chan Message),
		session:     session,
		stop:        make(chan bool),
		findHandler: findHandler,
	}
}

package main

import (
	"github.com/gorilla/websocket"
)

type FindHandler func(string) (Handler, bool)

type Message struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
}

type Client struct {
	conn        *websocket.Conn
	message     chan Message
	messageAll  chan Message
	stop        chan chan Message
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
	c.stop <- c.message
	c.conn.Close()
}

func (c *Client) Write() {
	for m := range c.message {
		if err := c.conn.WriteJSON(m); err != nil {
			break
		}
	}
	c.stop <- c.message
	c.conn.Close()
}

func newClient(conn *websocket.Conn, findHandler FindHandler, messageAll chan Message, stop chan chan Message) *Client {
	return &Client{
		conn:        conn,
		message:     make(chan Message),
		messageAll:  messageAll,
		stop:        stop,
		findHandler: findHandler,
	}
}

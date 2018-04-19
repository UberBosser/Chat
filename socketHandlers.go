package main

import (
	"github.com/mitchellh/mapstructure"
	godb "gopkg.in/gorethink/gorethink.v4"
	"log"
)

type ChatMessage struct {
	Id      string `gorethink:"id,omitempty"`
	Name    string `json:"name" gorethink:"name"`
	Message string `json:"message" gorethink:"message"`
}

func sendMessage(c *Client, data interface{}) {
	var message ChatMessage
	mapstructure.Decode(data, &message)
	err := godb.Table("messages").Insert(message).Exec(c.session)
	if err != nil {
		c.message <- Message{Code: "error", Data: err.Error()}
	}
}

func (c *Client) Subscribe() {
	result := make(chan godb.ChangeResponse)
	cursor, err := godb.Table("messages").Changes().Run(c.session)
	if err != nil {
		log.Println(err.Error())
		c.message <- Message{
			Code: "error",
			Data: err.Error(),
		}
		return
	}
	go func() {
		var change godb.ChangeResponse
		for cursor.Next(&change) {
			result <- change
		}
	}()
	for {
		select {
		case change := <-result:
			c.message <- Message{Code: "receive message", Data: change.NewValue}
		case <-c.stop:
			cursor.Close()
			return
		}
	}
}

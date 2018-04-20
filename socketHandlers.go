package main

type ChatMessage struct {
	Id      string `gorethink:"id,omitempty"`
	Name    string `json:"name" gorethink:"name"`
	Message string `json:"message" gorethink:"message"`
}

func sendMessage(c *Client, data interface{}) {
	c.messageAll <- Message{Code: "receive message", Data: data}
}

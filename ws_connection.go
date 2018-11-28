package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type connection struct {
	ws *websocket.Conn

	send chan interface{}
}

var (
	upgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func (c *connection) reader() {
	for {
		_, _, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
	}

	c.ws.Close()
}

func (c *connection) writer() {
	for change := range c.send {
		err := c.ws.WriteJSON(change)
		if err != nil {
			break
		}
	}

	c.ws.Close()
}

func wsHandler(o observer) http.HandlerFunc {
	log.Println("stating websocketserver...")

	return func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		c := &connection{send: make(chan interface{}, 256), ws: ws}
		o.register <- c
		defer func() { o.unregister <- c }()
		go c.writer()
		c.reader()
	}
}

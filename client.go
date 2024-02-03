package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	c *Chatroom

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

func Read(c *Client) {
	defer func() {
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			return
		}
		c.c.broadcast <- message
	}
}

func Write(c *Client) {
	defer func() {
		c.conn.Close()
		c.c.unregister <- c
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				return
			}

			err := c.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				return
			}
		}
	}
}

func serveWs(c *Chatroom, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{c: c, conn: conn, send: make(chan []byte, 256)}
	client.c.register <- client
	go Write(client)
	Read(client)
}

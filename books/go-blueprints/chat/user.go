package main

import (
	"github.com/gorilla/websocket"
)

const (
	messageBufferSize = 256
)

type User struct {
	socket *websocket.Conn
	send   chan []byte
	room   *Room
}

func (u *User) read() {
	defer u.socket.Close()
	for {
		_, msg, err := u.socket.ReadMessage()
		if err != nil {
			return
		}
		u.room.forward <- msg
	}
}

func (u *User) write() {
	defer u.socket.Close()
	for msg := range u.send {
		err := u.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}

func newUser(socket *websocket.Conn, r *Room) *User {
	return &User{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
}

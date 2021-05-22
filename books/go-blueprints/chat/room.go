package main

const (
	socketBufferSize = 1024
)

type Room struct {
	forward chan []byte
	join    chan *User
	leave   chan *User
	users   map[*User]bool
}

// Join, leave or forward message to room
func (r *Room) run() {
	for {
		select {
		case user := <-r.join:
			r.users[user] = true
		case user := <-r.leave:
			delete(r.users, user)
			close(user.send)
		case msg := <-r.forward:
			for user := range r.users {
				user.send <- msg
			}
		}
	}
}

func newRoom() *Room {
	return &Room{
		forward: make(chan []byte),
		join:    make(chan *User),
		leave:   make(chan *User),
		users:   make(map[*User]bool),
	}
}

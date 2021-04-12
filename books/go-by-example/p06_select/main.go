package p06_select

import (
	"fmt"
	"time"
	"math/rand"
)

func Run() {
	rand.Seed(time.Now().UnixNano())
	// Select lets you wait on multiple channel operations

	chimeMessages := make(chan string)
	slackMessages := make(chan string)

	go func() {
		sleepRandom("chime")
		chimeMessages <- "Hi on chime"
	}()

	go func() {
		sleepRandom("slack")
		slackMessages <- "Hi on slack"
	}()

	for i := 0; i < 2; i++ { // needs to recieve to messages
		select { // recieves from any of the channels
		case msg1 := <-chimeMessages:
			fmt.Println("recieved", msg1)
		case msg2 := <-slackMessages:
			fmt.Println("recieved", msg2)
		}
	}
}

func sleepRandom(id string) {
	r := rand.Intn(500)+100
	fmt.Println("Sleeping", id, "for", r)
	time.Sleep(time.Duration(r)*time.Millisecond)
}
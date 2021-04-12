package p05_channels

import (
	"time"
	"fmt"
)

func Run() {
	fmt.Println(" ---- Channels ---- ")

	// will only send more messages after first one is consumed
	fmt.Println(" -- blocking channels -- ")
	sendMessages(make(chan string))

	fmt.Println(" -- buffered channels -- ")
	sendMessages(make(chan string, 4))

	fmt.Println(" -- Single Worker -- ")
	singleWorkerCoordination()
}

func sendMessages(messageChan chan string) {
	go func() {
		messages := []string{"Hello", "There", "This", "is", "go", "lang", "for", "you"}
		for _, message := range messages {
			messageChan <- message
			fmt.Println("Sent", message)
		}
	}()

	message := <- messageChan
	fmt.Println("Recieved", message)
	time.Sleep(100*time.Millisecond)

	message = <- messageChan
	fmt.Println("Recieved", message)
	time.Sleep(100*time.Millisecond)
}

func singleWorkerCoordination() {

	worker := func(done chan<- bool) { // only accepts a chan for sending
		fmt.Println("Worker working ...")
		time.Sleep(100*time.Millisecond)
		fmt.Println("Worker done")
		// test := <- done // will fail
		done <- true
	}

	done := make(chan bool, 1)
	go worker(done)

	<- done // Wait till worker is done
}
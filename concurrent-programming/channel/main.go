package main

// To remember
// ch <- v    // Send v to channel ch.
// v := <-ch  // Receive from ch, and assign value to v.

import "fmt"

type Sender struct {
	Content string
}

type Receiver struct {
	Content string
}

func main() {
	sCh := make(chan Sender, 1)
	rCh := make(chan Receiver, 1)

	msg := Sender{"Hello, Channel"}

	sCh <- msg

	select {
	case conn := <-sCh:
		fmt.Println("Sender", conn)
	case conn := <-rCh:
		fmt.Println("Receiver", conn)
	default:
		fmt.Println("No messages received.")
	}
}

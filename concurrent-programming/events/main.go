package main

import (
	"concurrent-programming/events/button"
	"fmt"
)

func main() {
	btn := button.New()

	chMouseUp := make(chan string, 1)
	chMouseDown := make(chan string, 1)

	btn.AddEventListener("onmouseup", chMouseUp)
	btn.AddEventListener("onmousedown", chMouseDown)

	go func() {
		for {
			msg := <-chMouseUp
			fmt.Println("onmouseup", msg)
		}
	}()

	go func() {
		for {
			msg := <-chMouseDown
			fmt.Println("onmousedown", msg)
		}
	}()

	btn.DispatchEvent("onmouseup", "Dispatched")
	btn.RemoveEventListener("onmouseup", chMouseUp)
	// btn.DispatchEvent("onmouseup", "Dispatched again")

	btn.DispatchEvent("onmousedown", "Dispatched")

	fmt.Scanln()
}

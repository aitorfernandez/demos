package main

import "fmt"

type Status struct {
	Code int32
	Body string
}

func main() {
	status := &Status{
		Code: 200,
		Body: "read",
	}

	ch := make(chan *Status)
	go update(status, ch)
	res := <-ch

	fmt.Println("res", res)
}

func update(s *Status, ch chan *Status) {
	s.Body = "approved"
	ch <- s
}

package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)

	ch := make(chan string, 1)
	go func() {
		for {
			// Waiting for messages
			msg, ok := <-ch
			if ok {
				fmt.Printf("%s", msg)
			} else {
				break
			}
		}
	}()

	mutex := make(chan bool, 1)
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			mutex <- true
			go func() {
				msg := fmt.Sprintf("%d + %d = %d\n", i, j, i+j)
				ch <- msg
				<-mutex
			}()
		}
	}

	// ch := make(chan bool, 1)
	// for i := 1; i < 10; i++ {
	// 	for j := 1; j < 10; j++ {
	// 		// Message into the channel and it can't accept any other messages until the message is retrieved.
	// 		ch <- true
	// 		go func() {
	// 			fmt.Printf("%d + %d = %d\n", i, j, i+j)
	// 			<-ch
	// 		}()
	// 	}
	// }

	// No running in parallel
	// mutex := new(sync.Mutex)
	// for i := 1; i < 10; i++ {
	// 	for j := 1; j < 10; j++ {
	// 		mutex.Lock()
	// 		go func() {
	// 			fmt.Printf("%d + %d = %d\n", i, j, i+j)
	// 			mutex.Unlock()
	// 		}()
	// 	}
	// }

	// Sharing memory across multiple threads with multiple CPUs.
	// for i := 1; i < 10; i++ {
	// 	for j := 1; j < 10; j++ {
	// 		go func() {
	// 			fmt.Printf("%d + %d = %d\n", i, j, i+j)
	// 		}()
	// 	}
	// }

	fmt.Scanln()
}

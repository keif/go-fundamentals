package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan) // only usable if we know it would execute last
}

func main() {
	//dones := make([]chan bool, 4)
	done := make(chan bool)

	//dones[0] = make(chan bool)
	go greet("Nice to meet you!", done)
	//dones[1] = make(chan bool)
	go greet("How are you?", done)
	//dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", done)
	//dones[3] = make(chan bool)
	go greet("I hope you're liking the course!", done)
	// go will only continue after this line of code
	// some data has come out of the channel

	// one <- done for each go routine, but not ideal/efficient
	//<-done
	//<-done
	//<-done
	//<-done
	// second approach
	//for _, done := range dones {
	//	<-done
	//}
	//for doneChan := range done {
	//fmt.Println(doneChan)
	//}
	for range done {
	}
}

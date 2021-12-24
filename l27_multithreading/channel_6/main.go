package main

import (
	"fmt"
	"time"
)

func newGoRoutine(done chan bool) {
	fmt.Println("Hey, I'm newGoRoutine and I'm going sleep!")
	time.Sleep(4 * time.Second)
	fmt.Println("newGoRoutine has awoken and is going to send data to channel!")
	done <- true
}

func main() {
	done := make(chan bool)
	fmt.Println("I'm main Goroutine and I wanna call newGoRoutine!")
	go newGoRoutine(done)
	<-done
	fmt.Println("Ok, main Goroutine has received data and will die!")
}

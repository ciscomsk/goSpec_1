package main

import (
	"fmt"
	"time"
)

func newGoRoutine() {
	fmt.Println("Hey, I'm new Goroutine!")
}

func main() {
	go newGoRoutine()
	// Тормозим выполнение горутины main => за это время горутина newGoroutine успеет проинициализироваться и выполниться.
	time.Sleep(1 * time.Second)

	fmt.Println("Main Goroutine work!")
}

package main

import (
	"fmt"
	"time"
)

func printEvenNumbers() {
	for i := 1000; i < 1020; i += 2 {
		time.Sleep(250 * time.Millisecond) // sleep - пример ручного выделения ресурсов
		fmt.Printf("%d ", i)
	}
}

func printOddNumbers() {
	for i := 1; i < 20; i += 2 {
		time.Sleep(450 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main() {
	go printEvenNumbers()
	go printOddNumbers()
	time.Sleep(5000 * time.Millisecond)

	fmt.Println("Main Goroutine has died.")
}

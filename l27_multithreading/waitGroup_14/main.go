package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup ~ счетчик горутин.
// Когда горутина запускается - WaitGroup++, когда завершается - WaitGroup-- =>
// => когда WaitGroup == 0 - все горутины отработали.

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("Goroutine", i, "has started")
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d has ended\n", i)
	wg.Done() // == WaitGroup--
}

func main() {
	no := 3
	var wg sync.WaitGroup

	for i := 0; i < no; i++ {
		wg.Add(1) // == WaitGroup++
		go process(i, &wg)
	}
	wg.Wait() // Пока WaitGroup != 0 - на этой строке будет блокировка.
	fmt.Println("All Goroutines have finished execution.")
}

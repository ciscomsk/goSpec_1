package main

// Если горутины не общаются - использовать мьютекс.
// Если общаются - можно и через каналы.

import (
	"fmt"
	"sync"
)

var x int = 0

func increment(wg *sync.WaitGroup, ch chan bool) {
	ch <- true // 1-я горутина отправит true, остальные заблокируются
	x += 1
	<-ch
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, ch)
	}

	w.Wait()
	fmt.Println("Final value of x:", x)
}

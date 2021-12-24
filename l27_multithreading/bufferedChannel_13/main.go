package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Successfully sending", i, "to ch")
	}
	close(ch)
	// err - panic: send on closed channel - запись в закрытый канал.
	//ch <- 10
}

// При заполнении буфера - канал блокируется до тех пор, пока не будет освобождено место.
// Длина канала - сколько сейчас элементов в канале - len(ch).
// Вместимость канала - какой размер буфера у канала - cap(ch).

// !!! Закрытие канала происходит только после вычитывания всех элементов.

func main() {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)

	for v := range ch {
		fmt.Println("Read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}
}

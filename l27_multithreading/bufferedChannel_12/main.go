package main

import "fmt"

// Буферизированный канал - канал с буфером, в который можно поместить несколько значений (в зависимости от capacity) до блокировки канала.
// Пример - запись в канал выполняется быстро, а чтение - медленно. ~Backpressure.
// ch := make(chan int, capacityIntValue), по умолчанию capacityIntValue == 1.

// Буферизированный канал уменьшает нагрузку на cpu, но увеличивает потребление памяти.

func main() {
	ch := make(chan string, 2) // capacity == 2
	ch <- "Bob"                // если бы канал был небуфиризированным - здесь был бы deadlock
	ch <- "Alex"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

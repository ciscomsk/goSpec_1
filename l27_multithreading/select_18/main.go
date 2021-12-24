package main

import "fmt"

// Select как инструмент защиты от deadlock.
// default - страхует от появления deadlock в ходе выполнения и берет работу на себя.

func main() {
	var ch chan string

	select {
	case v := <-ch: // здесь не выполняется чтение - просто проверка есть ли данные в канале
		fmt.Println("Received value:", v)
	default:
		fmt.Println("Default case has executed.") // если убрать default - будет deadlock
	}
}

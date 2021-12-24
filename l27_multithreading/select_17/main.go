package main

import (
	"fmt"
	"time"
)

// На практике select чаще всего используется для того, чтобы выполнить какие-то действия,
// пока в канал еще не пришли данные.

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "Process successful"
}

func main() {
	ch := make(chan string)
	go process(ch)

	// v1 - просто ждем здесь 10.5 с на блокировке.
	//v := <- ch

	// v2 - во время ожидания можно выполнять какие-либо полезные действия (например - пинг сервера пока не получим ответ 200/перебор серверов).
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("Received value: ", v)
			return
		default:
			fmt.Println("No value received")
		}
	}
}

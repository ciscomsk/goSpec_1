package main

import (
	"fmt"
	"time"
)

// Select - инструмент, позволяющий выбирать из множества канальных операций (чтение/запись) для множества каналов.
// Если из 10 каналов что-то пришло в 1 - select выбирает его.
// Если из 10 каналов что-то пришло сразу в 2 и более - select выбирает случайный.

// Пример использования - несколько апи (отдают одинаковые данные) - на разных серверах (разная нагруженность/производительность)
// - надо получить данные из самого быстрого.

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)

	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2: // будет выбран этот  кейс, т.к. в этот канал данные будут помещены быстрее.
		fmt.Println(s2)
	}
}

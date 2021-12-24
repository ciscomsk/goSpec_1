package main

import (
	"fmt"
	"time"
)

// Если данные поступают в несколько каналов одновременно - выбирается случайный.

func server1(ch chan string) {
	ch <- "from server1"
}

func server2(ch chan string) {
	ch <- "from server2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)

	time.Sleep(1 * time.Second)
	// Гарантии, что s1 будет просмотрен первым - нет. Обход каналов с готовыми для чтения значениями - выполняется случайным образом.
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
	// На выводе будет случайное значение - from server1 или from server2
}

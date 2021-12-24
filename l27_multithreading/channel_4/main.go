package main

import "fmt"

// Каналы - средство коммуникации между горутинами.
// Канал можно рассматривать как соединительную трубу между горутинами.
// Каналы по умолчанию имеют zeroValue == nil. Поэтому их создают через make.
// Канал - нереферальный тип (как и интерфейс).

func main() {
	// Канал, через который будут передаваться данные типа int.
	var a chan int

	if a == nil {
		fmt.Println("Channel is nil, lets' define it.")
		a = make(chan int)
		fmt.Printf("Channel type: %T\n", a)
	}
}

package main

import "fmt"

// Итерация по каналу и закрытие канала.

// На получателе можно использовать синтаксис:
// val, ok := <- ch
// val - значение, помещенное в канал отправителем.
// ok == true/false - в зависимости от того, открыт канал или уже закрыт отправителем.
// Если канал закрыт => в val будет помещено zeroValue для типа канала.

// !!! Best practice - закрывать канал при завершении работы с ним.

func generator(ch chan int) {
	for i := 0; i < 25; i++ {
		ch <- i
	}
	// Закрытие канала с посылкой zeroValue.
	close(ch)
}

func main() {
	ch := make(chan int)
	go generator(ch)

	// v1 - более читабельная.
	for {
		val, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Received from channel: ", val)
	}

	// v2 - упрощенная конструкция:
	//for val := range ch {
	//	fmt.Println("Received from channel: ", val)
	//}
}

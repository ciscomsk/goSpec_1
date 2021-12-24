package main

// Каналы могут иметь направление.
// Однонаправленные каналы не имеют различий в производительности с двунаправленными и служат для логического разделения кода.

// Двунаправленный канал - возможна и отправка и чтение.
// biDirectionalChan := make(chan int)
// Канал только на ОТПРАВКУ.
// sendChan := make(chan <- int)
// Канал только на ЧТЕНИЕ.
// readChan := make(<- chan int)

func sender(sendChan chan<- int) {
	sendChan <- 10
}

func main() {
	sendChan := make(chan<- int)
	go sender(sendChan)
	// err - Invalid operation: <- sendChan (receive from the send-only type chan<- int) - нельзя читать из канала, который объявлен каналом для отправки.
	//<- sendChan
}

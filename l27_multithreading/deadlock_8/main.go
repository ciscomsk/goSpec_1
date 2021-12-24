package main

// Deadlock - ситуация, когда кто-то пишет в канал, но из него НИКОГДА НИКТО НЕ ПРОЧИТАЕТ (отсутствует получатель)
// или когда кто-то читает из канала, но в него НИКТО НИКОГДА НЕ ЗАПИШЕТ (отсутствует отправитель).

func main() {
	ch := make(chan int)
	// err - fatal error: all goroutines are asleep - deadlock!
	ch <- 10
}

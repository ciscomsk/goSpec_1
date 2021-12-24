package main

import "fmt"

// Берем число - подсчитываем сумму квадратов и сумму кубов, затем сложим получившиеся значения.

func squaresGoRoutine(num int, squareChan chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit
		num /= 10
	}
	squareChan <- sum
}

func cubesGoRoutine(num int, cubeChan chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit * digit
		num /= 10
	}
	cubeChan <- sum
}

func main() {
	number := 456
	squareChan := make(chan int)
	cubeChan := make(chan int)

	go squaresGoRoutine(number, squareChan)
	go cubesGoRoutine(number, cubeChan)
	squaresSum, cubesSum := <-squareChan, <-cubeChan

	fmt.Println("Total result is: ", squaresSum+cubesSum)
}

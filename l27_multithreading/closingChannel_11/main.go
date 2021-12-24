package main

import "fmt"

// 5 горутин - calcSquares + digits + calcCubes + digits + main.

func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)

	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
	close(squareop)
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)

	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
	close(cubeop)
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)

	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)

	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output: ", squares+cubes)
}

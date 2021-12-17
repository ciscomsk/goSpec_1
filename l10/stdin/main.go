package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		age  int
		name string
	)

	// v1
	fmt.Scan(&age)
	fmt.Println(age)
	fmt.Println()

	// v2
	input := bufio.NewScanner(os.Stdin)
	// Захват потока ввода (до символа окончания строки) и сохранение в буфер. Scan() и Text() всегда идут последовательно.
	if input.Scan() {
		name = input.Text()
	}
	fmt.Println(name)
	fmt.Println()

	// Чтение с выходом на пустой строке.
	fmt.Println("For loop started:")
	for {
		if input.Scan() {
			result := input.Text() // input.Bytes() - прочитает байты
			if result == "" {
				break
			}
			fmt.Println("input string is:", result)
		}
	}

	// Преобразование строкового литерала к числовым типам.
	numStr := "10"
	numInt, _ := strconv.Atoi(numStr) // Atoi == Anything to Int (именно int)
	fmt.Printf("%d is %T\n", numInt, numInt)

	numint64, _ := strconv.ParseInt(numStr, 10, 64)
	fmt.Printf("%d is %T\n", numint64, numint64)

	// !!! ParseFloat всегда возвращает float64, если bitSize == 32 - это просто гарантия, что результат без проблем можно сконвертировать в float32.
	numfloat32, _ := strconv.ParseFloat(numStr, 32)
	fmt.Printf("%.2f is %T\n", numfloat32, numfloat32)
	fmt.Println()

}

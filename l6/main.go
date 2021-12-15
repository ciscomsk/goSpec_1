package main

import (
	"fmt"
	"strings"
)

// go run l6/main.go
func main() {
	// init - блок инициализации переменных цикла.
	// В качестве init может быть использовано выражение присваивания ТОЛЬКО через :=
	// condition - условие (если верно - тело цикла выполняется, если нет - цикл завершается)
	// post - изменение переменной цикла (инкремент/декремент)
	//for init; condition; post {
	// тело цикла
	//}

	// i++ - постфиксный инкремент == i += 1. Префиксного инкремента в go нет.
	for i := 0; i <= 5; i++ {
		fmt.Printf("Current value: %d\n", i)
	}
	fmt.Println()

	// err - Unresolved reference 'i' - переменная недоступна вне цикла.
	//fmt.Println(i)

	// break - команда, прерывающая текущую итерацию цикла и передающая управление инструкция, следующим за циклом.
	for i := 0; i <= 15; i++ {
		if i > 12 {
			break
		}
		fmt.Printf("current value: %d\n", i)
	}
	fmt.Println("After for loop with BREAK.")
	fmt.Println()

	// continue - команда, прерывающая текущую итерацию цикла и передающая управление СЛЕДУЮЩЕЙ итерации цикла.
	for i := 0; i <= 20; i++ {
		if i > 10 && i <= 15 {
			continue
		}
		fmt.Printf("Current value: %d\n", i)
	}
	fmt.Println("After for loop with CONTINUE")
	fmt.Println()

	// Вложенные циклы.
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()

	// Лейблы - синтаксический сахар.
	var stopFlag bool

	for i := 0; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i: %d, j: %d, i+j: %d\n", i, j, i+j)
			if i == j {
				// Хотим выйти из всех циклов (т..е из внешнего тоже) - но выходим только из внутреннего.
				stopFlag = true
				break
			}
		}

		if stopFlag {
			break
		}
	}
	fmt.Println()

	// => с лейблом проще
outer:
	for i := 0; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i: %d, j: %d, i+j: %d\n", i, j, i+j)
			if i == j {
				break outer
			}
		}
	}
	fmt.Println()

	// Модификации цикла for.
	// 1. Цикл while do.
	var loopVar int

	// ; - можно не писать - опускает блоки init/post.
	//for ; loopVar < 10; {
	// ==
	for loopVar < 10 {
		fmt.Printf("In while like loop %d\n", loopVar)
		loopVar++
	}
	fmt.Println()

	// 2. Бесконечный цикл.
	var password string
	for {
		fmt.Print("Insert password: ")
		fmt.Scan(&password)

		if strings.Contains(password, "1234") {
			fmt.Println("Weak password. Try again.")
			//continue
		} else {
			fmt.Println("Password accepted.")
			// ??? В более ранних версиях break завершал условный оператор, но не цикл - требовалось добавить лейбл для цикла или return.
			break
		}
	}
	fmt.Println()

	// 3. Цикл с множественными переменными.
	for x, y := 0, 1; x <= 10 && y <= 12; x, y = x+1, y+2 {
		fmt.Printf("%d + %d = %d\n", x, y, x+y)
	}

}

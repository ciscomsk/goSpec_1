package main

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strconv"
)

// Механизмы сигнализации об аномальном поведении:
// 1. Error - каноничное исполнение на случай ненормального поведения (~исключение).
// Ошибки приведения/подключение к бд/наличие файла - т.е. в случаях, где возможен альтернативный вариант поведения.
// 2. Panic - не лучший вариант т.к. приводит к аварийному завершению , в принципе мог быть обычной ошибкой.
// Выполнение программы дальше продолжать нельзя.
// Выход за границы массива/обработка пустых интерфейсов и т.д.

func funcWithError(a int) (string, error) {
	if a%2 == 0 {
		return "Even", nil
	}
	return "", errors.New("this is Error")
}

func panicExample(firstName *string, lastname *string) {
	// Обработка паники - deferred функции вызываются до завершения программы.
	defer PanicRecover()

	if firstName == nil {
		panic("Runtime error: firstname can't be nil!")
	}

	if lastname == nil {
		panic("Runtime error: lastname can't be nil!")
	}

	// !!! После обработки паники - в функцию, вызвавшую панику, вернуться уже нельзя.

	fmt.Println("Full name:", *firstName, *lastname)
}

func PanicRecover() {
	if r := recover(); r != nil {
		fmt.Println("Panic satisfied: ", r)
		// Для вывода стектрейса.
		debug.PrintStack()
	}
}

func main() {
	numLiteral := "12"

	num, err := strconv.Atoi(numLiteral)
	if err != nil {
		// v1
		//log.Fatal("Can't convert this value to int: ", err)
		// ==
		// v2 - более каноничный вариант.
		fmt.Println("Can't convert this value to int:", err)
		return
	}
	fmt.Println("Conversion done: ", num)
	fmt.Println()

	ans, err := funcWithError(4)
	if err != nil {
		fmt.Println("Not use odd values.", err)
		return
	}
	fmt.Println(ans)
	fmt.Println()

	var name = "Bob"
	panicExample(&name, nil)
}

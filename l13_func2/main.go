package main

import "fmt"

// !!! Функции в go являются экземпляром 1-го уровня => функцию можно присваивать в переменную,
// передавать в качестве параметра, возвращать из других функций.

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}

// Использование анонимных функций внутри именованной.
func calcAndReturnResult(command string, a int, b int) int {
	if command == "addition" {
		return func(a, b int) int { return a + b }(a, b)
	} else if command == "subtraction" {
		return func(a, b int) int { return a - b }(a, b)
	} else {
		return func(a, b int) int { return a * b }(a, b)
	}
}

// Возврат функции в качестве результата.
func calcAndReturnFunc(command string) func(a, b int) int {
	if command == "addition" {
		return func(a, b int) int { return a + b }
	} else if command == "subtraction" {
		return func(a, b int) int { return a - b }
	} else {
		return func(a, b int) int { return a * b }
	}
}

// Функция как параметр в другой функции.
func receiveFuncAndReturnValue(f func(a, b int) int) int {
	var intA, intB int = 100, 200
	return f(intA, intB)
}

func main() {
	var command string
	command = "addition"

	res := calcAndReturnResult(command, 10, 20)
	fmt.Println("Result with command:", command, "value", res)
	resFunc := calcAndReturnFunc(command) // resFunc - функция.
	fmt.Println("Result with command:", command, "value", resFunc(10, 20))
	fmt.Println()

	// !!! Тип функции в go определяется входными и выходными параметрами.
	fmt.Printf("Type of resFunc is: %T\n", resFunc)
	fmt.Printf("Type pf calcAndReturnFunc: %T\n", calcAndReturnFunc)
	fmt.Println()

	fmt.Println("ReceiveFuncAndReturnValue(add):", receiveFuncAndReturnValue(add))
	fmt.Println(receiveFuncAndReturnValue(func(a, b int) int { // определяем анонимную функцию.
		return a*a + 2*a*b + b*b
	}))

}

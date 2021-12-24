package main

import "fmt"

// defer - оператор отложенного вызова функции/метода. Используется для обработка ошибок/закрытие файла или соединения с бд.

func CheckDBCloseConnection() {
	fmt.Println("Check started...")
	fmt.Println("Check done. Connection closed.")
}

// Если функция/метод принимает входные параметры и используется в  связке с defer -
// !!! параметры рассчитываются в момент их передачи в функцию/метод.
// Вызов функции осуществляется в момент завершения вышележащей функции.

func CheckDBCloseConnectionWithArg(a int) {
	fmt.Println("Check started... Value: ", a)
	fmt.Println("Check done. Connection closed.")
}

func OuterFunc() int {
	defer fmt.Println("I'm deferred print function.") // отработает перед return
	fmt.Println("OuterFunc has started.")
	var result = 10
	fmt.Println("OuterFunc has finished. Ready to return value.")

	return result
}

func main() {
	//defer CheckDBCloseConnection()

	var numIp = 10
	p := &numIp

	//defer CheckDBCloseConnectionWithArg(numIp)	// 10
	*p++

	fmt.Println("Value numIp in main(): ", numIp) // 11
	fmt.Println("Main function started.")
	fmt.Println("Main function ended.")
	fmt.Println()

	fmt.Println("Value from OuterFunc on main side is:", OuterFunc())
	fmt.Println()

	// Выполнится в порядке стека - 4 => 3 ...
	defer fmt.Println("Step 1 defer")
	defer fmt.Println("Step 2 defer")
	defer fmt.Println("Step 3 defer")
	defer fmt.Println("Step 4 defer")
}

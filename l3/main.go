package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		age  int
		name string
	)

	// & - указатель/ссылка.

	// Чтение из stdin.
	//fmt.Scan(&age)
	//fmt.Scan(&name)
	// => валидные разделители - \n и пробел.
	fmt.Scan(&age, &name)

	// Если будет введено значение неправильного типа - будет "нулевая" инициализация.
	fmt.Printf("My name is: %s\nMy age is: %d\n", name, age)

	// Ручное указание потока ввода.
	fmt.Fscan(os.Stdin, &age)
	fmt.Println("New age:", age)
}

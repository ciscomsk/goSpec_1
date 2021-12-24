package main

import "fmt"

type Employee struct {
	name   string
	salary int
}

// Метод, в котором получатель копируется и в его теле происходит работа с локальной копией (ValueReceiver).
func (e Employee) SetName(newName string) {
	e.name = newName
}

// Метод, в котором получатель передается по ссылке - в теле метода работает с ссылкой на экземпляр (PointerReceiver).
func (e *Employee) SetSalary(newSalary int) {
	e.salary = newSalary
}

// Методы с PointerReceiver используются в случае:
// 1. Изменения, выполняемые методом над экземпляром должны быть видны на вызывающей стороне.
// 2. Когда экземпляр имеет большой размер => его копирование дорого с точки зрения расходов ресурсов.
// 3. Любой вид экземпляра (значение/ссылка) может работать с такими методами - 33 строка.

// Best practice - делать все методы с PointerReceiver - оптимизация производительности.

func main() {
	e := Employee{"Alex", 3000}
	fmt.Println("Before setting parameters:", e)

	// !!! Никаких изменений не будет - т.к. меняется копия в функции.
	e.SetName("Bob")
	fmt.Println("After SetName call:", e)

	(&e).SetSalary(4500) // e.SetSalary == (&e).SetSalary - go неявно выполнит приведение.
	fmt.Println("After SetSalary call:", e)
	fmt.Println()
}

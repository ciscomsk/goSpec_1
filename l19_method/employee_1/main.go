package main

import "fmt"

// Метод - функция, привязанная к определенной структуре.
//func (s Struct) MethodName(parameters type) type {}
// (s Struct) - receiver, получатель метода.

type Employee struct {
	name     string
	position string
	salary   int
	currency string
}

func (e Employee) DisplayInfo() {
	fmt.Println("Name:", e.name)
	fmt.Println("Position:", e.position)
	fmt.Printf("Salary: %d %s\n", e.salary, e.currency)
}

func main() {
	emp := Employee{
		name:     "Bob",
		position: "Senior Golang developer",
		salary:   3000,
		currency: "USD",
	}

	// Вызов метода.
	emp.DisplayInfo()
}

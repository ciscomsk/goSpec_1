package main

import "fmt"

type University struct {
	city string
	name string
}

// !!! В структуру Professor встроены поля структуры University => все методы University можно вызывать на Professor == НАСЛЕДОВАНИЕ.
type Professor struct {
	fullName string
	age      int
	University
}

// Данный метод явно связан только с структурой University.
func (u *University) FullInfoUniversity() {
	fmt.Printf("Uni name: %s and City: %s\n", u.name, u.city)
}

func main() {
	p := Professor{
		fullName: "Boris Bobroff",
		age:      150,
		University: University{
			city: "Moscow",
			name: "BMSTU",
		},
	}
	// !!! Вызов метода структуры University через экземпляр структуры Professor.
	p.FullInfoUniversity()
}

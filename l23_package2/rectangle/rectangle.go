package rectangle

import "fmt"

// Инициализация переменных уровня пакета.
var (
	name string = "John"
	age  int    = 33
)

func init() {
	fmt.Println("Init function from rectangle package!")
	fmt.Println("Name:", name, "Age:", age)
}

type Rectangle struct {
	A, B int
}

func New(newA, newB int) *Rectangle {
	return &Rectangle{newA, newB}
}

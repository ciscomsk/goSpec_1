package main

import (
	"fmt"
	"math"
)

// Преимущества методов над функциями:
// 1. Наличие методов улучшает "консистентность" кода, т.к. напрямую влияет на его понимание.
// 2. В go создание функций с одинаковыми названиями - запрещено, созданием методов с одинаковыми названиями (для разных получателей) - разрешено.
// 3. В метод с value receiver можно передать экземпляр или ссылку, в функцию  с value receiver - только экземпляр (см PointerVersusValue_5).

type Circle struct {
	radius float64
}

type Rectangle struct {
	length, width int
}

// Функция и метод для структуры Circle.

func Perimeter(c Circle) float64 {
	return c.radius * 2 * math.Pi
}

func (c Circle) Perimeter() float64 {
	return c.radius * 2 * math.Pi
}

// Функция и метод для структуры Rectangle.

// err - 'Perimeter' redeclared in this package. Как выход - давать разные имена функциям - perimeterCircle/perimeterRectangle.
//func Perimeter(r Rectangle) int {
//	return 2 * (r.length + r.width)
//}

func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func main() {
	c := Circle{10.5}
	fmt.Println("Call function:", Perimeter(c))
	fmt.Println("Call method:", c.Perimeter())
	fmt.Println()

	r := Rectangle{10, 20}
	fmt.Println("Call method for rectangle:", r.Perimeter())
}

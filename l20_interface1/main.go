package main

import "fmt"

// Структура - явно декларированный заименованный набор атрибутов.
// Структура отвечает на вопрос - из ЧЕГО я должен состоять, чтобы считаться ТИПОМ, который декларирует структуру.
// Структура описывает паттерн СОСТОЯНИЕ (из чего состоит).

// Интерфейс - явно декларированный набор сигнатур ПОВЕДЕНИЙ (чаще всего в виде набора методов), удовлетворив который,
// можно считаться типом, который декларирует интерфейс (КОНТРАКТ).
// Интерфейсы в go декларируют полу-абстрактный тип (абстрактный тип - тип, объекты которого создать невозможно).
// Интерфейс отвечает на вопрос - что я должен УМЕТЬ ДЕЛАТЬ, чтобы считаться тем ТИПОМ, который декларирует интерфейс.
// Интерфейс - описывает паттерн ПОВЕДЕНИЕ.

// Объявление интерфейса.
type FigureBuilder interface {
	// Набор сигнатур методов, которые необходимо реализовать в структуре-претенденте.
	Area() int
	Perimeter() int
}

//func (fb FigureBuilder) GetInfo() {
//
//}

// Претенденты.

// Rectangle удовлетворяет условиям интерфейса (реализует интерфейс) FigureBuilder.
type Rectangle struct {
	length, width int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.length * r.width)
}

type Circle struct {
	radius int
}

// Circle реализует интерфейс FigureBuilder.
func (c Circle) Area() int {
	return 3 * c.radius * c.radius
}

func (c Circle) Perimeter() int {
	return 2 * 3 * c.radius
}

func main() {
	r1 := Rectangle{10, 10}
	c1 := Circle{5}

	totalArea := r1.Area() + c1.Area()
	fmt.Println(totalArea)
	fmt.Println()

	// Если экземпляров будет много - их нужно будет хранить в слайсах (для удобства расчета площади).
	r2 := Rectangle{30, 50}
	r3 := Rectangle{1, 1}
	c2 := Circle{10}
	c3 := Circle{15}

	rectangles := []Rectangle{r1, r2, r3}
	circles := []Circle{c1, c2, c3}

	// !!! ДУБЛИРОВАНИЕ КОДА.
	totalAreaOfRectangles := 0
	for _, rect := range rectangles {
		totalAreaOfRectangles += rect.Area()
	}

	totalAreaOfCircles := 0
	for _, circ := range circles {
		totalAreaOfCircles += circ.Area()
	}

	fmt.Println("Total area is:", totalAreaOfRectangles+totalAreaOfCircles)
	fmt.Println()

	// => интерфейс
	// Слайс экземпляров, удовлетворяющих интерфейсу FigureBuilder.
	figures := []FigureBuilder{r1, r2, r3, c1, c2, c3}

	// !!! ДУБЛИРОВАНИЯ КОДА - НЕТ.
	totalAreaOfFB := 0
	for _, fb := range figures {
		totalAreaOfFB += fb.Area()
	}
	fmt.Println("Total area of FB is:", totalAreaOfFB)

}

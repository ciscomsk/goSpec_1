package main

import "fmt"

type PerimeterCalculator interface {
	Perimeter() int
}

type AreaCalculator interface {
	Area() int
}

type Rectangle struct {
	a, b  int
	color string
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.a + r.b)
}

func (r Rectangle) Area() int {
	return r.a * r.b
}

// Объединение нескольких интерфейсов (создание единого уровня абстракции).
type FigureFeatureCalculator interface {
	PerimeterCalculator // встраивание интерфейса
	AreaCalculator

	// == наследование инткрфейсов в Java
}

func main() {
	rect := Rectangle{10, 20, "green"}
	var pCalc PerimeterCalculator
	var aCalc AreaCalculator

	fmt.Printf("Value: %v and Type: %T\n", pCalc, pCalc)
	pCalc = rect
	fmt.Printf("Value: %v and Type: %T\n", pCalc, pCalc)
	fmt.Println(pCalc.Perimeter())

	aCalc = rect
	fmt.Println(aCalc.Area())
	fmt.Println()

	var ffCalc FigureFeatureCalculator
	ffCalc = rect
	fmt.Println(ffCalc.Perimeter())
	fmt.Println(ffCalc.Area())
}

package main

import "fmt"

type Rectangle struct {
	length, width int
}

// Реализуем функцию и метод для подсчета периметра прямоугольника.
// !!! Все параметры передаем как копии.

// Value receiver.
// !!! Функцию с value receiver можно вызвать только на экземпляре, на ссылке - нельзя.
func Perimeter(r Rectangle) int {
	return 2 * (r.length + r.width)
}

// !!! При вызове метода с value receiver неважно, что мы передаем - экземпляр или ссылку.
func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

// Метод, который меняет поле структуры на новое с value receiver.
// !!!  Изменений на вызывающей стороне не будет - не важно, что было передано - экземпляр или ссылка - изменяться будет только локальная копия в теле метода.
func (r Rectangle) SetLength(newLength int) {
	r.length = newLength
}

// Pointer receiver.
// !!!  Функцию с pointer receiver можно вызвать только на ссылке, на экземпляре - нельзя.
func Area(r *Rectangle) int {
	return r.length * r.width
}

// !!! При вызове метода с pointer receiver неважно, что мы передаем - экземпляр или ссылку.
func (r *Rectangle) Area() int {
	return r.length * r.width
}

// !!! Изменения на вызывающей стороне - будут не важно, что было передано - экземпляр или ссылка.
func (r *Rectangle) setWidth(newWidth int) {
	r.width = newWidth
}

func main() {
	rectangleAsValue := Rectangle{10, 10}
	fmt.Println("Call function for rectangleAsValue:", Perimeter(rectangleAsValue))
	fmt.Println("Call method for rectangleAsValue:", rectangleAsValue.Perimeter())
	fmt.Println()

	rectangleAsPointer := &rectangleAsValue
	// err - Cannot use 'rectangleAsPointer' (type *Rectangle) as the type Rectangle
	//fmt.Println("Call function for rectangleAsPointer:", Perimeter(rectangleAsPointer))
	fmt.Println("Call method for rectangleAsPointer:", rectangleAsPointer.Perimeter())
	fmt.Println()

	// Method with value receiver.
	// Вызов метода SetLength у экземпляра Rectangle.
	fmt.Println("Before call method SetLength:", rectangleAsValue)
	rectangleAsValue.SetLength(9999)
	// Значение не изменится т.к. метод с value receiver - меняется локальная копия в методе.
	fmt.Println("After call method SetLength:", rectangleAsValue)

	// Вызов метода SetLength у ссылки Rectangle.
	rectangleAsPointer.SetLength(99999999)
	// Значение не изменится т.к. метод с value receiver - меняется локальная копия в методе.
	fmt.Println("After call method SetLength for &Rectangle:", rectangleAsPointer)
	fmt.Println()
	fmt.Println()

	fmt.Println("Call Area function for &Rectangle:", Area(rectangleAsPointer))
	fmt.Println("Call Area method for &Rectangle:", rectangleAsPointer.Area())
	fmt.Println()

	// err - Cannot use 'rectangleAsValue' (type Rectangle) as the type *Rectangle
	//fmt.Println("Call Area function for &Rectangle:", Area(rectangleAsValue))
	fmt.Println("Call Area method for Rectangle:", rectangleAsValue.Area())
	fmt.Println()

	// Method with pointer receiver.
	// Вызов метода SetWidth у ссылки Rectangle.
	fmt.Println("Before changing width:", rectangleAsValue)
	rectangleAsPointer.setWidth(8888)
	// Значение изменится.
	fmt.Println("After change via method SetWidth for &Rectangle", rectangleAsValue)

	// Вызов метода SetWidth у экземпляра Rectangle.
	rectangleAsValue.setWidth(88888888)
	// Значение изменится.
	fmt.Println("After change via method SetWidth for Rectangle:", rectangleAsValue)

}

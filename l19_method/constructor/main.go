package main

import "fmt"

type Rectangle struct {
	length, width int
}

// В go принятно возвращать из конструктора не сам экземпляр, а ссылку на него.
// В конструкторе можно реализовать валидацию параметров (нужные проверки).
func New(newLength, newWidth int) *Rectangle {
	return &Rectangle{newLength, newWidth}
}

func (r *Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func (r *Rectangle) Area() int {
	return r.length * r.width
}

func main() {
	r := New(10, 20)
	fmt.Printf("Type: %T and value: %v\n", r, r)
	fmt.Println("Perimeter:", r.Perimeter())
	fmt.Println("Area:", r.Area())
}

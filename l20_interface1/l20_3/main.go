package main

import "fmt"

// Пустой интерфейс - нет требований к поведению.
// !!! Любой тип данных реализует пустой интерфейсу.
type Empty interface {
}

func Describer(pretender interface{}) { // interface{} == Empty
	fmt.Printf("Type: %T and value: %v\n", pretender, pretender)
}

type Student struct {
	name string
}

// Type Assertion.
func SemiGeneric(pretender interface{}) {
	val, ok := pretender.(int)
	fmt.Println("Value:", val, "Ok?:", ok)
}

func main() {
	strSample := "Hello world!"
	// передача параметра без присваивания в промежуточную переменную.
	Describer(strSample)

	intSample := 200
	Describer(intSample)

	Describer(Student{"Vova"})
	fmt.Println()

	SemiGeneric(10)
	SemiGeneric(Student{"Fedya"})
	SemiGeneric("Hello world")

}

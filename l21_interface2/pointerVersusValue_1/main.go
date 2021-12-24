package main

import "fmt"

// В Go принятно называть интерфейсы с окончанием "er" (Describer/Worker/Pooler).
type Describer interface {
	Describe()
}

type Student struct {
	name string
	age  int
}

func (std Student) Describe() {
	fmt.Printf("Student name: %s and age: %d\n", std.name, std.age)
}

type Professor struct {
	name string
	age  int
}

func (prof *Professor) Describe() {
	fmt.Printf("Professor nameL %s and age: %d\n", prof.name, prof.age)
}

func main() {
	var descr1 Describer

	stud1 := Student{"Alex", 23}
	descr1 = stud1 // экземпляр реализует интерфейс Describer - func (std Student) Describe()
	descr1.Describe()

	stud2 := &Student{"Bob", 21}
	descr1 = stud2 // Ссылка разыменовывается компилятором
	descr1.Describe()
	fmt.Println()

	var descr2 Describer
	prof1 := &Professor{"Viktor Wann", 72}
	descr2 = prof1 // ссылка реализует интерфейс Describer - func (prof *Professor) Describe()
	descr2.Describe()

	prof2 := Professor{"Bob Brown", 65}
	// err - Type does not implement 'Describer' as the 'Describe' method has a pointer receiver
	//descr2 = prof2	// экземпляр не реализует интерфейс Describer
	// !!! Интерфейс не является референсным типом =>
	// при работе с методом - компилятор может взять ссылку (референс) на экземпляр
	// но при попытке сделать это через интерфейс - компилятор не видит в нем никаких ссылок
	// Решение - создавать конструкторы, которые возвращают указатели.
	fmt.Println(prof2)
}

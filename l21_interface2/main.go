package main

import "fmt"

type Describer interface {
	Describe()
}

type Student struct {
	name string
	age  int
}

func (std Student) Describe() {
	fmt.Printf("%s and %d y.o.\n", std.name, std.age)
}

func TypeFinder(i interface{}) {
	// В переменную v присваиваем значение.
	// !!! Без присвоения - i.Describe() - работать не будет, т.к. тип кастится при проверке с присваиванием значения.
	switch v := i.(type) {
	case string:
		fmt.Println("This is  string")
	case int:
		fmt.Println("This is int")
	// С типом интерфейса можно сравниваться, как и с любыми другими типами в языке => интерфейсы - полу-абстрактные типы.
	case Describer:
		fmt.Println("I'm describer")
		v.Describe()
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	std := Student{"Alex", 23}

	TypeFinder(10)
	TypeFinder("Hello")
	TypeFinder(nil)
	TypeFinder(std)
}

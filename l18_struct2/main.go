package main

import "fmt"

// Вложенные структуры - использование структуры, как тип поля в другой структуре.
type University struct {
	yearBased int
	infoShort string
	infoLong  string
}

type Student struct {
	firstName  string
	lastName   string
	university University
}

// Встроенные структуры - добавление полей одной структуры к другой. В go через этот механизм реализовано наследование.
type Professor struct {
	firstName string
	lastName  string
	age       int
	greatWork string
	//papers map[string]string	// для примера сравнения с сыслочными типами данных
	University // добавление всех полей структуры University в Professor
}

func main() {
	// Создание экземпляра с вложенной структурой.
	stud := Student{
		firstName: "Fedya",
		lastName:  "Petrov",
		university: University{
			yearBased: 1991,
			infoShort: "cool university",
			infoLong:  "very cool University",
		},
	}

	// Обращение к вложенным полям структуры.
	fmt.Println("YearBased:", stud.university.yearBased)
	fmt.Println("InfoLong:", stud.university.infoLong)
	fmt.Println()

	// Создание экземпляра с встраиванием структур.
	prof := Professor{
		firstName: "Anatoly",
		lastName:  "Smirnov",
		age:       125,
		greatWork: "Ultimate C programming",
		University: University{
			yearBased: 1734,
			infoShort: "short info",
			// err unknown field 'age' in struct literal of type University.
			// ??? раньше так было можно.
			//age: 2021 - 1734,

		},
	}

	// Обращение к полям с встроенной структурой.
	fmt.Println("YearBased:", prof.yearBased)
	fmt.Println("InfoShort:", prof.infoShort)
	fmt.Println()

	// Сравнение экземпляр на равенство. При сравнение экземпляров - происходит сравнение полей.
	// !!! Если хотя бы одно из полей структур не сравнимо (ссылочные типы - слайс/мапа) - вся структура не сравнима.
	profLeft := Professor{}
	profRight := Professor{}
	fmt.Println(profLeft == profRight)
}

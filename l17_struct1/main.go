package main

import "fmt"

// Структура - заименованный набор полей (свойств/состояний/атрибутов), определяющий новый тип данных.
// Класс дополнительно к этому определяет поведение, структура - нет. Структура == класс, состоящий из атрибутов.

// Явное определение структуры.
type Student struct {
	firstName string
	lastName  string
	age       int
}

// Если есть несколько состояний одного типа, можно определять так:
type Student2 struct {
	firstName, lastName, groupName string
	age, courseNumber              int
}

func prettyPrint(std Student) {
	fmt.Println("====")
	fmt.Println("FirstName:", std.firstName)
	fmt.Println("LastName", std.lastName)
	fmt.Println("Age:", std.age)
}

func main() {
	// Создание экземпляра структуры - поля можно задавать в произвольном порядке.
	stud1 := Student{
		firstName: "Fedya",
		age:       21,
		lastName:  "Petrov",
	}
	fmt.Println("Stud1:", stud1)
	fmt.Println()

	prettyPrint(stud1)
	fmt.Println()

	// Более короткий вариант создания структуры - поля должны указываться в порядке заданном в структуре.
	stud2 := Student{"Petya", "Ivanov", 19}
	prettyPrint(stud2)
	fmt.Println()

	// Не определенные полябудут заполнены "нулевыми" значениями типов.
	stud3 := Student{
		firstName: "Vasya",
	}
	prettyPrint(stud3)
	fmt.Println()

	// Анонимные структуры - структура без имени. Часто используются для отладки.
	anonStudent := struct {
		age           int
		groupID       int
		professorName string
	}{
		age:           23,
		groupID:       2,
		professorName: "Alexeev",
	}
	fmt.Println("AnonStudent", anonStudent)
	fmt.Println()

	// Доступ к состояниям и их модификация.
	studVova := Student{"Vova", "Ivanov", 19}
	fmt.Println("firstName:", studVova.firstName)
	fmt.Println("lastName:", studVova.lastName)
	fmt.Println("age:", studVova.age)
	fmt.Println()

	studVova.age += 2
	fmt.Println("new age:", studVova.age)
	fmt.Println()

	// Инициализация пустой структуры.
	// v1
	emptyStudent := Student{}
	// v2
	var emptyStudent2 Student
	prettyPrint(emptyStudent)
	prettyPrint(emptyStudent2)
	fmt.Println()

	// Указатели на экземпляры структур.
	studPointer := &Student{
		firstName: "Igor",
		lastName:  "Sidorov",
		age:       22,
	}
	fmt.Println("Value studPointer:", studPointer) // == &{Igor Sidorov 22}, & == указатель.

	secondPointer := studPointer
	(*secondPointer).age += 20
	fmt.Println("Value after pointer modification:", studPointer)

	studPointerNew := new(Student)
	fmt.Println(studPointerNew)
	fmt.Println()

	// Обращение к полям структуры через указатель.
	// v1
	fmt.Println("Age via (*...).age", (*secondPointer).age)
	// v2 - более удобный способ
	fmt.Println("Age via .age", studPointer.age) // неявное разыменование указателя и запрос соответствующего поля
	fmt.Println()

	// Структура с анонимными полями.
	type Human1 struct {
		string
		int
		bool
	}

	type Human2 struct {
		firstName string
		lastName  string
		string
		int
		bool
	}

	human := &Human2{
		firstName: "Bob",
		lastName:  "Johnson",
		string:    "Additional Info",
		int:       -1,
		bool:      true,
	}
	fmt.Println(human)
	fmt.Println("Anon field string:", human.string)

}

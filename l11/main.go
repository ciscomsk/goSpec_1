package main

import "fmt"

func main() {
	// Map - набор пар key:value.

	//Инициализация пустой мапы.
	mapper := make(map[string]int)
	fmt.Println("Empty map:", mapper)

	// Добавление пар в существующую мапу.
	mapper["Alice"] = 24
	mapper["Bob"] = 25
	fmt.Println("Mapper after adding pairs:", mapper)
	fmt.Println()

	// Инициализация мапы с указанием пар.
	newMapper := map[string]int{
		"Alice": 1000,
		"Bob":   1000,
	}

	newMapper["Jo"] = 3000
	fmt.Println("NewMapper:", newMapper)
	fmt.Println()

	// !!! Особенности мапы.
	// 1. Мапа в go - не упорядочена.
	// 2. Ключом в мапе может быть только сравнимый тип (==, !=).
	// 3. Нулевое значение мапы == nil. Для создания пустой мапы - использовать make.
	var mapZeroValue map[string]int // mapZeroValue == nil - указатель в пустоту. Используется для сравнения/в тестах.
	// err - panic: assignment to entry in nil map
	//mapZeroValue["Alice"] = 12
	fmt.Println(mapZeroValue)
	fmt.Println()

	// Получение элементов из map.
	// Получение элемента, который представлен в мапе.
	testPerson := "Alice"
	fmt.Println("Salary of testPerson:", newMapper[testPerson])
	// Получение элемента, который не представлен в мапе - будет получено "нулевое" значение для значения.
	testPerson = "Derek"
	fmt.Println("Salary of new testPerson:", newMapper[testPerson])
	fmt.Println(newMapper)
	fmt.Println()

	// Проверка вхождения ключа.
	employee := map[string]int{
		"Den":   0,
		"Alice": 0,
		"Bob":   0,
	}
	// !!! Значения для ключей с нулевым значением не отличимы от значений для отсутствующих ключей.
	fmt.Println("Den and value:", employee["Den"])
	fmt.Println("Jo and value:", employee["Jo"])
	fmt.Println()

	// При обращении по ключу - возвращается 2 значения.
	if value, ok := employee["Den"]; ok {
		fmt.Println("Den and value:", value)
	} else {
		fmt.Println("Den doesn't exists in map")
	}

	if value, ok := employee["Jo"]; ok {
		fmt.Println("Jo and value:", value)
	} else {
		fmt.Println("Jo doesn't exists in map")
	}
	fmt.Println()

	// Итерация по элементам мапы.
	for key, value := range employee {
		fmt.Printf("%s and value %d\n", key, value)
	}
	fmt.Println()

	// Удаление пар.
	// Удаление существующей пары.
	fmt.Println("Before deleting:", employee)
	delete(employee, "Den")
	fmt.Println("After deleting:", employee)
	// !!! Удаление несуществующей пары - не будет ошибки. Заранее надо убедиться, что пара существует. delete - очень дорогая операция.
	if _, ok := employee["Anna"]; ok {
		delete(employee, "Anna")
	}
	fmt.Println("After second deleting:", employee)
	fmt.Println()

	// Длина мапы == количество пар.
	fmt.Println("Pair amount in map:", len(employee))

	// !!! Мапа (как и слайс) - ссылочный тип.
	words := map[string]string{
		"One": "Один",
		"Two": "Два",
	}
	fmt.Println()

	newWords := words
	newWords["Three"] = "Три"
	delete(newWords, "One")
	fmt.Println("words map:", words)
	fmt.Println("newWords map:", newWords)
	fmt.Println()

	// Сравнение массивов - ок. Массив можно использовать как ключ мапы.
	if [3]int{1, 2, 3} == [3]int{1, 2, 3} {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}

	// Сравнение слайсов - нельзя. Т.к. тип ссылочный можно сравнить только с nil.
	// err - Invalid operation: []int{1, 2, 3} == []int{1, 2, 3} (the operator == is not defined on []int)
	//if []int{1, 2, 3} == []int{1, 2, 3} {
	//}

	// Сравнение мап - нельзя. Т.к. тип ссылочный можно сравнить только с nil.
	aMap := map[string]int{
		"a": 1,
	}
	bMap := map[string]int{
		"b": 1,
	}
	//if aMap == bMap {
	//}

	if aMap == nil {
		fmt.Println("Zero value map")
	}
	if bMap == nil {
		fmt.Println("Zero value of map bMap")
	}

	// !!! Если мапа/слайс является составляющей какой-либо структуры => структура автоматически становится несравнимой.
	// Строки это тоже слайс байт, но строки - базовый встроенный тип с переопределенным сравнением - "нулевое" значение
	// для них - пустой набор байт, а не nil.
}

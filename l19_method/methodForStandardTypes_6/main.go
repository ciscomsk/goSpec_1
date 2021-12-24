package main

import "fmt"

// Методы для стандартных типов (int/string/bool/...).

// Наивная попытка - так нельзя. Компилятор запрещает добавление новых методов к существующим базовым типам.
// err - cannot define new methods on non-local type int
//func (i *int) IsEven() bool {
//	if *i % 2 == 0 {
//		return true
//	}
//	return false
//}

// Правильный вариант - создать собственный тип и написать методы для него.
type SuperInt int

func (si *SuperInt) IsEven() bool {
	if *si%2 == 0 {
		return true
	}
	return false
}

func main() {
	num1 := SuperInt(202)
	fmt.Println(num1.IsEven())
	num2 := SuperInt(201)
	fmt.Println(num2.IsEven())

	// Приведение SuperInt => int
	var num3 int = int(num2)
	fmt.Println(num3)

}

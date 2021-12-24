package main

import (
	"fmt"
	"reflect"
)

// !!! Reflection - UNSAFE инструмент в Go + убивает читаемость.
// Используется для поточной кодогенерации - например низкоуровневое создание sql-запросов.

// Kind - к какому семейству принадлежит экземпляр типа (структура/интерфейс/функция/примитив).

type Order struct {
	orderId    int
	customerId int
}

func checkMyType(i interface{}) {
	// В runtime определим тип аргумента.
	typeOfSample := reflect.TypeOf(i)
	valueOfSample := reflect.ValueOf(i)
	kindOfSample := typeOfSample.Kind()
	fmt.Println("Type:", typeOfSample)
	fmt.Println("Value:", valueOfSample)
	fmt.Println("Kind:", kindOfSample)

	if reflect.ValueOf(i).Kind() == reflect.Struct {
		v := reflect.ValueOf(i)
		fmt.Println("Number of fields:", v.NumField())

		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field: %d, type: %T, value: %v\n", i, v.Field(i), v.Field(i))
		}
	}
}

func main() {
	ord := Order{23, 25}
	checkMyType(ord)
	fmt.Println()

	name := "Vasya"
	checkMyType(name)
	fmt.Println()

	checkMyType(func() {})
	fmt.Println()
}

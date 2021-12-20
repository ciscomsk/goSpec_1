package main

import (
	"fmt"
	"math"
)

// Константа - неизменяемая переменная, которые служат для:
// 1. Более строгого понимания кода.
// 2. Для того, чтобы случайно не поменять значение (предполагается, что значение константы неизменно).
// 3. Для удобных перобразований.

const MAIN_PORT = "8001"

func checkConstant() bool {
	// err - Unresolved reference 'ipAddress' - константа ipAddress вне области видимости функции.
	//if ipAddress == "127.0.0.1" {
	if MAIN_PORT == "8001" {
		return true
	}
	return false
}

func main() {
	// Объявление константы. Используются заглавные буквы + подчеркивания. Константа может быть только примитивным типом.
	const A = 10
	fmt.Println("A value:", A)

	// Константу нельзя поменять в ходе работы программы.
	// err - Cannot assign to A.
	//A = 25

	// Объявление блока констант с областью видимости - функция main.
	const (
		ipAddress = "127.0.0.1"
		port      = "8000"
		dbName    = "postgres"
	)
	fmt.Println("ipAddress value:", ipAddress)
	fmt.Println(checkConstant())
	fmt.Println()

	// !!! Значения констант ДОЛЖНЫ БЫТЬ ИЗВЕСТНЫ на момент компиляции.
	var sqrt float64 = math.Sqrt(25)
	fmt.Println("Var sqrt:", sqrt)
	fmt.Println()

	// Нельзя присвоить в константу что-либо, что является результатом вызова функции/метода.
	// err - Const initializer 'math.Sqrt(25)' is not a constant - на момент компиляции значение неизвестно.
	//const sqrt1 float64 = math.Sqrt(25)

	// Типизированные и нетипизированные константы.
	// untyped constant - компилятор сам подставляет тип в зависимости от переданного значения.
	const NUMERIC = 100
	var numInt = NUMERIC
	fmt.Printf("Value %v ant Type %T\n", numInt, numInt)

	// При использовании нетипизированных констант мы разрешаем компилятору использовать
	// неявное приведение типов в момент присваивания значений констант в переменные.
	var numInt8 int8 = NUMERIC
	var numInt32 int32 = NUMERIC
	var numInt64 int64 = NUMERIC
	var numFloat32 float32 = NUMERIC
	var numComplex complex64 = NUMERIC
	fmt.Printf("numInt8 value: %v type: %T\n", numInt8, numInt8)
	fmt.Printf("numInt32 value: %v type: %T\n", numInt32, numInt32)
	fmt.Printf("numInt64 value: %v type: %T\n", numInt64, numInt64)
	fmt.Printf("numFloat32 value: %v type: %T\n", numFloat32, numFloat32)
	fmt.Printf("numComplex value: %v type: %T\n", numComplex, numComplex)
	fmt.Println()

	// https://stackoverflow.com/questions/54605650/arithmetic-on-uint8-int8
	// 100 + 100 == 200, но max Int8 == 127
	//
	// 01100100
	// +
	// 01100100
	// ________
	// 11001000	== -128 (10000000) + 72 (1001000) == -56
	fmt.Printf("%v + %v is %v\n", numInt8, NUMERIC, numInt8+NUMERIC)

	// typed const. Указание типа - повышение читабельности кода.
	const ADMIN_EMAIL string = "admin@admin.com"

	// Константы в go хранятся во внутреннем стеке (помещаются туда а момент компиляции) RUNTIME программы и не очищаются до ее окончания.

}

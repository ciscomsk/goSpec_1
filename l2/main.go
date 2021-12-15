package main

import (
	"fmt"
	"math"
)

// go run l2/main.go
func main() {
	// Вывод в консоль.
	fmt.Println("Hello world", "Hello another")
	fmt.Println("Second line")

	// Вывод в консоль без переноса строки.
	fmt.Print("First")
	fmt.Print("Second")
	fmt.Print("Third")
	fmt.Println()
	fmt.Println()

	// Формитированный вывод - стандартный вывод os.Stdout c флагами форматирования.
	fmt.Printf("Hello, my name is %s\nMy age is %d\n", "Bob", 42)

	// Декларирование переменных - процесс связывания имени переменной с типом значения.
	// Инициализация - присвоение значения.

	// При объявлении переменная инициализируется "нулевым" значением. В go не бывает неинициализированных переменных.
	var age int // == 0.  var age int = 0 - плохая практика
	fmt.Println("My age is:", age)
	age = 32
	// err - (type untyped float) cannot be represented by the type int
	//age = 32.4
	fmt.Println("Age after assignment:", age)

	// Декларирование и инициализация пользовальским значением.
	var height int = 183
	fmt.Println("My height is:", height)

	// Полустрогая типизация. Рекомендуется опускать типы при работе с примитивами.
	var uid = 12345
	fmt.Println("My uid:", uid)
	// %T - получение типа переменной.
	fmt.Printf("%T\n", uid)
	fmt.Println()

	// Декларирование и инициализация переменных ОДНОГО типа (множественный случай). Тип можно опустить.
	var firstVar, secondVar int = 20, 30
	// err - переменные разных типов.
	//var firstVar float32, secondVar int = 20, 30
	fmt.Printf("Firstvar:%d, Secondvar:%d\n", firstVar, secondVar)

	// Декларирование блока переменных. Типы можно опускать
	var (
		personName string = "Bob"
		personAge  int    = 42
		personUID  int
	)
	fmt.Printf("Name: %s\nAge %d\nUID: %d\n", personName, personAge, personUID)
	fmt.Println()

	// Если не указывать типы - можно задекларировать несколько переменных с разными типами.
	// Плохая практика - лучше объявлять разные типы в блоках.
	var a, b = 30, "Vova"
	fmt.Println(a, b)
	fmt.Println()

	// Повторное декларирование переменной - приводит к ошибке компиляции.
	// err - redeclared in this block
	//var a = 200

	// Короткая декларация (объявление) - :=, значение нужно задать. var == :=
	count := 10
	fmt.Println("Count:", count)
	count++ // == count= count + 1
	fmt.Println("Count:", count)

	// Множественное присваивание через :=
	aArg, bArg := 10, "Vova"
	// err - Assignment count mismatch: 2 = 1
	//aArg, bArg := 10
	fmt.Println(aArg, bArg)

	// Переприсвоение значений.
	intA, intB := 10, 20
	// err - No new variables on the left side of ':='
	//intA, intB := 30, 40
	intA, intB = 30, 40
	fmt.Println(intA, intB)

	// Если есть хотя бы 1 новая переменная - можно использовать :=
	intB, intC := 50, 60
	fmt.Println(intA, intB, intC)
	fmt.Println()

	width, length := 20.5, 30.2
	// %.3f - 3 знака после запятой.
	fmt.Printf("Min dimensional of rectangle is: %.3f\n", math.Min(width, length))

	word1 := "имя"
	word2 := "твое"
	word3 := "мне"
	word4 := "знакомо"

	fmt.Println(word4, word3, word2, word1)
	fmt.Println(word3, word1, word2, word4)
	fmt.Println(word2, word4, word1, word3)
}

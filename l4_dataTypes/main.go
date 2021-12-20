package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

// go run l4/main.go
func main() {
	// Boolean.
	var firstBoolean bool // default => false
	fmt.Println(firstBoolean)

	aBoolean, bBoolean := true, false
	fmt.Println("AND:", aBoolean && bBoolean) // AND - логическое умножение
	fmt.Println("OR:", aBoolean || bBoolean)  // OR - логическое сложение
	fmt.Println("NOT:", !aBoolean)            // NOT - логическое отрицание

	// err - Invalid operation: the operator + is not defined on bool
	//fmt.Println(aBoolean + bBoolean)
	fmt.Println()

	// Numerics. Integers.
	// int8, int16, int32, int64, int. int - платформозависимый тип (32/64).
	// uint8, uint16, uint, 32, uint64, uint. - Беззнаковые целые.
	var a int = 32
	b := 32 // при коротком объявлении тип == int
	fmt.Println("Value of a:", a, "Value of b:", b, "a + b:", a+b)

	// Вывод типа через %T форматирование.
	fmt.Printf("Type is %T\n", a)

	// Сколько байт занимает переменная типа int.
	fmt.Printf("Type %T size of %d bytes\n", a, unsafe.Sizeof(a)) // 8
	// Короткое объявление => тип == int.
	fmt.Printf("Type %T size of %d bytes\n", b, unsafe.Sizeof(b)) // 8

	var first32 int32 = 12
	var second64 int64 = 13
	fmt.Println(first32, second64)
	// err - Invalid operation: mismatched types int32 and int64
	//fmt.Println(first32 + second64)
	// => !!! явное приведение типа к меньшему - возможна потеря данных (не влезет).
	fmt.Println(first32 + int32(second64))
	// => !!! явное приведение типа к большему - потери данных не будет.
	fmt.Println(int64(first32) + second64)

	// При выполнении арифметических операций над int и intX - также обязательно использовать механизм приведения. Т.к. int != int64
	var third64 int64 = 16123414
	var fourthInt int = 156234
	fmt.Println(third64, fourthInt)
	// err - Invalid operation: mismatched types int64 and int. Несмотря на то, что оба типа занимают по 8 байт.
	//fmt.Println(third64 + fourthInt)
	fmt.Println(third64 + int64(fourthInt))

	// Арифметические операции: + - * / %

	// Аналогичным образом устроены uint типы.
	fmt.Println()

	// Numerics. Float.
	// float32, float64
	floatFirst, floatSecond := 5.67, 12.54 // float64 - платформозависимый.
	fmt.Printf("Type of floatFirst %T, type of floatSecond %T\n", floatFirst, floatSecond)

	sum := floatFirst + floatSecond
	sub := floatFirst - floatSecond
	fmt.Println("Sum:", sum, "Sub:", sub)
	// Вывод вещественного числа с указанной точностью.
	fmt.Printf("Sum: %.3f, Sub: %.3f\n", sum, sub)
	fmt.Println()

	// Numerics. Complex.
	c1 := complex(5, 7)
	c2 := 12 + 32i
	fmt.Println(c1 + c2)
	fmt.Println(c1 * c2)
	fmt.Println()

	// !!! String. Строка в go - это набор байт.
	name := "Mikhail"
	lastname := "Kuznetsov"
	concat := name + " " + lastname
	fmt.Println("Full name:", concat)

	// len возвращает количество элементов в наборе.
	fmt.Println("Length of string:", name, len(name))         // 7 байт
	fmt.Println("Length of string:", "Михаил", len("Михаил")) // 12 байт

	// Для хранения utf символов используются руны (аналог char). 1 руна == 1 utf символ.
	fmt.Println("Amount of chars:", name, utf8.RuneCountInString(name))         // 7
	fmt.Println("Amount of chars:", "Михаил", utf8.RuneCountInString("Михаил")) // 6

	// rune -> alias int32
	var sampleRune rune
	fmt.Println(sampleRune)
	var anotherRune rune = 'Q' // инициализация руны символьным значением
	fmt.Printf("Rune as char %c\n", anotherRune)
	var thirdRune rune = 234542 // инициализация руны int32
	fmt.Printf("Rune as char %c\n", thirdRune)

	// Поиск подстроки в строке.
	totalString, subString := "ABCDEDFG", "CDE"
	fmt.Println(strings.Contains(totalString, subString))

	// Сравнение строк.
	// v1 - новый способ.
	fmt.Println("abcd" > "def")
	// v2 - старый способ.
	// -1 - левая строка меньше правой, 0 - элементы равны, 1 - левая строка больше правой.
	fmt.Println(strings.Compare("abcd", "def"))
	fmt.Println()

	// Byte -> alias uint8
	var aByte byte
	fmt.Println("Byte:", aByte)
}

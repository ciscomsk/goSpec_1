package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Строка - слайс байтов.
	name := "Hello world"
	fmt.Println(name)
	fmt.Println()

	word := "Sample word"
	fmt.Printf("String %s\n", word)

	// Получение значений байтов в слайсе.
	fmt.Print("Bytes: ")
	for i := 0; i < len(word); i++ {
		// %x - формат представления 16-ричного байта.
		fmt.Printf("%x ", word[i])
	} // !!! 53 61 6d 70 6c 65 20 77 6f 72 64 - именно эти значения считает функция len
	fmt.Println()

	// Обращение к отдельным символам.
	// !!! Строки в go хранятся как наборы UTF8-символов. Символ может занимать больше 1 байта (пример - кириллица - 2 байта).
	fmt.Print("Characters: ")
	for i := 0; i < len(word); i++ {
		// %c - формат представления char.
		fmt.Printf("%c ", word[i])
	} // S a m p l e   w o r d
	fmt.Println()
	fmt.Println()

	wordС := "Тестовая строка"
	fmt.Printf("String %s\n", wordС)

	fmt.Print("Bytes: ")
	for i := 0; i < len(wordС); i++ {
		fmt.Printf("%x ", wordС[i])
	} // d0 a2 d0 b5 d1 81 d1 82 d0 be d0 b2 d0 b0 d1 8f 20 d1 81 d1 82 d1 80 d0 be d0 ba d0 b0
	fmt.Println()

	fmt.Print("Characters: ")
	for i := 0; i < len(wordС); i++ {
		fmt.Printf("%c ", wordС[i])
	} // Ð ¢ Ð µ Ñ Ñ Ð ¾ Ð ² Ð ° Ñ   Ñ  Ñ  Ñ  Ð ¾ Ð º Ð °
	fmt.Println()
	fmt.Println()

	// Руна - стандартный тип в Go (alias -> int32), позволяющий хранить неделимый UTF символ
	// вне зависимости от того, сколько байт он занимает.

	// Преобразование слайса байт к слайсу рун. Наборот - []byte(sliceRune)
	runeSlice := []rune(wordС)

	for i := 0; i < len(runeSlice); i++ {
		fmt.Printf("%c ", runeSlice[i])
	}
	fmt.Println()
	fmt.Println()

	// Итерирование по строке с использованием рун.
	for id, runeVal := range wordС {
		fmt.Printf("%c starts at position %d\n", runeVal, id) // !!! id - индекс байта, с которого начинается руна runeVal
	}
	fmt.Println()

	// Создание строки из слайса байт.
	myByteSlice := []byte{0x40, 0x41, 0x42, 0x43}
	myStr := string(myByteSlice)
	fmt.Println(myStr)

	myDecimalByteSlice := []byte{100, 101, 102, 103} // синтаксический сахар - использование десятичного представления байтов. byte -> uint8
	myDecimalStr := string(myDecimalByteSlice)
	fmt.Println(myDecimalStr)
	fmt.Println()

	// Создание строки из слайса рун.
	// Руны как hex.
	runeHexSlice := []rune{0x45, 0x46, 0x47, 0x48} // также можно передавать десятичные значения.
	myStrFromRunes := string(runeHexSlice)
	fmt.Println("From Runes(hex):", myStrFromRunes)

	// Руны как литералы.
	// '' - обозначение руны.
	runeLiteralSlice := []rune{'M', 'i', 'k', 'e'}
	myStrFromRuneLiterals := string(runeLiteralSlice)
	fmt.Println("From Runes(literals):", myStrFromRuneLiterals)
	fmt.Printf("%s and %T\n", myStrFromRuneLiterals, myStrFromRuneLiterals)
	fmt.Println()

	// Длина и емкость строки.
	// Длина len() - количество байт в слайсе.
	fmt.Println("Length of Миша:", len("Миша"), "bytes")
	// Длина utf8.RuneCountInString() - количество рун.
	fmt.Println("Length of Миша:", utf8.RuneCountInString("Миша"), "runes")
	fmt.Println()

	// !!! Строки в go - иммутабельны => вычисление емкости бессмысленно. Append строки невозможен.
	// err - Invalid argument for the cap function
	//fmt.Println(cap("Миша"))
	// cap для строки всегда равен len (т.к. строка - неизменяема)
	fmt.Println(cap([]rune("Миша")))
	fmt.Println()

	// Сравнение строк == и != - нормально работает с go 1.6.
	word1, word2 := "Вася", "Петя"
	if word1 == word2 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}
	fmt.Println()

	// Конкатенация - значение исходных строк не изменяется, создается новая.
	word3 := word1 + word
	fmt.Println(word3)
	fmt.Println()

	// StringBuilder
	firstName := "Alex"
	secondName := "Johnson"
	fullName := fmt.Sprintf("%s #### %s", firstName, secondName)
	fmt.Println(fullName)
	fmt.Println()

	// !!! Строки - иммутабельны.
	// err - Cannot assign to fullName[0]
	//fullName[0] = byte(0x41)

	// !!! Слайсы изменяемы.
	fullNameSlice := []rune(fullName)
	fullNameSlice[0] = 'F'           // изменили слайс
	fullName = string(fullNameSlice) // переприсвоили переменную
	fmt.Println("String mutation:", fullName)
	fmt.Println()

	// Сравнение рун.
	if 'Q' == 'q' {
		fmt.Println("Runes equal")
	} else {
		fmt.Println("Runes not equal")
	}

	// !!! Методы работы со строками - import "strings"
}

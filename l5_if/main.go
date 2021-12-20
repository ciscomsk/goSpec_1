package main

import (
	"fmt"
	"strings"
)

// go run l5/main.go
func main() {
	// Условный оператор.
	var condition1, condition2 bool
	if condition1 {
		// body
	} else if condition2 {
		// body
	} else {
		// body
	}

	var value int
	fmt.Scan(&value)

	if value%2 == 0 {
		fmt.Println("The number", value, "is even")
	} else {
		fmt.Println("The number", value, "is odd")
	}

	var color string
	fmt.Scan(&color)

	if strings.Compare(color, "green") == 0 {
		fmt.Println("Color is green")
	} else if strings.Compare(color, "red") == 0 {
		fmt.Println("Color is red")
	} else {
		fmt.Println("Unknown color")
	}

	fmt.Println()

	// Инициализация в блоке условного оператора. Инициализируемая переменная будет доступна только в области видимости условного оператора (в телах if/else if/else).
	// !!! Блок присваивания - ТОЛЬКО :=
	if num := 10; num%2 == 0 {
		fmt.Println("EVEN")
	} else {
		fmt.Println("ODD")
		fmt.Println(num)
	}

	// err - Unresolved reference 'num' - переменная недоступна вне условного оператора.
	//fmt.Println(num)

	//if data, err := someFunc(); err != nil {
	// data будет существовать только здесь.
	//}

	// Переносить блоки в условном операторе - нельзя.
	var age int = 10

	// OK
	if age > 7 {
		fmt.Println("Go to school")
	} else {
		fmt.Println("Another case")
	}

	// NOT OK
	/*
		if age > 7 {
			fmt.Println("Go to school")
		} // фактически == } ; - т.е. следующий блок else не имеет родительского блока if.
		// err - 'else' unexpected
		else {
			fmt.Println("Another case")
		}
	*/

	// Не идеоматично.
	if width := 100; width > 100 {
		fmt.Println("Width > 100")
	} else {
		fmt.Println("Width <= 100")
	}

	// Идеоматично.
	// !!! В Go стараются избегать блоков else.
	if height := 100; height > 100 {
		fmt.Println("Height > 100")
		return
	}
	fmt.Println("Height <= 100")
}

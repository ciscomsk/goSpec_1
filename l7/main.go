package main

import "fmt"

func main() {
	// Switch.
	var price int
	fmt.Scan(&price)

	if price == 100 {
		fmt.Println("First case")
	} else if price == 110 {
		fmt.Println("Second case")
	} else if price == 120 {
		fmt.Println("Third case")
	} // ...
	// =>
	switch price {
	case 100:
		fmt.Println("First case")
	case 110:
		fmt.Println("Second case")
	case 120:
		fmt.Println("Third case")
	// err - Duplicate case 120 - дублирующиеся условия запрещены.
	//case 120:
	//	fmt.Println("Third case")
	case 130:
		fmt.Println("Another case")
	default:
		// Отработает, если не один из вышеперечисленных кейсов не подходит.
		fmt.Println("Default case")
	}

	// case с множеством вариантов.
	var ageGroup string = "e" // возможные варианты - a/b/c/d/e

	switch ageGroup {
	case "a", "b", "c":
		fmt.Println("AgeGroup 10-40.")
	case "d", "e":
		fmt.Println("AgeGroup 50-70.")
	default:
		fmt.Println("You are too young/old.")
	}

	// case с выражениями.
	var age int
	fmt.Scan(&age)

	switch {
	case age <= 18:
		fmt.Println("Too young.")
	case age > 18 && age <= 30:
		fmt.Println("Second case.")
	case age > 30 && age <= 100:
		fmt.Println("Too old.")
	default:
		fmt.Println("Who are you?")
	}

	// case с "проваливаниями" - fallthrough.
	var number1 int
	fmt.Scan(&number1)

	switch {
	case number1 < 100:
		fmt.Printf("%d is less then 100\n", number1)
		fallthrough
	case number1 < 200:
		fmt.Printf("%d less then 200\n", number1)
	}
	fmt.Println()

	// !!! "Проваливания" выполняют даже ЛОЖНЫЕ кейсы - УСЛОВИЯ СЛЕДУЮЩЕГО КЕЙСА НЕ ПРОВЕРЯЮТСЯ, А СРАЗУ ВЫПОЛНЯЮТСЯ.
	var number2 int
	fmt.Scan(&number2)

	switch {
	case number2 < 100:
		fmt.Printf("%d is less then 100\n", number2)

		if number2%2 == 0 {
			// Так "проваливания" не будет.
			break
		}

		fallthrough
	case number2 > 200:
		fmt.Printf("%d GREATER then 200\n", number2)
		// Если бы fallthrough не было - здесь выполнение switch было закончено
		fallthrough
	default:
		fmt.Printf("%d DEFAULT", number2)
		// err - Cannot use 'fallthrough' in the final case of the 'switch' statement. Из дефолтного кейса "провалиться" нельзя.
		//fallthrough
	}
	fmt.Println()

	// switch можно лейблить.
outer:
	switch {
	default:
		break outer
	}

	// Терминация цикла for из switch
	var i int

uberloop:
	for {
		fmt.Scan(&i)

		switch {
		case i%2 == 0:
			fmt.Printf("Value is %d and it's even\n", i)
			break uberloop

		}
	}

	fmt.Println("END")

}

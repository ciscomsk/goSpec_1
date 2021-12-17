package main

import "fmt"

// Явная функция (явное определение) - локально-определенный блок кода, имеющий имя.
// Функцию необходимо определить + функцию необходимо вызвать.

// Сигнатура функций и их определение.
//func functionName(params type) returnValueType {
// body
//}

// Параметры и тип выходного результата - опциональны.
// !!! Если функция ничего не возвращает - тип выходного значения не указывается.
func functionWO() {
}

func add(a int, b int) int {
	result := a + b
	return result
}

// Функции можно определять как до момента ее вызова в функции main, так и в любом месте пакета .

// Функция с однотипными параметрами.
func mult(a, b, c, d int) int {
	result := a * b * c * d
	return result
}

// Возврат больше чем 1 значения => (returnType1, returnType2, ...)
func rectangleParameters(length, width float64) (float64, float64) {
	var perimeter float64 = 2 * (length + width)
	var area float64 = length * width

	return perimeter, area
}

// Именованный возврат значений - редко используется.
func namedReturn(a, b int) (perimeter int32, area int) {
	// err - No new variables on the left side of ':=' - компилятор уже инициализировал переменную, т.к. мы указали ее в результатах.
	//perimeter := 2 * (a + b)

	perimeter = int32(2 * (a + b))
	area = a * b
	return // не указываем возвращаемые переменные.
}

// При вызове оператора return функция прекращает свое выполнение и возвращает результат.
func funcWithReturn(a, b int) int {
	if a > b {
		return a - b // При a == 10, b == 7 - функция закончит свое выполнение на этой строке.
	}

	if a == b {
		return a
	}

	return b // При a == 7, b == 10
}

func funcWithReturn2(a, b int) (int, bool) {
	if a > b {
		return a - b, true
	}

	if a == b {
		return a, true
	}

	return b, false
}

// Если return опущен или пустой - функция не будет возвращать ничего.
func emptyReturn(a int) {
	fmt.Println("I'm emptyReturn with parameter:", a)
	return // если опустить - результат аналогичен
}

// Variadic parameters (континуальные параметры). От 0 до множества.
func helloVariadic(a ...int) {
	fmt.Println(a)
	fmt.Printf("value %v and type %T\n", a, a)
}

// Смешение параметров с variadic.
// 1. Континуальный параметр указывается в конце.
// 2. Может быть всего один variadic параметр на функцию.
func someStrings(a, b int, words ...string) {
	fmt.Println("Parameter:", a)
	fmt.Println("Parameter:", b)

	var result string
	for _, word := range words {
		result += word
	}
	fmt.Println("Concat result:", result)

}

// Передача слайса или использование variadic parameters?
// !!! Использование вариадического параметра - более производительный и удобный вариант.
func sumVariadic(nums ...int) int {
	var sum int
	for _, val := range nums {
		sum += val
	}

	return sum
}

func sumSlice(nums []int) int {
	var sum int
	for _, val := range nums {
		sum += val
	}

	return sum
}

// Анонимная функция внутри явной.
func bigFunction(aArg, bArg int) int {
	// func(a, b int) int { return a + b } - определение анонимной функции.
	// (aArg, bArg) - передача параметров в анонимную функцию.
	return func(a, b int) int { return a + b + 1 }(aArg, bArg)
}

func main() {
	// Вызов функции.
	// В момент вызова проверяется количество и типы аргументов - управление передается в заданный блок кода.
	// return возвращает управление вызывающей стороне.
	res := add(10, 20)
	fmt.Println("Result of add(10, 20):", res)

	fmt.Println("Result of mult(1, 2, 3, 4):", mult(1, 2, 3, 4))
	println()

	per, area := rectangleParameters(10.5, 10.5)
	fmt.Println("Area of rect(10.5, 10.5):", area)
	fmt.Println("Perimeter of rect(10.5, 10.5):", per)

	// Пропуск полученных значений.
	newPer, _ := rectangleParameters(10, 10)
	fmt.Println("NewPer:", newPer)
	fmt.Println()

	namedPer, namedArea := namedReturn(10, 20)
	fmt.Println(namedPer, namedArea)
	fmt.Println()

	emptyReturn(10) // ок
	// err - emptyReturn(10) used as a value, but it returns nothing - emptyReturn ничего не возвращает.
	//resEmpty := emptyReturn(10)
	fmt.Println()

	helloVariadic(10, 20, 30, 40, 50, 60)
	helloVariadic(10)
	helloVariadic()
	fmt.Println()

	someStrings(2, 10, "Bob", "Alex", "Vito")
	fmt.Println()

	sum1 := sumVariadic(10, 20, 30)
	sliceNums := []int{10, 20, 30}
	sum2 := sumVariadic(sliceNums...) // распаковка слайса
	fmt.Println(sum1, sum2)

	fmt.Println(sumSlice([]int{10, 20, 30})) // дополнительное выделение памяти при создании массива + дополнительная нагрузка на gc.
	fmt.Println(sumSlice(sliceNums))
	fmt.Println()

	// Анонимная функция.
	// Используется для организации горутин/кастомных ключей сортировки/небольшие промежуточные вычисления.
	anon := func(a, b int) int {
		return a + b
	}
	fmt.Println("Anon", anon(20, 30))
	fmt.Println()

	fmt.Println("BigFunction(10, 20):", bigFunction(10, 20))

}

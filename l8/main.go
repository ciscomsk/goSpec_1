package main

import "fmt"

func main() {
	// Массив - структура данных, располагающаяся на непрерывном участке памяти. Обращение к элементу массива - по индексу.
	// Массив - неизменяемая структура, хранит элементы одного типа.

	// Декларация массива.
	// Массив для 5 int элементов. При инициализации массива нужно указать количество элементов в нем.
	var arr [5]int
	fmt.Println("This is my array:", arr)

	// Определение элементов массива (после предварительной инициализации).
	arr[0] = 10
	arr[1] = 20
	arr[3] = -500
	arr[4] = 720
	fmt.Println("After elements init:", arr)

	// err - Invalid array index '5' (out of bounds for the 5-element array)
	//arr[5] = 100

	// Определение массива с указанием элементов.
	// Если при нициализации количество заданных элементов меньше длины массива - недостающие элементы будут инициализированы нулевым значением.
	newArr := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Short declaration and init:", newArr)
	fmt.Println()

	// Создание массива через инициализацию переменных.
	// [...]int{} - пустой массив.
	arrWithValues := [...]int{10, 20, 30, 40}
	fmt.Println("Arr declaration with [...]:", arrWithValues, "length:", len(arrWithValues))
	arrWithValues[0] = 100000
	fmt.Println("Arr declaration with [...]:", arrWithValues, "length:", len(arrWithValues))
	fmt.Println()

	// Размер массива нельзя изменить. Можно сложить 1-й и 2-й массив в 3-й.
	// В массив нельзя положить элемент с другим типом.

	// !!! Массив == набор ЗНАЧЕНИЙ. Т.е. при всех манипуляциях - массив копируется (жестко - на уровне компилятора).
	first := [...]int{1, 2, 3}
	second := first // !!! second - отдельный массив, а не ссылка на first
	second[0] = 10000
	fmt.Println("First arr:", first)
	fmt.Println("Second arr:", second)
	fmt.Println()

	// !!! Массив и его размер - составляющие типа (размер массива - часть типа).
	var aArr [5]int
	var bArr [6]int
	aArr[0] = 100
	bArr[1] = 200
	// err - Cannot use 'aArr' (type [5]int) as the type [6]int - массивы с разными размерами - разные типы.
	//bArr = aArr

	// Итерирование по массиву - цикл for.
	floatArr := [...]float64{12.5, 13.5, 15.2, 10.0, 12.0}

	for i := 0; i < len(floatArr); i++ {
		fmt.Printf("%d element of arr is %.2f\n", i, floatArr[i])
	}
	fmt.Println()

	// Итерирование по массиву - цикл range based for - возвращает индекс + значение.
	var sum float64
	for id, val := range floatArr {
		fmt.Printf("%d element of arr is %.2f\n", id, val)
		sum += val
	}
	fmt.Println("Total sum is:", sum)
	fmt.Println()

	// Игнорирование id/val в range.
	//for id, _ := range floatArr {	// == for id := range floatArr { - 2-ю переменную можно опустить - не рекомендуется, т.к. ухудшается читаемость.
	for _, val := range floatArr {
		fmt.Printf("%.2f value WO id\n", val)
	}
	fmt.Println()

	// Многомерные массивы - должны содержать массивы одного типа.
	words := [2][2]string{
		{"Bob", "Alice"},
		{"Victor"}, // можно указывать не все элементы
	}
	fmt.Println("Multidimensional array:", words)
	fmt.Println()

	// Итерирование по многомерному массиву.
	for _, val1 := range words {
		for _, val2 := range val1 {
			fmt.Printf("%s ", val2)
		}
		fmt.Println()
	}
	fmt.Println()

	// Общение со срезом. Срез ~ массив, без указания начального размера.
	slice := []int{10, 20, 30}
	slice[0] = slice[0] * 10
	slice[1] = 200
	// Добавление нового элемента.
	slice = append(slice, 200)

	for id, val := range slice {
		fmt.Println(id, val)
	}
	fmt.Println()

	var emptySlice []int
	emptySlice = append(emptySlice, 200)
	fmt.Println(emptySlice)
}

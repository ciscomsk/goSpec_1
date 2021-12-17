package main

import "fmt"

func main() {
	// Слайсы (срезы) - динамическая обертка над массивом (интерфейс над массивом)
	startArr := [4]int{10, 20, 30, 40}

	// !!! Инициализация: Array - [number]type, Slice - []type. Слайс инициализируется пустыми квадратными скобками.

	// Создание слайса на основе массива.
	// [:] - весь массив
	// [0:2] - от 0 до 2 не включая 2
	var startSlice []int = startArr[0:2]
	fmt.Println("Slice[0:2]:", startSlice)

	// Создание слайса без явной инициализации массива.
	secondSlice := []int{15, 20, 30, 40}
	fmt.Println(secondSlice)
	fmt.Println()

	// Изменение элементов среза.
	// !!! При изменении элементов слайса - меняются и элементы нижележащего массива.
	// !!! Срез - набор ссылок на элементы нижележащего массива.
	originArr := [...]int{30, 40, 50, 60, 70, 80}
	firstSlice := originArr[1:4]

	for i, _ := range firstSlice {
		firstSlice[i]++
	}
	fmt.Println("OriginArr:", originArr)
	fmt.Println("FirstSlice:", firstSlice)
	fmt.Println()

	// Один массив и два производных среза.
	// !!! Изменения во всех срезах меняют нижележащий массив.
	oArr := [...]int{30, 40, 50, 60, 70, 80}
	fSlice := oArr[:]
	sSlice := oArr[2:5]

	fmt.Println("Before modifications: Arr:", oArr, "fSlice:", fSlice, "sSlice", sSlice)
	fSlice[3]++
	sSlice[1]++
	fmt.Println("After modifications: Arr:", oArr, "fSlice:", fSlice, "sSlice", sSlice)
	fmt.Println()

	// Срез как встроенный тип.
	//type slice struct {
	//	Length int
	//	Capacity int
	//	ZeroElement *byte - указатель на первый элемент
	//}

	// Длина и емкость слайса.
	// Длина - количество фактических элементов.
	// Емкость (capacity) - значение, показывающее сколько элементов можно разместить в слайсе без выделения дополнительной памяти под нижележащий массив.
	wordsSlice := []string{"one", "two", "three"}
	fmt.Println("WordSlice:", wordsSlice, "Length", len(wordsSlice), "Capacity:", cap(wordsSlice))
	wordsSlice = append(wordsSlice, "four")
	fmt.Println("WordSlice:", wordsSlice, "Length", len(wordsSlice), "Capacity:", cap(wordsSlice))
	fmt.Println()

	// Пример - инициализировали срез с 3 элементами без явного указания массива.
	// => под капотом
	// 1. Компилятор создал массив на 3 элемента + вернул его адрес. Срез запомнил адрес и теперь ссылается на него.
	// 2. Деформируем слайс - увеличим количество элементов (> 3).
	// Проблема - в нижележащем массиве только 3 ячейки.
	// 3. Компилятор ищет в памяти место для массива 3*2 (в общем случае n*2, n - исходный размер массива)
	// 4. В найденное место копируются старые элементы + добавляемый элемент.
	// 5. Компилятор возвращает слайсу новый адрес в памяти, где находится новый массив.

	numerics := []int{1, 2}

	for i := 0; i < 200; i++ {
		if i%5 == 0 {
			fmt.Println("Current len:", len(numerics), "Current cap:", cap(numerics))
		}
		numerics = append(numerics, i)
	}
	fmt.Println()

	// !!! После выделения памяти под новый массив, ссылки со старым будут перенесены в новый.
	numArr := [2]int{1, 2}
	numSlice := numArr[:]

	numSlice = append(numSlice, 3) // !!! в этот момент numSlice больше не ссылается на numArr
	numSlice[0] = 10
	fmt.Println(numArr)
	fmt.Println(numSlice)
	fmt.Println()

	// Эффективный способ создания слайсов.
	// make() - функция, позволяющая более детально создавать срезы.

	// err - len larger than cap in make([]int) - такое невозможно.
	//sl := make([]int, 10, 9)
	sl := make([]int, 10, 15)
	// => под капотом
	// 1. инициализируется нижележащий массив arr = [15]int
	// 2. По массиву делается срез arr[0:10] и возвращается
	fmt.Println(sl)
	println()

	// Добавление элементов в срез.
	myWords := []string{"one", "two", "three"}
	fmt.Println("myWords:", myWords)
	myWords = append(myWords, "four")
	fmt.Println("myWords:", myWords)

	// Добавление слайса в слайс.
	anotherSlice := []string{"five", "six", "seven"}

	// v1 - range
	for _, val := range anotherSlice {
		myWords = append(myWords, val)
	}
	fmt.Println(myWords)

	// v2 - распаковка слайса.
	myWords = append(myWords, anotherSlice...)
	// ==
	//myWords = append(myWords, "five", "six", "seven")
	fmt.Println(myWords)
	fmt.Println()

	// Многомерный срез.
	mSlice := [][]int{
		{1, 2},
		{3, 4, 5, 6},
		{10, 20, 30},
		{},
	}
	fmt.Println(mSlice)

}

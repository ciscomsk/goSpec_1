package main

import "fmt"

// Указатели на массивы (и на мапы) - почему так делать не надо.

func mutation(arr *[3]int) { // ограничение на тип аргумента - строго 3 элемента
	//(*arr)[1] = 909
	//(*arr)[2] = 100000
	// ==> если функция принимает указатель то можно записать короче - go сам разыменует указатель на массив
	arr[1] = 909
	arr[2] = 100000
}

// Более универсально и идеоматично использовать слайс.
func mutationSlice(slc []int) { // ограничений на тип аргумента - нет
	slc[1] = 909
	slc[2] = 100000
}

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println("Arr before mutation:", arr)
	mutation(&arr)
	fmt.Println("Arr after mutation", arr)
	fmt.Println()

	newArr := [3]int{1, 2, 3}
	fmt.Println("NewArr before mutationSlice:", newArr)
	mutationSlice(newArr[:])
	fmt.Println("NewArr after mutationSlice:", newArr)
}

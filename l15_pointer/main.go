package main

import "fmt"

// Указатель - переменная, хранящая в качестве значения адрес в памяти другой переменной.
// Главный плюс - отсутствие затрат на копирование при изменении значения.

// Определение функции, принимающей параметр как указатель.
func changeParam(val *int) {
	*val += 100
}

func changeWOPointer(val int) {
	// Изменение val происходят в локальном скоупе функции и не видно ивне.
	val += 100
}

// Инициализация функции возвращающей указатель.
func returnPointer() *int {
	var numeric int = 321
	// Возвращает указатель на локально определенную переменную (в c++ результат такой операции - неопределен).
	return &numeric // в момент возврата go перемещает данную переменную в кучу.

	// В стеке лежат неизменяемые (статические данные) - например числа/строки/массивы. Поиск в стеке выполняется гораздо быстрее, чем поиск в куче.
	// В куче лежат изменяемые (динамические) данные - например слайс.
}

func main() {
	// Определение указателя.
	var variable int = 30
	var pointer *int = &variable // * - указатель, & - операция взятия адреса в памяти.
	fmt.Println(pointer)         // 0xc0000b8000 - место в памяти, где хранится int значение 30.
	fmt.Printf("Type of pointer: %T\n", pointer)
	fmt.Println()

	// Нулевое значение указателя == nil - указатель в никуда.
	var zeroPointer *int
	fmt.Printf("Value: %v and type: %T\n", zeroPointer, zeroPointer)

	if zeroPointer == nil {
		zeroPointer = &variable
		fmt.Printf("After initialization value: %v and type: %T\n", zeroPointer, zeroPointer)
	}
	fmt.Println()

	// Указатель на указатель.
	zeroPointerToPointer := &zeroPointer
	fmt.Println(zeroPointerToPointer)
	fmt.Println()

	// Разыменование указателя == получение значения.
	// & - применяется к значениям => получаем указатель.
	// * - применяется к указателям => получаем значение.
	var numericValue int = 32
	var pointerToNumeric *int = &numericValue
	fmt.Printf("Value in numericValue is: %v\nAddress is: %v\n", *pointerToNumeric, pointerToNumeric)
	fmt.Println()

	// Создание указателей на "нулевые" значения типов.
	// v1
	var zeroVar int
	var zeroPoint *int = &zeroVar
	fmt.Printf("Value in *zeroPointer: %v\nAddress is: %v\n", *zeroPoint, zeroPoint)

	// v2 new - более удобный вариант.
	// new - под капотом создает zeroValue для int и возвращает адрес, где хранится это zeroValue.
	var zeroPoint2 = new(int)
	fmt.Printf("Value in *zeroPointer2: %v\nAddress is: %v\n", *zeroPoint2, zeroPoint2)
	fmt.Println()

	// Изменение значения хранимого по адресу через указатель.
	zeroPointerToInt := new(int)
	fmt.Printf("Address is: %v and value: in zeroPointerToInt is: %v\n", zeroPointerToInt, *zeroPointerToInt)
	*zeroPointerToInt += 40
	fmt.Printf("Address is: %v and new value: in zeroPointerToInt is: %v\n", zeroPointerToInt, *zeroPointerToInt)
	fmt.Println()

	b := 345
	fmt.Println(b)
	a := &b
	*a++
	fmt.Println(b)
	c := &b
	*c += 100
	fmt.Println(b)
	fmt.Println()

	// Указательная арифметика - из адреса одной ячейки можно продвинуться в другие ячейки.
	// В go нет указательной арифметики (как в c++).
	cppLike := "cpp"
	cppPointer := &cppLike
	fmt.Println(cppPointer)
	// err - Invalid operation: cppPointer++ (non-numeric type *string)
	//cppPointer++
	fmt.Println()

	// Передача указателей в функции.
	// !!! Передача указателя сокращает объем используемой памяти/ (новые переменные не создаются).
	// !!! При обычной передаче аргументов (в виде значений) - происходит их копирование.
	sample := 1
	samplePointer := &sample
	fmt.Println("Origin value of sample:", sample)
	changeParam(samplePointer) // == 101
	fmt.Println("After changing sample is:", sample)

	changeWOPointer(sample) // значение не изменится.
	fmt.Println("After local changing sample is:", sample)
	fmt.Println()

	// Best practice - возврат указателя из функции (особенно полезно при работе с конструктором).
	// Два разных указателя.
	ptr1 := returnPointer()
	ptr2 := returnPointer()
	fmt.Printf("Ptr1: %T and address: %v and value: %v\n", ptr1, ptr1, *ptr1)
	fmt.Printf("Ptr2: %T and address: %v and value: %v\n", ptr2, ptr2, *ptr2)

}

package main // пакет main

import (
	"fmt"                   // стандартные пакеты берутся из GOROOT
	"l22_package/rectangle" // путь к пакету rectangle
)

// Разделяющая область видимости.
// 2 модуля main и calculator размещены в одной директории - l22_package.
// == Модули main и calculator находятся в одном пакете (директории).

// Разделяющая область видимости - все, что находится внутри данного пакета - доступно из любого модуля без импортирования.

// !!! Функция main - может быть одна на весь ПРОЕКТ.
func main() {
	//// Функции видны компилятору, т.к. модуль calculator входит в состав пакета main.
	//resAdd := Add(10, 20)
	//resSub := Sub(30, 40)
	//resMult := Mult(50, 4)
	//resDiv := Div(40, 4)
	//
	//fmt.Println(resAdd, resSub, resMult, resDiv)
	//fmt.Println()

	r := rectangle.New(10, 20, "green")
	fmt.Println("Perimeter:", r.Perimeter())
	fmt.Println()

	// Экспортируемость как инкапсуляция.
	newR := rectangle.Rectangle{
		A: 10,
		B: 7,
		// err - Unexported field 'color' usage. поле color в пакете rectangle написано с маленькой буквы - не экспортируемо.
		//color: 12,
	}
	fmt.Println(newR)
}

// go run l22_package/main.go
// => err
//l22_package/main.go:13:12: undefined: Add
//l22_package/main.go:14:12: undefined: Sub
//l22_package/main.go:15:13: undefined: Mult
//l22_package/main.go:16:12: undefined: Div
// => правильно
// v1 - go run l22_package/main.go l22_package/calculator.go - т.к. язык компилируемый - нужно указать все файлы
// v2 - go build l22_package/main.go l22_package/calculator.go - создает бинарный исполняемый файл. Выполнить его - ./main
// v3 - go install l22_package/main.go l22_package/calculator.go - создает бинарный исполняемый файл и кладет его в GOPATH/bin (== GOBIN),

// !!! Функции, которые могут быть экспортированы в main - начинаются с большой буквы.
// Если имя сущности (переменная/функция/поле в структуре/структура/метод/интерфейс) - сущность можно экспортировать.
// (т.е. можно передавать в другие модули и пакеты)
// Если имя сущности начинается с маленькой буквы - данная сущность неэкспортируема (нельзя передавать за пределы данного ПАКЕТА).

// 1. Создаем дополнительный пакет rectangle.
// mkdir rectangle && cd rectangle && touch rectangle.go
// Реализуем внутри rectangle структуру, метод и конструктор.
// 2. Определим альтернативные пути.
// Альтернативный путь - определяет точку входа в ПРИЛОЖЕНИЕ. Для создания альтернативного пути необходимо инициализировать файл go.mod
// go.mod - будет отправной точкой для импортирования всех пакетов внутри проекта.
// go mod init <project_name>
// cd l22_package (!!! КОРНЕВАЯ директория проекта) && go mod init l22_package
// 3. Импортируем пакет rectangle.
// Создаем экземпляр + комментируем вызов функций.
// cd l22_package && go run main.go - OK
// cd .. && go run main.go - err - l22_package/main.go:4:2: package l22_package/rectangle is not in GOROOT (/usr/local/go/src/l22_package/rectangle)
// 4. Наличие файла go.mod упрощает сборку - компилятор получает отправную точку из go.mod
// cd l22_package && go run main.go && go build && go install

// !!! Всегда использовать инициализацию go.mod

// Называть пакеты именами стандартных пакетов - можно, но не рекомендуется.

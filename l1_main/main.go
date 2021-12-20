// Каждая программа начинается с директивы package <name> - указывает на принадлежность go модуля к исполняемому пакету.
package main

import "fmt"

// main - точка входа в программу.
func main() {
	fmt.Println("Hello world!")
}

// go build <path_to_file.go> - создает (компилирует) исполняемый файл в месте вызова команды.
// go build l1/main.go
// ./main - запуск получившегося исполняемого файла.

// go install <path_to_file.go> - создает исполняемый файл в GOPATH/bin.
// go install l1/main.go

// go run <path_to_file.go> - создает исполняемый файл в tmp-директории + запускает этот файл + удаляет/кэширует его.
// --work - отображать путь к временной директории.
// go run l1/main.go
// go run --work l1/main.go

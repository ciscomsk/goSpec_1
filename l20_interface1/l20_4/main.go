package main

import "fmt"

func DoSomething(pretender interface{}) {
	// Извлечение нижележащего типа.
	switch pretender.(type) {
	case string:
		fmt.Println("This is string!")
	case int:
		fmt.Println("This is int!")
	default:
		fmt.Println("Unknown type!")
	}
}

func main() {
	DoSomething(10)
	DoSomething("Hello world!")
	DoSomething(func(a, b int) int { return a + b })
}

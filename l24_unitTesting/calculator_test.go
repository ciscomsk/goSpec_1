package main

import (
	"log"
	"testing"
)

// На каждый юнит - создается своя тестирующая функция (Test<unit_name>).
// Все, что начинается с Test - будет запущено командой go test

// cd l24_unitTesting && go mod init l24_unitTesting && go test
// go test -v - более детальный выводю

// Coverage (покрытие) - показывает сколько % кода покрыто модульными тестами.
// go test -cover -v
// Достаточно - 75-80% coverage.

func TestAdd(t *testing.T) {
	// 1-й тест кейс.
	if res := Add(10, 20); res != 30 {
		// log.Fatalf - завершает тестирование на первой ошибке.
		log.Fatalf("Add(10, 30) fail. Expected %d, got %d\n", 30, res)

		// !!! сообщение будет выведено - но статус теста будет PASS.
		//fmt.Printf("Add(10, 30) fail. Expected %d, got %d\n", 30, res)
	}

	// 2-й тест кейс
	if res := Add(30, 30); res != 60 {
		log.Fatalf("Add(10, 30) fail. Expected %d, got %d\n", 60, res)
	}
}

func TestSub(t *testing.T) {
}

func TestMult(t *testing.T) {
}

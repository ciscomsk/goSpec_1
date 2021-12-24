package main

import (
	"fmt"
	"unicode/utf8"
)

type BigWord interface {
	IsBig() bool
}

type SuperString string

func (ss SuperString) IsBig() bool {
	if utf8.RuneCountInString(string(ss)) > 10 {
		return true
	}
	return false
}

func main() {
	sample := SuperString("akjsdlaskhdkasjljdkjasd")

	// Объявление переменной типа BigWord.
	var interfaceSample BigWord
	// Присваивание возможно т.к. тип переменной sample - SuperString - реализует интерфейс BigWord (поэтому интерфейсы - ПОЛУ-АБСТРАКТНЫЕ)
	interfaceSample = sample
	fmt.Println("IsBig?", interfaceSample.IsBig())

	var interfaceBadSample BigWord
	// err - '"string"' (type string) cannot be represented by the type BigWord - string не реализует интерфейс BigWord.
	//interfaceBadSample = "string"
	fmt.Println(interfaceBadSample)
	fmt.Println()
}

package main

import "fmt"

type Rider interface {
	Ride()
	Gas()
	Stop()
}

func main() {
	var r Rider
	// Значение и тип для "нулевого" экземпляра интерфейса - nil
	if r == nil {
		fmt.Printf("r is nil. Value: %v and type: %T\n", r, r)
	}

	// err - panic: runtime error: invalid memory address or nil pointer dereference
	// !!! У экземпляра интерфейса можно вызвать значение, но т.к. значение == nil => nil.Ride()
	// Если код падает с invalid memory address or nil pointer dereference - 99%, что
	// происходит обращение к нулевому экземпляру интерфейса.
	r.Ride()
}

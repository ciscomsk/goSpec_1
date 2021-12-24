package circle

import "fmt"

func init() {
	fmt.Println("Init function from circle package!")
}

type Circle struct {
	Radius float64
}

func New(newRadius float64) *Circle {
	return &Circle{newRadius}
}

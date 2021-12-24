package rectangle

type Rectangle struct {
	A, B  int    // с большорй буквы
	color string // с маленькой буквы - данное поле не экспортируемо (т.е. ее область видимости - ПАКЕТ)
}

func New(newA int, newB int, newColor string) *Rectangle {
	return &Rectangle{newA, newB, newColor}
}

func (r *Rectangle) Perimeter() int { // сбольшой буквы
	return 2 * (r.A + r.B)
}

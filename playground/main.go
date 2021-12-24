package main

import (
	"fmt"
	"math/big"
)

func Factorial(i int) *big.Int {
	if i > 0 {
		res := new(big.Int).SetInt64(int64(1))
		for j := 1; j <= i; j++ {
			res.Mul(res, new(big.Int).SetInt64(int64(j)))
		}
		return res
	}
	return new(big.Int).SetInt64(int64(1))
}

func combination(n int, m int) *big.Int {
	// !!! При 30 11 - переполнение
	//if n >= m {
	//	return Factorial(n) / (Factorial(m) * Factorial(n - m))
	//}
	//return 0
	factN := Factorial(n)
	factM := Factorial(m)
	factNsubM := Factorial(n - m)
	return new(big.Int).Div(factN, new(big.Int).Mul(factM, factNsubM))
}

func main() {
	var set, subset int

	fmt.Scan(&set, &subset)
	//fmt.Println(^uint(0))	// MaxUint
	//fmt.Printf("%T, %v\n", Factorial(set), Factorial(set))
	//fmt.Printf("%T, %v\n", Factorial(subset), Factorial(subset))
	//fmt.Printf("%T, %v\n", Factorial(19), Factorial(19))
	//fmt.Printf("%T, %v\n", new(big.Int).Mul(Factorial(subset), Factorial(19)), new(big.Int).Mul(Factorial(subset), Factorial(19)))
	fmt.Println(combination(set, subset))

	// 18446744073709551615 == MaxUint
	// 265252859812191058636308480000000 == Factorial(30)
	// 4855683143999265177600000 == Factorial(11) * Factorial(19)
}

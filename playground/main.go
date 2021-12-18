package main

import (
	"fmt"
)

func main() {
	var (
		first      int = 1
		second     int
		third      int
		totalSum   int
		totalCount int
		sumLeft    int
		sumLeft2   int
		sumLeft3   int
		count      int
		count2     int
		count3     int
	)

	fmt.Scan(&totalSum, &totalCount)
	for i := 1; i <= totalCount; i++ {
		first = i

		if totalSum-i*20 >= 0 {
			sumLeft = totalSum - i*20
			count += first

			//fmt.Println(sumLeft)
			//fmt.Println(count)

			if sumLeft == 0 && count == totalCount {
				fmt.Println(first, second, third)
			} else {
				for j := 0; j <= totalCount-first; j++ {
					second = j
					sumLeft2 = sumLeft - j*10
					count2 = j

					if sumLeft2 == 0 && count+count2 == totalCount {
						fmt.Println(first, second, third)
					} else {
						for k := 0; k <= totalCount-first-second; k++ {
							third = k
							sumLeft3 = sumLeft2 - k*5
							count3 = k

							if sumLeft3 == 0 && count+count2+count3 == totalCount {
								fmt.Println(first, second, third)
							}

							sumLeft3 = sumLeft2
							count3 = 0

						}

						sumLeft2 = sumLeft
						count2 = 0
					}
				}

				sumLeft = totalSum
				count = 0
				second = 0
			}
		}
	}
}

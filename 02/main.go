package main

import (
	"fmt"
)

var limit, sum, result int = 4000000, 0, 0

//The project Euler problem considers the firts 10 numbers = 1, 2, 3, 5, 8, 13, 21, 34, 55, 89
//This is wrong as the sequence starts as 1,1,2,3...
//To be 1,2,3,5... just change the if to a <= 1
func fibRec(a int) int {
	if a < 2 {
		return a
	}
	return fibRec(a-1) + fibRec(a-2)
}

func main() {
	for i := 1; result < limit; i++ {
		result := fibRec(i)
		if result > limit {
			break
		}

		if result%2 != 0 {
			sum += result
		}
	}

	fmt.Println(sum)
}

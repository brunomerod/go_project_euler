package main

import (
	"fmt"
)

var limit, sum, result int = 4000000, 0, 0

//The project Euler problem considers the firts
//10 numbers = 1, 2, 3, 5, 8, 13, 21, 34, 55, 89
//This is wrong as the sequence starts as 1,1,2,3...
//To be 1,2,3,5... just change the if to a <= 1
func fibRec(n int) int {
	if n < 2 {
		return n
	}
	return fibRec(n-1) + fibRec(n-2)
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func main() {
	for i := 1; fibRec(i) < limit; i++ {
		result = fibRec(i)

		switch {
		case result > limit:
			break
		case isEven(result):
			sum += result
		}
	}

	fmt.Println(sum)
}

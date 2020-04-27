package main

import (
	"fmt"
)

const count int = 1000

var sum = 0

//Refactored to add a function to check if it's divisible,
//this is faster.
func findDiv(n int) bool {
	switch {
	case n%3 == 0, n%5 == 0:
		return true
	}
	return false
}

func main() {
	for i := 1; i < count; i++ {
		if findDiv(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}

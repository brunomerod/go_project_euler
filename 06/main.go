package main

import (
	"fmt"
)

func main() {
	var fn int
	fmt.Printf("First numbers: ")
	_, err := fmt.Scanf("%d", &fn)

	var sqr int
	var sum int

	for i := 1; i <= fn; i++ {
		sqr = sqr + (i * i)
	}

	for i := 1; i <= fn; i++ {
		sum = sum + i
	}

	sum = sum * sum

	if err != nil || fn < 0 {
		fmt.Println("Not possible")
	} else {
		fmt.Printf("\n%d - %d = %d\n", sum, sqr, sum-sqr)
	}

}

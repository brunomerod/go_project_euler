package main

import (
	"fmt"
)

func isPrime(n int) bool {
	var div []int

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			div = append(div, i)
		}
		//fmt.Println(".")
		if len(div) > 2 {
			break
		}
	}

	//fmt.Println(div)
	if len(div) == 2 {
		return true
	}

	return false
}

func main() {
	var np int
	fmt.Printf("The nth prime: ")
	_, err := fmt.Scanf("%d", &np)
	//For project Euler, 1 is not a prime :(
	for n, i := 2, 1; i <= np; n++ {
		if isPrime(n) {
			i++

			//fmt.Printf("%d, ", n)

			//Print only the last one
			if i > np {
				fmt.Println("\nNth number", n)
			}
		}
	}

	//fmt.Println(n)
	if err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"fmt"
)

func isPrime(n int) bool {
	var div int = 1
	for i := 2; i < n; i++ {
		if n%i == 0 {
			div++
		}
		if div > 1 {
			return false
		}
	}

	return true
}

func main() {
	//max number
	var mn int = 2000
	var s int = 0
	/*
		fmt.Printf("Valor máximo: ")
		_, err := fmt.Scanf("%d", &mn)

		if err != nil {
			fmt.Println(err)
		}
	*/
	for i := 2; i <= mn; i++ {
		if isPrime(i) {
			s = s + i
		}
	}

	fmt.Printf("A soma dos primos até o valor máximo é\n%d\n", s)

}

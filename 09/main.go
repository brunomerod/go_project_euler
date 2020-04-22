package main

import (
	"fmt"
)

func bruteWay(found bool, a, b, c, s int) {
	for a = 1; a < s/3; a++ {
		for b = a; b < s/2; b++ {
			c = s - a - b
			if (a*a)+(b*b) == (c * c) {
				found = true
				break
			}
		}
		if found == true {
			break
		}
	}

	fmt.Printf("%d + %d + %d = %d\n", a, b, c, s)
	fmt.Printf("%d * %d * %d = %d\n", a, b, c, a*b*c)
}

//TODO:  a more mathematical aproach using Euclid's algorithm
func main() {
	//Brute Force Way
	var a, b, c, s int = 0, 0, 0, 100
	var found bool = false

	bruteWay(found, a, b, c, s)

}

package main

import (
	"fmt"
	"math"
)

/*
Formula for triangle numbers
(n*(n+1))/2

The 7th triangle number
(7*(7+1))/2 = (7*8)/2 = 28

There's no need to iterate all integers
*/

//TriangleNumber : Creates Triangle Number
func TriangleNumber(n int) int {
	return (n * (n + 1)) / 2
}

//TODO: Read the overview https://projecteuler.net/overview=012
//and try to implement it
func main() {
	for i := 1; ; i++ {
		//Divisors
		var div int = 0
		//The triangle Number to nth (i)
		/*
		* The divisors of a sqrt root * 2 gives the total divisors.
		* -1 if it's a perfect sqrt
		 */
		n := float64(TriangleNumber(i))
		//Int for sqrt
		sqrt := int(math.Sqrt(n))
		for s := 1; s <= sqrt; s++ {
			if int(n)%s == 0 {
				div += 2
			}
		}

		if sqrt*sqrt == int(n) {
			div--
		}

		if div > 500 {
			fmt.Printf("Nth = %d\n", i)
			fmt.Println(int(n))
			break
		}

	}
}

package main

/*
*	Brute force method. It'll first get the palindrome with 99 * 99[...].
*	Then it'll enter another loop taking into consideration the first palindrome found.
*	It's working well up to 4 digits.
 */

import (
	"fmt"
	"strconv"
)

// Verify if it's a palindrome, string and back for comparisson.
func isPalindrome(n int) bool {
	strN := strconv.Itoa(n)
	revStr := ""

	for lenght := len(strN); lenght > 0; lenght-- {
		revStr = revStr + string(strN[lenght-1])

		if revStr == strN {
			return true
		}
	}
	return false
}

//Will verify every possible product
func allPossibleProducts(n int) {
	/*
	   99 * 99
	   98 * 99 (98 * 98 / 97 * 98 / 96 * 98)
	   97 * 99	(97 * 97 / 96 * 97 / 95 * 97)
	   96 * 99
	   95 * 99
	   ...
	*/
	//var palis []int
	var largePali, x, y int
	var minP int

	for i := 1; i < n; i++ {
		if isPalindrome((n) * (n - i)) {
			x, y, largePali = n, n-i, n*(n-i)
			minP = n - i
			break
		}
	}

	//Yes, I'm aware of this mess...
	for i := n; i > minP; i-- {
		for b := 0; b < n-minP; b++ {
			if isPalindrome(i*(n-b)) && i*(n-b) > largePali {
				x, y, largePali = i, n-b, i*(n-b)
				minP = n - i
				fmt.Printf(".")
				break
			}
		}
		//fmt.Printf("=")
	}

	fmt.Printf("\n%d * %d = %d \n", x, y, largePali)
	return
}

func main() {
	var digits int
	fmt.Printf("N of digits: ")
	_, err := fmt.Scanf("%d", &digits)

	//Easier to transform string to int
	var p string = ""

	//Forming the var to products
	for i := 1; i <= digits; i++ {
		p = p + "9"
	}

	p1, err := strconv.Atoi(p)

	fmt.Printf("Base: %d \n", p1)

	allPossibleProducts(p1)

	if err != nil {
		fmt.Println(err)
	}

}

package main

import (
	"fmt"
)

/*
* 	Brute force. Multiples of 9 divisible by 20 first in array. I'm sure they
* 	will be divisible by 3,5, 9,10,20....!no
* 	Change of plans. IF % with 1-20.
* 	Nested fors n²
 */

func main() {
	//This is the fastest LOL
	var i int = 1

	for {
		if i%1 == 0 &&
			i%2 == 0 &&
			i%3 == 0 &&
			i%4 == 0 &&
			i%5 == 0 &&
			i%6 == 0 &&
			i%7 == 0 &&
			i%8 == 0 &&
			i%9 == 0 &&
			i%10 == 0 &&
			i%11 == 0 &&
			i%12 == 0 &&
			i%13 == 0 &&
			i%14 == 0 &&
			i%15 == 0 &&
			i%16 == 0 &&
			i%17 == 0 &&
			i%18 == 0 &&
			i%19 == 0 &&
			i%20 == 0 {
			fmt.Printf("First: %d\n", i)
			break
		}
		//If I use i = i + 20 it's really slower
		i++
	}

	//Fors n²
	for i := 1; ; i++ {
		for n := 1; n <= 20; n++ {
			if i%n != 0 {
				break
			}
			if n == 20 {
				fmt.Printf("Second: %d\n", i)
				return
			}
		}
	}
}

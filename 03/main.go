package main

/*
Decided to write a code to get all (array) prime factors of any number [the target].
*/

import (
	"fmt"
	"math/big"
)

var target int = 600851475143
var primes []int

func multiplyArr(arr []int) int {
	var r int = 1
	for i := 0; i < len(arr); i++ {
		r = arr[i] * r
	}
	return r
}

//TODO: This is really slow...
func main() {
	//create a copy of target (yes should split all parts bellow into func)
	var dummyTarget int = target

	//Check all the 2s
	for dummyTarget%2 == 0 {
		primes = append(primes, 2)
		dummyTarget = dummyTarget / 2
	}

	//Find all the primes to the dummyTarget, if is prime, try to divide,
	//if % == 0 then sum
	for i := 3; i <= dummyTarget; i = i + 2 {
		//Function from the math/big library to find primes til 2^64
		if big.NewInt(int64(i)).ProbablyPrime(0) {
			for dummyTarget%i == 0 {
				primes = append(primes, i)
				dummyTarget = dummyTarget / i
			}
		}

		if multiplyArr(primes) == target {
			fmt.Println(primes)
			return
		}

	}
}

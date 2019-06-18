package main

import (
	"fmt"
)

var count = 1000
var sum = 0

func main() {
	for i := 0; i < count; i++ {
		switch {
		case i%3 == 0, i%5 == 0:
			sum += i
			//fmt.Println(i)
		}
	}
	fmt.Println(sum)
}

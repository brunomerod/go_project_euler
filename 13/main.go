package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
)

func main() {
	var n, result = new(big.Int), new(big.Int)

	//File on memory
	file, err := os.Open("numbers.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Bufio to read all the lines
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Sscan convertion to the type of n
		fmt.Sscan(scanner.Text(), n)
		//Addition with big int. result += n doesn't work.
		result = result.Add(result, n)
	}

	fmt.Printf("The sum is: %d \n", result)
	fmt.Printf("The first 10 digits are: %s \n", result.String()[:10])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

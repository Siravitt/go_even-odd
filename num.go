package main

import "fmt"

type num []int

func createNum() num {
	numSlice := num{}

	for i := 0; i <= 10; i++ {
		numSlice = append(numSlice, i)
	}
	return numSlice
}

func (n num) checkEvenOrOdd() {
	for _, number := range n {
		if number%2 == 0 {
			fmt.Printf("Number %d is Even\n", number)
		} else {
			fmt.Printf("Number %d is Odd\n", number)
		}
	}
}

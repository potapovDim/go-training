package main

import "fmt"

func main() {
	var testArr [10]int
	var i, j int

	for i = 0; i < len(testArr); i++ {
		testArr[i] = i + 100
		j = 0
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, testArr[j])
	}
}

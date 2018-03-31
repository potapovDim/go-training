package main

import "fmt"

func someRun(str string, num1 int, num2 int) int {
	var result int

	if str == "test" {
		result = num1 + num2
	} else if str == "tes1" {
		result = num1 - num2
	}
	return result
}

func main() {
	fmt.Println(someRun("test", 2, 4))
}

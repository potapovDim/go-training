package main

import "fmt"

/* global variable declaration */
var a int = 27

var a int = 700

func main() {
	var b int = 20
	var c int = 0

	fmt.Printf("value of a in main() = %d\n", a)
	c = sum(a, b)
	fmt.Printf("value of c in main() = %d\n", c)
}

func sum(a, b int) int {
	fmt.Printf("value of a in sum() = %d\n", a)
	fmt.Printf("value of b in sum() = %d\n", b)

	return a + b
}

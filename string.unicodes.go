package main

import "fmt"

func main() {
	var str string = "Hello world !"

	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}
}

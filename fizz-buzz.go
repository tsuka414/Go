package main

import (
	"fmt"
)

func main() {

	i:= 1
	for i < 101 {
		switch {
		case i % 15 == 0:
			fmt.Println("fizz-buzz")
		case i % 5 == 0:
			fmt.Println("fizz")
		case i % 3 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}

		i++
	}

}
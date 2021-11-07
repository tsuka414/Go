package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

// defer へ渡した関数は　LIFO の順番で実行される
// https://go-tour-jp.appspot.com/flowcontrol/13
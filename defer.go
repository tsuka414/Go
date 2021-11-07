package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

// defer へ渡した関数の引数は、すぐに評価されるが、
// その関数は呼び出し元の関数が return するまで実行されません。
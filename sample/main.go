package main

import (
	"fmt"
)

func Mul(x int, y int) int {
	return x * y
}

func Min(a int, b int) int {
	if a > b{
		return b
	}
	return a
}

func main() {
	fmt.Println("Hello World")
}

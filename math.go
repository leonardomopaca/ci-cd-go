package main

import "fmt"

func main() {
	fmt.Println(soma(112, 10))
	fmt.Println(mult(2, 4))
}

func soma(a int, b int) int {
	return a + b
}

func mult(a int, b int) int { return a * b }

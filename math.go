package main

import "fmt"

func main() {
	fmt.Println(soma(112, 10))
	fmt.Println(multiplica(2, 4))
}

func soma(a int, b int) int {
	return a + b
}

func multiplica(a float64, b float64) float64 { return a * b }

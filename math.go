package main

import "fmt"

func main() {
	fmt.Println(soma(10, 10))
}

func soma(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func times(a int, b int) int {
	return a * b
}

func somaX(a int, b int) int {
	return a + b + a + a
}

func subX(a int, b int) int {
	return a - b - a
}

func timesX(a int, b int) int {
	return a * b * a
}
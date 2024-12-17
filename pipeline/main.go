package main

import "fmt"

func main() {
	value := 1

	fmt.Println(add(multiply(value, 2), 3))

	fmt.Println(multiply(add(value, 2), 3))
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

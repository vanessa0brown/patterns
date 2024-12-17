package main

import "fmt"

func main() {
	items := []int{10, 20, 30, 40, 50}

	dataChannel := generator(items)

	process(dataChannel)
}

func generator(items []int) chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for _, item := range items {
			ch <- item
		}
	}()

	return ch
}

func process(ch chan int) {
	for item := range ch {
		fmt.Println(item)
	}
}

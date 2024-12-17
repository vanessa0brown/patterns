package main

import "fmt"

func add(doneCh chan struct{}, inputCh chan int) chan int {
	resultChan := make(chan int)

	go func() {
		defer close(resultChan)

		for value := range inputCh {
			result := value + 2

			select {
			case <-doneCh:
				return
			case resultChan <- result:
			}
		}
	}()

	return resultChan
}

func multiply(doneCh chan struct{}, inputCh chan int) chan int {
	resultChan := make(chan int)

	go func() {
		defer close(resultChan)

		for value := range inputCh {
			result := value * 3

			select {
			case <-doneCh:
				return
			case resultChan <- result:
			}
		}
	}()

	return resultChan
}

func generator(doneCh chan struct{}, numbers []int) chan int {
	outputCh := make(chan int)

	go func() {
		defer close(outputCh)

		for _, num := range numbers {
			select {
			case <-doneCh:
				return
			case outputCh <- num:
			}
		}
	}()

	return outputCh
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	doneCh := make(chan struct{})
	defer close(doneCh)

	inputCh := generator(doneCh, numbers)

	addCh := add(doneCh, inputCh)
	resultCh := multiply(doneCh, addCh)

	for res := range resultCh {
		fmt.Println(res)
	}
}

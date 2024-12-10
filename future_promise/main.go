package main

import (
	"fmt"
	"time"
)

func Promise(task func() int) chan int {
	resultChan := make(chan int, 1)

	go func() {
		result := task()
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

func main() {
	longRunningTask := func() int {
		time.Sleep(2 * time.Second)
		return 42
	}

	future := Promise(longRunningTask)

	fmt.Println("Start task")

	result := <-future

	fmt.Println("Finish task")
	fmt.Println("Result:", result)
}

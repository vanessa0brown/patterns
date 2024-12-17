package main

import (
	"fmt"
	"time"
)

type Result struct {
	value int
	err   error
}

func Promise(task func() (int, error)) chan Result {
	resultChan := make(chan Result, 1)

	go func() {
		value, err := task()
		resultChan <- Result{value, err}
		close(resultChan)
	}()

	return resultChan
}

func main() {
	taskWithError := func() (int, error) {
		time.Sleep(2 * time.Second)
		//return 0, errors.New("Something went wrong")
		return 42, nil
	}

	future := Promise(taskWithError)

	fmt.Println("Start task")

	result := <-future
	if result.err != nil {
		fmt.Println("Error: ", result.err)
	} else {
		fmt.Println("Result: ", result.value)
	}
}

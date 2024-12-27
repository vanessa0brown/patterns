package main

import (
	"errors"
	"fmt"
)

type Result struct {
	data int
	err  error
}

func main() {
	input := []int{1, 2, 3, 4, 5}

	resultCh := make(chan Result)

	go consumer(generator(input), resultCh)

	for res := range resultCh {
		if res.err != nil {
			fmt.Println("Ошибка", res.err)
		} else {
			fmt.Println("Результат:", res.data)
		}
	}
}

func generator(input []int) chan int {
	resultCh := make(chan int)

	go func() {
		defer close(resultCh)

		for _, data := range input {
			resultCh <- data
		}
	}()

	return resultCh
}

func consumer(inputCh chan int, resultCh chan Result) {
	defer close(resultCh)

	for data := range inputCh {
		resp, err := CallDataBase(data)
		resultCh <- Result{data: resp, err: err}
	}
}

func CallDataBase(data int) (int, error) {
	return data, errors.New("Ошибка обращения к бд")
}

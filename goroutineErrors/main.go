package main

import (
	"context"
	"errors"
	"log"

	"golang.org/x/sync/errgroup"
)

type Result struct {
	data int
	err  error
}

func main() {
	g, _ := errgroup.WithContext(context.Background())
	input := []int{1, 2, 3, 4, 5}

	inputCh := generator(input)

	for data := range inputCh {
		data := data
		g.Go(func() error {
			_, err := CallDataBase(data)
			if err != nil {
				return err
			} else {
				return nil
			}
		})
	}

	if err := g.Wait(); err != nil {
		log.Println("Ошибка: ", err)
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
	if data == 3 {
		return data, errors.New("Ошибка обращения к бд")
	} else {
		return data, nil
	}
}

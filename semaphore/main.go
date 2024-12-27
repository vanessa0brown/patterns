package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	semaCh chan struct{}
}

func NewSemaphore(maxReq int) *Semaphore {
	return &Semaphore{
		semaCh: make(chan struct{}, maxReq),
	}
}

func (s *Semaphore) Acquire() {
	s.semaCh <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.semaCh
}

func main() {
	var wg sync.WaitGroup

	sema := NewSemaphore(3)

	for i := range 10 {
		wg.Add(1)

		go func(taskId int) {
			defer wg.Done()
			sema.Acquire()

			fmt.Printf("Запущен рабочий %v \n", taskId)
			time.Sleep(1 * time.Second)
			fmt.Printf("Процесс %v завершен \n", taskId)

			wg.Done()
			sema.Release()
		}(i)
	}

	wg.Wait()
}

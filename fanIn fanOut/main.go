package main

import (
  "sync"
  "time"
)

func genenerator(doneCh chan struct{}, numbers []int) chan int {
	dataStream := make(chan int)
	go func() {
		defer close(dataStream)

		for _, num := range numbers {
			select {
			case <-doneCh:
				return
			case dataStream <- num:
			}
		}
	}()

	return dataStream
}

func add(doneCh chan struct{}, inputCh chan int) chan int {
	resultStream := make(chan int)

    go func() {
      defer close(resultStream)

      for num := range inputCh {
        time.Sleep(1 * time.Second)
        result := num + 1

        select {
          case <-doneCh:
            return
          case resultStream<-result
        }

      }
    }()

    return resultStream
}

func multiply(doneCh chan struct{}, inputCh chan int) chan int {
  resultStream := make(chan int)

  go func() {
    defer close(resultStream)

    for num := range inputCh {
      result:= num * 2

      select {
        case <-doneCh:
          return
        case resultStream<-result:
      }
    }
  }()

  return resultStream
}

func fanOut(doneCh chan struct{}, inputCh chan int, workers int) []chan int {
  resultChan := make([]chan int, workers)

  for i := range workers {
    resultChan[i] = add(doneCh, inputCh)
  }

  return resultChan
}

func fanIn(doneCh chan struct{}, channels ...chan int) chan int {
  finalStream := make(chan int)
  var wg sync.WaitGroup

  for _, ch := range channels {
    chCopy := ch
    wg.Add(1)

    go func() {
      defer wg.Done()
        for value := range chCopy {
          select {
            case <-doneCh:
              return
            case finalStream<-value:
          }
        }
    }()
  }

  go func() {
    wg.Wait()
    close(finalStream)
  }()

  return finalStream
}

func main() {

}

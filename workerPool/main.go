package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Printf("Рабочий %v начал работу над задачей %v\n", id, job)
		time.Sleep(1 * time.Second)
		fmt.Printf("Рабочий %v закончил работу над задачей %v\n", id, job)
		result <- job * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := range numJobs {
		jobs <- j
	}

	close(jobs)
	
	for r := 1; r <= numJobs; r++ {
		res := <-results
		fmt.Println(res)
	}
}

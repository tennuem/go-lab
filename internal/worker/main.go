package main

import (
	"fmt"
	"time"
)

func main() {
	workers := 4
	jobs := 1000
	jobCh := make(chan int, 100)
	resCh := make(chan int, 100)

	for i := 1; i < workers; i++ {
		go worker(i, jobCh, resCh)
	}

	for j := 1; j < jobs; j++ {
		jobCh <- j
	}
	close(jobCh)

	var result []int
	for i := 1; i < jobs; i++ {
		result = append(result, <-resCh)
	}

	fmt.Println("result", result)
}

func worker(id int, jobs chan int, result chan int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		result <- j * 2
	}
}

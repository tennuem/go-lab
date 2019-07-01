package main

import (
	"fmt"
	"sync"
)

func main() {
	quit := make(chan struct{})
	defer close(quit)

	ch1 := write(quit)
	ch2 := write(quit)

	for n := range merge(quit, ch1, ch2) {
		fmt.Println(n)
	}
}

func write(quit <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < 10; i++ {
			select {
			case out <- i:
			case <-quit:
				return
			}
		}
	}()

	return out
}

func merge(quit <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				select {
				case out <- n:
				case <-quit:
					return
				}
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

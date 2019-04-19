package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var outChan chan int
var quit chan bool

func main() {
	outChan = make(chan int)
	quit = make(chan bool)

	ch1 := make(chan int)
	ch2 := make(chan int)

	wg.Add(2)
	go mergeChanels(ch1, ch2)
	go read()

	ch1 <- 1
	ch2 <- 2

	close(ch1)
	close(ch2)

	wg.Wait()
}

func mergeChanels(inChan ...chan int) {
	defer func() {
		wg.Done()

	}()

	wg.Add(len(inChan))

	for _, ch := range inChan {
		go func(ch chan int) {
			defer wg.Done()
			for {
				v, ok := <-ch
				if !ok {
					return
				}
				outChan <- v
			}
		}(ch)
	}

	wg.Wait()
}

func read() {
	defer wg.Done()
	for v := range outChan {
		fmt.Println(v)
	}
}

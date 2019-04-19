package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	runtime.GOMAXPROCS(1)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("go routine")
	}()
	wg.Wait()

	fmt.Println("num CPU: ", runtime.NumCPU())
	fmt.Println("num go routine: ", runtime.NumGoroutine())
}

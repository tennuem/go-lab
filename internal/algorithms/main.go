package main

import (
	"fmt"

	"github.com/tennuem/go-lab/internal/algorithms/sort/merge"
)

func main() {
	slice := []int{5, 4, 2, 3, 1, 9, 7, 8, 0}

	//sort := bubble.NewBubble()
	//sort := insertion.NewInsertion()
	sort := merge.NewMerge()

	fmt.Println("Result:", sort.Result(slice))
	fmt.Println("Time duration:", sort.Duration())
	fmt.Println("Count operations:", sort.Operation())
}

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Fprintf(os.Stdout, "\r%d%%", i)
		time.Sleep(time.Second * 1)
	}
}

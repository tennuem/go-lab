package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "вывод промежуточных результатов")
var wg sync.WaitGroup
var sema = make(chan struct{}, 20)

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSize := make(chan int64)

	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, fileSize)
	}

	go func() {
		wg.Wait()
		close(fileSize)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(time.Millisecond * 500)
	}

	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSize:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d файлов %fGB\n", nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, fileSizes chan<- int64) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "catalog walker: %v\n", err)
		return nil
	}
	return entries
}

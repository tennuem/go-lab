package main

import (
	"os/exec"
	"sync"
)

func main() {
	var mu sync.Mutex
	for {
		mu.Lock()
		err := exec.Command("ffmpeg", "-i", "inputfile.avi", "-c:v", "libx264", "-b:v", "1000k", "-c:a", "aac", "-f", "mp4", "outfile.mp4", "-y").Run()
		if err != nil {
			panic(err)
		}
		mu.Unlock()
	}
}

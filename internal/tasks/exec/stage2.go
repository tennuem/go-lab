package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("ffmpeg", "-i", "inputfile.avi", "-c:v", "libx264", "-b:v", "1000k", "-c:a", "aac", "-f", "mp4", "outfile.mp4", "-y")
	stderr, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("start")

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("read")
	slurp, _ := ioutil.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)

	fmt.Println("wait")
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}

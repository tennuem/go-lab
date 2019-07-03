package pkg

// Worker gets job from chan, after make some work and send result in other chan
func Worker(id int, jobs chan int, result chan int) {
	for j := range jobs {
		result <- j * 2
	}
}

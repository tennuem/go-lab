package pkg

func Worker(id int, jobs chan int, result chan int) {
	for j := range jobs {
		result <- j * 2
	}
}

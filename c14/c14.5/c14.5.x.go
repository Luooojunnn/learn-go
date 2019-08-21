// 使用gorutine和channel实现一个线程池

package c14

import (
	"time"
	"fmt"
)

func init()  {
	jobs := make(chan int, 100)
	res := make(chan int, 100)

	for w := 1; w < 4; w++ {
		go worker(w, jobs, res)
	}

	for i := 1; i < 5; i++ {
		jobs <- i
	}

	close(jobs)

	for a := 1; a < 5; a++ {
		<- res
	}
}

func worker(id int, job <-chan int, result chan<- int){ 
	for j := range job {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished  job", j)
		result <- j * 2
	}
}
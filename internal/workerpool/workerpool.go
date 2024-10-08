package workerpool

import (
	"fmt"
)

type Job struct {
	Req  string
	Resp chan<- string
}

// Job Queue:
var JobQueue = make(chan Job, 100)

func worker(id int, jobs <-chan Job) {
	// Exectue Work
	for j := range jobs {
		fmt.Println(j.Req)
		j.Resp <- ("Resp" + string(id))
	}
}

func startWorkerPool(numWorker int) {
	for i := 0; i < numWorker; i++ {
		go worker(i, JobQueue)
	}
}

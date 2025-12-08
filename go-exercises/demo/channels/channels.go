package main

import (
	"fmt"
	"time"
)

type ControlMsg int

const (
	Doexit = iota
	ExitOk
)

type Job struct {
	data int
}

type Result struct {
	result int
	job    Job
}

func doubler(jobs <-chan Job, results chan<- Result, control chan ControlMsg) {
	for {
		select {
		case msg := <-control:
			switch msg {
			case Doexit:
				fmt.Println("Exit goroutine")
				control <- ExitOk
				return
			default:
				panic("unhandled  control messaage")
			}
		case job := <-jobs:
			results <- Result{result: job.data * 2, job: job}
		}
	}
}

func main() {
	//Job channel to send Job to the goroutine
	//Result channel to send results back from the gorotine
	//Control channel should indicate whether the goroutine should terminate

	jobs := make(chan Job, 50)
	results := make(chan Result, 50)
	control := make(chan ControlMsg)

	go doubler(jobs, results, control)

	for i := 0; i <= 30; i++ {
		jobs <- Job{i}
	}

	for {
		select {
		case result := <-results:
			fmt.Println(result)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("timed out")
			control <- Doexit
			<-control
			fmt.Println("program exited")
			return
		}
	}

}

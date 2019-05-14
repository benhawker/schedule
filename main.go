package main

import "fmt"

type job struct {
	fn        func([]int) int
	inputData []int
}

// Global var that holds all scheduledJobs
var scheduledJobs []job

// Add an arbitrary number of jobs
func Schedule(jobs []job) {
	for _, job := range jobs {
		scheduledJobs = append(scheduledJobs, job)
	}
}

// Iterate the jobs slice and invoke each jobFunc
func Run() {
	for _, job := range scheduledJobs {
		fmt.Println(job.fn(job.inputData))
	}
}


func sum(input []int) int {
	sum := 0

	for i := range input {
		sum += input[i]
	}

	return sum
}


// Consider this a client calling the Scheduler
func main() {
	// create a slice of functions
	jobs := []job{}

	// Append a job to it. The job sums the input slice
	fn := func(input []int) int { return sum(input) }
	jobs = append(jobs, job{fn, []int{1, 2}})

	// Append another job to it.
	fn = func(input []int) int { return sum(input) }
	jobs = append(jobs, job{fn, []int{1, 2, 3}})

	// Append another job to it.
	fn = func(input []int) int { return sum(input) }
	jobs = append(jobs, job{fn, []int{1, 2, 3, 4}})

	// Invoke Schedule to append to the global scheduledJobs var.
	Schedule(jobs)

	// Run it.
	Run()
}

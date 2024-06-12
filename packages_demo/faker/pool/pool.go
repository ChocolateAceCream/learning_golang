package pool

import "sync"

type Worker interface {
	DoJob(chan Job, *sync.WaitGroup)
}

type Job struct {
	batch types.Batch
}

func NewProcessorJob(batch types.Batch) Job {
	return Job{
		batch: batch,
	}
}

func NewWorker(workFunc func(Job, *sync.WaitGroup)) Worker {
	return &Worker{
		workFunc: workFunc,
	}
}

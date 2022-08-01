package internal

import "sync"

//NewWorkerPool creates pool for specified workers count
func NewWorkerPool[T any](workersCount int) (WorkIn[T], WorkOut[T]) {
	workIn := make(chan Worker[T])
	workOut := make(chan Result[T])

	var wg = new(sync.WaitGroup)
	wg.Add(workersCount)
	for i := 0; i < workersCount; i++ {
		go func() {
			for worker := range workIn {
				v, err := worker()

				workOut <- Result[T]{v, err}
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(workOut)
	}()

	return workIn, workOut
}

type Result[T any] struct {
	Value T
	Err   error
}

type Worker[T any] func() (T, error)

type WorkIn[T any] chan<- Worker[T]
type WorkOut[T any] <-chan Result[T]

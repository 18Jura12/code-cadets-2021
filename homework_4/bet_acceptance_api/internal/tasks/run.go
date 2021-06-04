package tasks

import (
	"context"
	"sync"
)

func RunTasks(tasks ...Task) {
	wg := &sync.WaitGroup{}
	wg.Add(len(tasks))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i, task := range tasks {
		go func(i int, task Task) {
			defer wg.Done()
			defer cancel()

			task.Start(ctx)
		}(i, task)
	}

	wg.Wait()
}

type Task interface {
	Start(ctx context.Context)
}


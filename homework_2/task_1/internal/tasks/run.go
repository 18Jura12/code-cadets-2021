package tasks

import (
	"context"
	"fmt"
	"sync"
)

func RunTasks(tasks ...Task) {
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(len(tasks))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// run each task in separate goroutine
	for i, task := range tasks {
		go func(i int, task Task) {
			defer waitGroup.Done()
			defer cancel()

			err := task.Start(ctx)
			fmt.Printf(`"%v" finished with "%v" error`, task, err)
		}(i, task)
	}

	// wait for all tasks to finish
	waitGroup.Wait()
	fmt.Println("all tasks finished")
	//
	// when first task finishes, signal to the other goroutines that application should stop

}

type Task interface {
	Start(ctx context.Context) error
}

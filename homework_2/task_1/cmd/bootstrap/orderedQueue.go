package bootstrap

import "task_1/internal/infrastructure/queue"

func NewOrderedQueue() *queue.OrderedQueue {
	return queue.NewOrderedQueue()
}

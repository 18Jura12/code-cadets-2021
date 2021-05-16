package main

import (
	"task_1/cmd/bootstrap"
	"task_1/internal/tasks"
)

func main() {
	signalHandler := bootstrap.NewSignalHandler()

	feed := bootstrap.NewAxilisOfferFeed()
	queue := bootstrap.NewOrderedQueue()
	feedProcessorService := bootstrap.NewFeedProcessorService(feed, queue)

	tasks.RunTasks(signalHandler, feed, queue, feedProcessorService)
}

package main

import (
	"task_1/cmd/bootstrap"
	"task_1/internal/domain/services"
	"task_1/internal/tasks"
)

func main() {
	signalHandler := bootstrap.NewSignalHandler()

	feed := bootstrap.NewAxilisOfferFeed()
	feed2 := bootstrap.NewAxilisOfferFeed2()
	queue := bootstrap.NewOrderedQueue()
	feedProcessorService := bootstrap.NewFeedProcessorService(queue, []services.Feed {feed, feed2})

	tasks.RunTasks(signalHandler, feed, feed2, queue, feedProcessorService)
}

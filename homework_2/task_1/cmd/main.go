package main

import (
	"code-cadets-2021/lecture_2/05_offerfeed/cmd/bootstrap"
	"code-cadets-2021/lecture_2/05_offerfeed/internal/tasks"
)

func main() {
	signalHandler := bootstrap.NewSignalHandler()

	feed := bootstrap.NewAxilisOfferFeed()
	queue := bootstrap.NewOrderedQueue()
	feedProcessorService := bootstrap.NewFeedProcessorService(feed, queue)

	tasks.RunTasks(signalHandler, feed, queue, feedProcessorService)
}

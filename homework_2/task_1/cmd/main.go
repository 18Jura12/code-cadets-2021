package main

import (
	stdhttp "net/http"
	"task_1/cmd/bootstrap"
	"task_1/internal/domain/services"
	"task_1/internal/tasks"
)

func main() {
	signalHandler := bootstrap.NewSignalHandler()

	client := stdhttp.Client{}

	feed := bootstrap.NewAxilisOfferFeed(client)
	feed2 := bootstrap.NewAxilisOfferFeed2(client)
	queue := bootstrap.NewOrderedQueue()
	feedProcessorService := bootstrap.NewFeedProcessorService(queue, []services.Feed {feed, feed2})

	tasks.RunTasks(signalHandler, feed, feed2, queue, feedProcessorService)
}

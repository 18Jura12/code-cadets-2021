package main

import (
	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/services"
	"code-cadets-2021/lecture_2/05_offerfeed/internal/infrastructure/http"
	"code-cadets-2021/lecture_2/05_offerfeed/internal/infrastructure/queue"
	"context"
	stdhttp "net/http"
	"time"
)

func main() {

	newQueue := queue.NewOrderedQueue()
	/*source := newQueue.GetSource()

	source <- models.Odd{
		Id:          "",
		Name:        "",
		Match:       "",
		Coefficient: 0,
		Timestamp:   time.Time{},
	}

	close(source)

	newQueue.Start(nil)*/

	ctx, cancel := context.WithCancel(context.Background())

	feed := http.NewAxilisOfferFeed(stdhttp.Client{})

	go feed.Start(ctx)
	go newQueue.Start(ctx)

	service := services.NewFeedProcessorService(feed, newQueue)

	go service.Start(ctx)

	time.Sleep(time.Second * 3)

	cancel()

	time.Sleep(time.Second * 3)
}

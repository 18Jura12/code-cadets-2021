package bootstrap

import "task_1/internal/domain/services"

func NewFeedProcessorService(feed services.Feed, queue services.Queue) *services.FeedProcessorService {
	return services.NewFeedProcessorService(feed, queue)
}
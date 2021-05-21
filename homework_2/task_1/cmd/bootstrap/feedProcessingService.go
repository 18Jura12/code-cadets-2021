package bootstrap

import "task_1/internal/domain/services"

func NewFeedProcessorService(queue services.Queue, feed []services.Feed) *services.FeedProcessorService {
	return services.NewFeedProcessorService(queue, feed)
}
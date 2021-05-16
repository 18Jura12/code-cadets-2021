package services

import (
	"context"
	"task_1/internal/domain/models"
)

type FeedProcessorService struct {
	feed Feed
	queue Queue
}

func NewFeedProcessorService(_feed Feed, _queue Queue) *FeedProcessorService {
	// it should receive "Feed" & "Queue" interfaces through constructor
	return &FeedProcessorService{
		_feed,
		_queue,
	}
}

func (f *FeedProcessorService) Start(ctx context.Context) error {
	// initially:
	// - get updates channel from feed interface
	// - get source channel from queue interface
	updates := f.feed.GetUpdates()
	source := f.queue.GetSource()
	//
	// repeatedly:
	// - range over updates channel
	// - multiply each odd with 2
	// - send it to source channel
	for update := range updates {
		update.Coefficient *= 2
		source <- update
	}
	close(source)
	//
	// finally:
	// - when updates channel is closed, exit
	// - when exiting, close source channel
	return nil
}

// define feed interface here
type Feed interface {
	GetUpdates() chan models.Odd
	Start(ctx context.Context) error
}

// define queue interface here
type Queue interface {
	GetSource() chan models.Odd
	Start(ctx context.Context) error
}
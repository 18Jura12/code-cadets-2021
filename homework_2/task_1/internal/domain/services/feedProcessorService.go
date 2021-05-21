package services

import (
	"context"
	"reflect"
	"task_1/internal/domain/models"
)

type FeedProcessorService struct {
	feed []Feed
	queue Queue
}

func NewFeedProcessorService(_queue Queue, _feed []Feed) *FeedProcessorService {
	return &FeedProcessorService{
		_feed,
		_queue,
	}
}

func (f *FeedProcessorService) Start(ctx context.Context) error {
	var updates []chan models.Odd
	for _, feedInstance := range f.feed {
		updates = append(updates, feedInstance.GetUpdates())
	}

	source := f.queue.GetSource()

	cases := make([]reflect.SelectCase, len(updates))
	for i, update := range updates {
		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(update),
			Send: reflect.Value{},
		}
	}

	for {
		if len(cases) == 0 {
			break
		}
		index , input, ok := reflect.Select(cases)
		if !ok {
			cases = append(cases[:index], cases[index+1:]...)
			continue
		}
		odd := input.Interface().(models.Odd)
		odd.Coefficient *= 2
		source <- odd
	}

	close(source)

	return nil
}

type Feed interface {
	GetUpdates() chan models.Odd
	Start(ctx context.Context) error
}

type Queue interface {
	GetSource() chan models.Odd
	Start(ctx context.Context) error
}

func (f *FeedProcessorService) String() string {
	return "feed processor service"
}

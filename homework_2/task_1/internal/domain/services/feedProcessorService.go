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
	// it should receive "Feed" & "Queue" interfaces through constructor
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
		index , input, ok := reflect.Select(cases)
		if !ok {
			cases = append(cases[:index], cases[index+1:]...)
			continue
		}
		if len(cases) == 0 {
			break
		}
		odd := input.Interface().(models.Odd)
		odd.Coefficient *= 2
		source <- odd
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

func (f *FeedProcessorService) String() string {
	return "feed processor service"
}

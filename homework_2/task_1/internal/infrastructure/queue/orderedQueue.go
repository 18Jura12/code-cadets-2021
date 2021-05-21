package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"task_1/internal/domain/models"

	"github.com/pkg/errors"
)

type OrderedQueue struct {
	queue  []models.Odd
	source chan models.Odd
}

func NewOrderedQueue() *OrderedQueue {
	return &OrderedQueue{
		source: make(chan models.Odd),
	}
}

func (o *OrderedQueue) Start(ctx context.Context) error {
	err := o.loadFromFile()
	if err != nil {
		return err
	}

	for sourceInput := range o.source {
		o.queue = append(o.queue, sourceInput)
	}

	fmt.Println("prošao")

	err = o.storeToFile()
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderedQueue) GetSource() chan models.Odd {
	return o.source
}

func (o *OrderedQueue) loadFromFile() error {
	f, err := os.Open("queue.txt")
	if os.IsNotExist(err) {
		return nil

	} else if err != nil {
		return errors.Wrap(err, "load from file, open")
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&o.queue)
	if err != nil {
		return errors.Wrap(err, "load from file, decode")
	}

	return nil
}

func (o *OrderedQueue) storeToFile() error {
	fmt.Println("file")
	f, err := os.Create("queue.txt")
	if err != nil {
		return errors.Wrap(err, "store to file, create")
	}
	defer f.Close()

	serializedQueue, err := json.MarshalIndent(o.queue, "", "    ")
	if err != nil {
		return errors.Wrap(err, "store to file, marshal")
	}

	n, err := f.Write(serializedQueue)
	if err != nil {
		return errors.Wrap(err, "store to file, write")

	} else if len(serializedQueue) != n {
		return errors.Wrapf(err, "store to file, write len; n: %d, serializedLen: %d", n, len(serializedQueue))
	}

	err = f.Sync()
	if err != nil {
		return errors.Wrap(err, "store to file, sync")
	}

	return nil
}

func (o *OrderedQueue) String() string {
	return "ordered queue"
}

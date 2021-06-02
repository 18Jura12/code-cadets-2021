package tasks

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

type SignalHandler struct{}

func NewSignalHandler() *SignalHandler {
	return &SignalHandler{}
}

func (s *SignalHandler) Start(ctx context.Context) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-signals:
	case <-ctx.Done():
	}
}


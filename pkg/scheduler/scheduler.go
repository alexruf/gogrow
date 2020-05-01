package scheduler

import (
	"context"
	"log"
)

type Scheduler interface {
	Start() error
	Shutdown(ctx context.Context) error
}

type TimeScheduler struct {
	logger *log.Logger
}

func NewTimeScheduler(logger *log.Logger) Scheduler {
	return TimeScheduler{logger: logger}
}

func (s TimeScheduler) Start() error {
	s.logger.Println("TimeScheduler running!")
	return nil
}

func (s TimeScheduler) Shutdown(ctx context.Context) error {
	return nil
}

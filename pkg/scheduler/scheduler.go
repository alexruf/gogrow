package scheduler

import (
	"context"
	"github.com/alexruf/gogrow/pkg/config"
	"github.com/alexruf/gogrow/pkg/snapshot"
	"github.com/spf13/viper"
	"log"
	"time"
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

	go func() {
		interval := viper.GetInt(config.ShootIntervalMinutes)
		s.logger.Printf("Processing timer: %s", time.Now().Format(time.RFC850))
		if err := snapshot.CaptureSnapshot(); err != nil {
			s.logger.Printf("Error capturing snapshot: %s\n", err)
		}
		for range time.Tick(time.Duration(interval) * time.Minute) {
			if err := snapshot.CaptureSnapshot(); err != nil {
				s.logger.Printf("Error capturing snapshot: %s\n", err)
			}
		}
	}()

	return nil
}

func (s TimeScheduler) Shutdown(ctx context.Context) error {
	s.logger.Println("TimeScheduler shutting down!")
	return nil
}

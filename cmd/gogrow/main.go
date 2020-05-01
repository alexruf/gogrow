package main

import (
	"context"
	"github.com/alexruf/gogrow/pkg/scheduler"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"os/user"
	"path/filepath"
	"syscall"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "[GoGrow]", log.LstdFlags)
	logger.Println("Starting...")

	logger.Println("Initializing config...")
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	usr, err := user.Current()
	if err != nil {
		logger.Fatal(err)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Join(usr.HomeDir, ".gogrow"))
	viper.AddConfigPath(dir)
	err = viper.ReadInConfig()
	if err != nil {
		logger.Fatalf("Fatal error reading in config file: %s\n", err)
	}

	exitChan := make(chan bool, 1)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	timeScheduler := scheduler.NewTimeScheduler(logger) //pass config
	go gracefullShutdown(timeScheduler, logger, sigChan, exitChan)
	if err := timeScheduler.Start(); err != nil {
		logger.Fatalf("Could not start the scheduler: %v\n", err)
	}

	<-exitChan // Blocks until everything has shutdown

	logger.Println("Writing config file...")
	err = viper.WriteConfig()
	if err != nil {
		logger.Fatalf("Fatal error writing in config file: %s\n", err)
	}

	logger.Println("Done, shutting down!")
}

func gracefullShutdown(scheduler scheduler.Scheduler, logger *log.Logger, termChan <-chan os.Signal, quitChan chan<- bool) {
	<-termChan
	logger.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := scheduler.Shutdown(ctx); err != nil {
		logger.Fatalf("Could not gracefully shutdown the scheduler: %v\n", err)
	}
	close(quitChan)
}

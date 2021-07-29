package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	errStopped := errors.New("stopped")
	logger := log.New(os.Stdout, "", log.Lshortfile|log.Ldate|log.Ltime)

	gr, ctx := errgroup.WithContext(context.Background())

	gr.Go(func() error {
		configs, err := NewConfig()
		if err != nil {
			return err
		}
		var service Service //FIXME: Replace
		transport := NewTransport(service)
		return Listen(ctx, configs, transport.Routes())
	})

	gr.Go(func() error {
		//TODO: Register metrics
		return nil
	})

	gr.Go(func() error {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
		defer signal.Stop(signals)

		select {
		case <-signals:
			logger.Println("Caught signal. Exiting...")
			return errStopped
		case <-ctx.Done():
			return nil
		}
	})

	if err := gr.Wait(); err != nil && err != errStopped {
		logger.Fatal(err)
	}
}

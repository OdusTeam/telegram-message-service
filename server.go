package main

import (
	"context"
	"net/http"
	"time"
)

func Listen(ctx context.Context, cfg *Config, handler http.Handler) error {
	server := http.Server{
		Addr:        cfg.HttpAddr,
		Handler:     handler,
		IdleTimeout: time.Duration(cfg.HttpTimeout) * time.Second,
	}

	errCh := make(chan error, 1)

	go func() {
		errCh <- server.ListenAndServe()
	}()

	cancelCtx, cancelFunction := context.WithTimeout(ctx, 30*time.Second)
	defer cancelFunction()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		return server.Shutdown(cancelCtx)
	}
}

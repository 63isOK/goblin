package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func NewLogger() *log.Logger {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)
	logger.Print("Executing NewLogger.")
	return logger
}

func NewLogger2() *log.Logger {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)
	logger.Print("Executing NewLogger2.")
	return logger
}

func NewHandler(logger *log.Logger) (http.Handler, error) {
	logger.Print("Executing NewHandler.")
	return http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		logger.Print("Got a request.")
	}), nil
}

func NewMux(lc fx.Lifecycle, logger *log.Logger) *http.ServeMux {
	logger.Print("Executing NewMux.")
	// First, we construct the mux and server. We don't want to start the server
	// until all handlers are registered.
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Print("Starting HTTP server.")
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Print("Stopping HTTP server.")
			return server.Shutdown(ctx)
		},
	})

	return mux
}

func Register(mux *http.ServeMux, h http.Handler) {
	mux.Handle("/", h)
}

func main() {
	app := fx.New(
		fx.Provide(
			NewLogger,
			NewLogger2,
			NewHandler,
			NewMux,
		),
		fx.Invoke(Register),

		fx.WithLogger(
			func() fxevent.Logger {
				return fxevent.NopLogger
			},
		),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	if _, err := http.Get("http://localhost:8080/"); err != nil {
		log.Fatal(err)
	}

	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}

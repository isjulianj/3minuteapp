package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"
)

func (app *Application) ServeHTTP(mux http.Handler) error {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config.Addr),
		Handler:      mux,
		ErrorLog:     slog.NewLogLogger(app.Logger.Handler(), slog.LevelError),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Info(fmt.Sprintf("starting http server on port %s", server.Addr))
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Listen and serve:", err)
	}

	return nil
}

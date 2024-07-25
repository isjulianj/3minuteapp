package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

const (
	defaultPort        = 6969
	defaultEnvironment = "development"
	defaultStaticDir   = "assets"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	run(logger)
}

func run(logger *slog.Logger) {
	var config Config

	app := NewApplication(config, logger)

	flag.IntVar(&config.Addr, "addr", defaultPort, "The server listen address")
	flag.StringVar(&config.Environment, "env", defaultEnvironment, "The server environment")
	flag.StringVar(&config.AssetDirectory, "assetDir", defaultStaticDir, "Static asset Directory")

	flag.Parse()

	err := http.ListenAndServe(fmt.Sprintf(":%d", app.Config.Addr), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("Hello World!")); err != nil {
			panic(err)
		}
	}))

	if err != nil {
		fmt.Println("It's over anyways")
	}

	return
}

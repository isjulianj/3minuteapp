package main

import (
	"flag"
	"fmt"
	"log/slog"
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

	flag.IntVar(&config.Addr, "addr", defaultPort, "The server listen address")
	flag.StringVar(&config.Environment, "env", defaultEnvironment, "The server environment")
	flag.StringVar(&config.AssetDirectory, "assetDir", defaultStaticDir, "Static asset Directory")
	flag.Parse()

	app := NewApplication(config, logger)

	err := app.ServeHTTP(app.muxHandler())

	if err != nil {
		fmt.Println("It's over anyways")
	}

	return
}

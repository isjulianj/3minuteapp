package main

import "log/slog"

type Config struct {
	Addr           int
	Environment    string
	AssetDirectory string
	Debug          bool
}

type Application struct {
	Config Config
	Logger *slog.Logger
}

func NewApplication(config Config, logger *slog.Logger) *Application {
	return &Application{Config: config, Logger: logger}
}

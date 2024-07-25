package main

import (
	"flag"
	"fmt"
	"net/http"
)

type Config struct {
	Addr           int
	Environment    string
	AssetDirectory string
}

const (
	defaultPort        = 6969
	defaultEnvironment = "development"
	defaultStaticDir   = "assets"
)

func main() {
	run()
}

func run() {
	var config Config

	flag.IntVar(&config.Addr, "addr", defaultPort, "The server listen address")
	flag.StringVar(&config.Environment, "env", defaultEnvironment, "The server environment")
	flag.StringVar(&config.AssetDirectory, "assetDir", defaultStaticDir, "Static asset Directory")

	flag.Parse()

	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Addr), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("Hello World!")); err != nil {
			panic(err)
		}
	}))

	if err != nil {
		fmt.Println("It's over anyways")
	}

	return
}

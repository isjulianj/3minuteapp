package main

import "net/http"

func (app *Application) muxHandler() http.Handler {
	mux := http.NewServeMux()

	// home
	mux.HandleFunc("GET /{$}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("Hello World!")); err != nil {
			app.Logger.Error(err.Error())
		}
	}))

	return mux

}

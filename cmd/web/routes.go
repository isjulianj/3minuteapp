package main

import (
	"github.com/isjulianj/3minuteapp/cmd/ui/pages"
	"net/http"
)

func (app *Application) muxHandler() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./cmd/ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// home
	mux.HandleFunc("GET /{$}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := pages.Home().Render(r.Context(), w)
		if err != nil {
			app.Logger.Error(err.Error())
		}
	}))

	return mux

}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":6969", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("Hello World!")); err != nil {
			panic(err)
		}
	}))

	if err != nil {
		fmt.Println("It's over anyways")
	}
}

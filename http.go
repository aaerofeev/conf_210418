package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, AdvertPresenter(r.URL.Query().Get("id"), dataLayer))
}

// Задает обработчики запросов
func mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/show", handler)
	
	return mux
}

func run() {
	http.ListenAndServe(":8081", mux())
}

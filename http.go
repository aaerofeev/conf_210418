package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, AdvertPresenter(r.URL.Query().Get("id"), dataLayer))
}

func mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/show", handler)

	return mux
}

func RunServer() {
	http.ListenAndServe(":8081", mux())
}

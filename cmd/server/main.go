package main

import (
	"fmt"
	"net/http"
)

func main() {
	addr := "127.0.0.1:8080"
	http.Handle("/", &noCache{Handler: http.FileServer(http.Dir("./public"))})
	fmt.Printf("Listening on http://%s\n", addr)
	http.ListenAndServe(addr, nil)
}

type noCache struct {
	http.Handler
}

func (h *noCache) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache")
	h.Handler.ServeHTTP(w, r)
}

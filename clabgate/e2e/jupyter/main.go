package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		Addr:              ":8888",
		ReadHeaderTimeout: 5 * time.Second,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			_, _ = fmt.Fprint(w, "<!doctype html><title>Clabgate smoke Jupyter</title><h1>workspace ready</h1>")
		}),
	}
	log.Printf("smoke Jupyter listening on %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}

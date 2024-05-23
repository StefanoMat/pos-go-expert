package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request started")
	defer log.Println("Request ended")
	select {
	case <-time.After(5 * time.Second):
		log.Println("Request processed")
		w.Write([]byte("Request processed!"))
	case <-ctx.Done():
		log.Println("Request canceled by client")
		http.Error(w, "Request canceled by client", http.StatusRequestTimeout)
	}
}

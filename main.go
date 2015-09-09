package main

import (
	"log"
	"net/http"
	"time"
)

func handle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(200)

	for {
		w.Write([]byte("young\r\n"))
		w.(http.Flusher).Flush()
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {

	http.HandleFunc("/", handle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

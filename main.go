package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

type flushWriter struct {
	f http.Flusher
	w io.Writer
}

func (fw *flushWriter) Write(p []byte) (n int, err error) {
	n, err = fw.w.Write(p)
	if fw.f != nil {
		fw.f.Flush()
	}
	return
}

func handler(w http.ResponseWriter, r *http.Request) {
	for {
		fw := flushWriter{w: w}
		if f, ok := w.(http.Flusher); ok {
			fw.f = f
		}
		time.Sleep(time.Millisecond * 10000)
		fw.Write([]byte("WHEEE\n"))
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

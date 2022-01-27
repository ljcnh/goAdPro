package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var logger = log.New(os.Stdout, "", 0)

func hello(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("hello"))
}

func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		timeStart := time.Now()

		next.ServeHTTP(writer, request)

		timeElapsed := time.Since(timeStart)
		logger.Println(timeElapsed)
	})
}

func main() {
	http.Handle("/", timeMiddleware(http.HandlerFunc(hello)))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}))
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"net/http"
)

type application struct{}

func main() {
	app := &application{}

	log.Printf("Starting the server at port :8080")
	err := http.ListenAndServe(":8080", app.routes())
	if err != nil {
		log.Fatal(err)
	}
}

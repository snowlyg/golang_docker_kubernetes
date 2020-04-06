package main

import (
	"log"
	"net/http"
)

const PORT = "80"

func main() {
	router := NewRouter()

	//listen and serve
	log.Println("Client Server listening on port " + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}

package main

import (
	"net/http"
	"log"
)

const PORT = "8001"

func main () {
	router := NewRouter()

	//listen and serve
	log.Println("API Server listening on port " + PORT)
	log.Fatal(http.ListenAndServe(":" + PORT, router))
}

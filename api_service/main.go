package main

import (
	"log"
	"net/http"
)

const PORT = "8001"

func main() {
	router := NewRouter()

	//listen and serve
	log.Println("API Server listening on port " + PORT)
	log.Fatal(http.ListenAndServe("http://api_service.com:"+PORT, router))
}

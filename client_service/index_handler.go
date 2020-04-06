package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Greeting struct {
	Greeting string `json:"greeting"`
}

func GetNameQueryKey() string {
	return "name"
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get(GetNameQueryKey())

	apiServiceUri := os.Getenv("API")
	if apiServiceUri == "" {
		http.Error(w, "Missing API service uri", http.StatusInternalServerError)
		return
	}

	response, err := http.Get(apiServiceUri + "/api/greeting?name=" + name)
	if err != nil {
		panic(err)
	}
	greeting := &Greeting{}
	json.NewDecoder(response.Body).Decode(greeting)
	fmt.Fprint(w, greeting.Greeting)
}

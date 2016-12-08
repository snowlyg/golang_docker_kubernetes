package main

import (
	"net/http"
	"fmt"
	"os"
	"encoding/json"
)

type Greeting struct {
	Greeting	string	`json:"greeting"`
}

func GetNameQueryKey() string {
	return "name"
}

func IndexHandler(w http.ResponseWriter, r *http.Request){
	name := r.URL.Query().Get(GetNameQueryKey())


	apiServiceUri := os.Getenv("api")
	if apiServiceUri == "" {
		http.Error(w, "Missing api service uri", http.StatusInternalServerError)
		return
	}
	response, _ := http.Get(apiServiceUri + "/api/greeting?name=" + name)
	defer response.Body.Close()
	greeting := &Greeting{}
	json.NewDecoder(response.Body).Decode(greeting)
	fmt.Fprint(w, greeting.Greeting)
}

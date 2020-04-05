package main

import (
	"fmt"
	"net/http"
	// "os"
	"encoding/json"
)

type Greeting struct {
	Greeting string `json:"greeting"`
}

func GetNameQueryKey() string {
	return "name"
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get(GetNameQueryKey())

	apiServiceUri := "http://api_service.com:8001"
	response, _ := http.Get(apiServiceUri + "/api/greeting?name=" + name)
	defer response.Body.Close()
	greeting := &Greeting{}
	json.NewDecoder(response.Body).Decode(greeting)
	fmt.Fprint(w, greeting.Greeting)
}

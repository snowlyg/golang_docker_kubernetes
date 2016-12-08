package main

import (
	"net/http"
	"encoding/json"
)

type GreetingResponse struct {
	Greeting	string	`json:"greeting"`
}

func GetGreetingHandler(w http.ResponseWriter, r *http.Request){
	response := GreetingResponse{}

	name := r.URL.Query().Get("name")

	if name == "" {
		response = GreetingResponse{
			Greeting: "Hello World from Distelli!",
		}
	} else {
		greeting := ConstructGreeting(name)

		response = GreetingResponse{
			Greeting: greeting,
		}
	}



	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)

}

func ConstructGreeting(name string) (string) {

	return "Hello " + name + " you are awesome!"
}
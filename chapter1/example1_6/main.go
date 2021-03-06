package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldRequest struct {
	Name string `json:"name"`
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

const (
	imagesDirectory    = "./images"
	helloworldEndPoint = "/helloworld"
	catEndpoint        = "/cat/"
)

func main() {
	port := 8080

	catHandler := http.FileServer(http.Dir(imagesDirectory))
	http.Handle(catEndpoint, http.StripPrefix(catEndpoint, catHandler))

	http.HandleFunc(helloworldEndPoint, helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%v", port), nil),
	)
}
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: "Hello " + request.Name + " nice to meet you."}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

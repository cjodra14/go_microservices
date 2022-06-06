package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
	Author  string `json:"-"`
	Date    string `json:",omitempty"`
	Id      int    `json:"id,string"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello World!", Author: "Christian", Date: "05/07/2022", Id: 1}
	data, err := json.Marshal(response)
	if err != nil {
		panic("json marshaling error")
	}
	fmt.Fprint(w, string(data))
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type validationContextKey string

type helloWorldRequest struct {
	Name string `json:"name"`
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

func main() {
	port := 8080

	handler := newValidationHandler(newHelloWorldHandler())
	http.Handle("/helloworld", handler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%v", port), nil),
	)
}

type validationHandler struct {
	next http.Handler
}

func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

func (vH validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)

		return
	}

	ctx := context.WithValue(r.Context(), validationContextKey("name"), request.Name)
	r = r.WithContext(ctx)

	vH.next.ServeHTTP(rw, r)
}

type helloWorldHandler struct{}

func newHelloWorldHandler() http.Handler {
	return helloWorldHandler{}
}

func (hWH helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	name := r.Context().Value(validationContextKey("name")).(string)
	reponse := helloWorldResponse{Message: "Hello" + name}

	encoder := json.NewEncoder(rw)
	encoder.Encode(reponse)
}
